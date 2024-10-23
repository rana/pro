module Ml
using Knet
using Random
using Statistics
using AutoGrad
using Serialization
# using Threads
# using JLD2
# using Plots

const pth = "/home/rana/data/dev/"
imgPth = "/home/rana/Pictures/sys/"

function setImgPth(pth::String)
  global imgPth = pth
end

# Dense layer
struct Dense; w; b; f; end
Dense(i::Int,o::Int,f=relu) = Dense(param(o,i), param0(o), f)
function (d::Dense)(x)
  d.f.(d.w * mat(x) .+ d.b)
end

# Chain of layers
struct Chain; lyrs; end
function (c::Chain)(x; pdrop=(0.2,0.0)) # pdrop=0)
  for (i,lyr) in enumerate(c.lyrs)
    p = (i <= length(pdrop) ? pdrop[i] : pdrop[end])
    # println("i:$i  p:$p")
    x = dropout(x, p)     # <-- This one line helps generalization
    x = lyr(x)
  end
  x
end

# Neural network used to fit hst data and predict rlt. Stored to disk and retrieved during rlt for prediction.
mutable struct Net
  AvgsX::Array{Float32,1}
  StdsX::Array{Float32,1}
  Model::Chain
end

# Calculate model loss (mean square error) for a specified minibatch pair (x, y).
loss(model, x, y) = (sum(abs2,y-model(x)) / size(x,2))
# Calculate model loss (mean squared error) for the specified dataset.
loss(model, data::Data) = mean(loss(model, x, y) for (x, y) in data)

const nets = Dict{String,Net}()


function fit(key::String, x::Array{Float32,2}, y::Array{Float32,1})::Array{UInt8,1}
  println("JULIA: fit: start")
  # println("Threads.nthreads() ", Threads.nthreads())
  # @show key
  # @show summary(x)
  # @show x

  println("JULIA: fit: building...")
  inCnt = size(x, 1)
  outCnt = 1
  α = 4
  hdnCnt = α * (inCnt + outCnt)
  hdnCnt2 = round(Int, hdnCnt/3)
  hdnCnt3 = round(Int, hdnCnt/3)
  hdnCnt4 = round(Int, hdnCnt/3)
  hdnCnt5 = round(Int, hdnCnt/3)
  # model = Chain((Dense(inCnt,hdnCnt), Dense(hdnCnt,outCnt,identity)))
  model = Chain((
      Dense(inCnt,hdnCnt),
      Dense(hdnCnt,hdnCnt2),
      Dense(hdnCnt2,hdnCnt3),
      Dense(hdnCnt3,hdnCnt4),
      Dense(hdnCnt4,hdnCnt5),
      Dense(hdnCnt5,outCnt,identity)
      ))
  net = Net(zeros(size(x, 1)), zeros(size(x, 1)), model)
  for i = 1:size(x, 1)  # normalize all x-features in-place
    row = view(x, i, :)
    net.AvgsX[i] = mean(row)                # record AvgX for use by rlt
    net.StdsX[i] = stdm(row, net.AvgsX[i])  # record StdX for use by rlt
    for j in 1:size(x, 2)
      x[i,j] = (x[i,j]-net.AvgsX[i])/net.StdsX[i] # normalize x-feature in-place
    end
  end
  # @show summary(net)

  println("JULIA: fit: training...")
  # optimizer=SGD()
  # optimizer=Momentum()
  # optimizer=Nesterov()
  # optimizer=Rmsprop(lr=0.01, gclip=0, rho=0.7, eps=1e-6)
  optimizer=Rmsprop()
  # optimizer=Adagrad()
  # optimizer=Adadelta()
  # optimizer=Adam()
  epochs = 20
  epochs *= 1
  batchsize = 1
  spltPct = 0.6
  lim = convert(Int64, floor(spltPct*length(y)))
  dtrn = minibatch(x[:,1:lim], y[1:lim], batchsize)
  dtst = minibatch(x[:,lim+1:size(x, 2)], y[lim+1:length(y)], batchsize)
  results = Float64[]
  lossMin = typemax(Float64)
  modelMin = model
  ps = params(model)
  for param in ps
      param.opt = deepcopy(optimizer)
  end
  for epoch = 1:epochs
    for (x,y) in dtrn
      J = @diff loss(model,x,y)
      for param in ps
        g = grad(J,param)
        update!(value(param),g,param.opt)
      end
    end
    push!(results, loss(model,dtrn), loss(model,dtst))
    println((:epoch,epoch,:trn,results[end-1],:tst,results[end]))
    # if results[end] < lossMin
    #   lossMin = results[end]
    #   modelMin = deepcopy(model)
    # end
  end
  # println("*** lossMin final:", lossMin)
  # store in-memory for hst prediction
  net.Model = model
  nets[key] = net

  # println("JULIA: fit: ploting...")
  # results = reshape(results, (2,:))
  # @show summary(results)
  # p = plot(
  #   [results[1,:], results[2,:]],
  #   labels=[:trn :tst], ylim=(0.0,1.0),
  #   xlabel="Epochs",ylabel="Loss",
  #   title="spltPct:$spltPct, batchsize:$batchsize, epochs:$epochs")
  # png(p, imgPth*"loss.png")

  # write to array for dsk storage and later rlt usage
  # TODO: CONSIDER JLD2 SERIALIZATION
  println("JULIA: fit: serializing...")
  buf = Array{UInt8,1}()
  io = IOBuffer(buf;read=true,write=true)
  serialize(io, net)
  flush(io)

  # net2 = deserialize(IOBuffer(buf))
  # nets[key]=net2
  # @show length(buf)
  # @show buf
  # @show length(buf)

  println("JULIA: fit: end")
  return buf
end
function loadNet(key::String, buf::Array{UInt8,1})
  println("JULIA: loadNet: start")
  # @show buf
  net = deserialize(IOBuffer(buf))
  # @show net
  # delete!(nets, key)
  nets[key] = net
  println("JULIA: loadNet: end")
end
function predict(key::String, x::Array{Float32,1})::Float32
  # println("JULIA: predict: haskey: ", haskey(nets, key))
  if !haskey(nets, key) # missing key for stgy with no available trds
    println("JULIA: predict: key missing")
    println(key)
    return typemin(Float32)
  end
  net = nets[key]
  # @show net.AvgsX
  # @show net.StdsX
  # @show x
  for i = 1:length(x) # normalize x-feature in-place
    x[i] = (x[i]-net.AvgsX[i])/net.StdsX[i]
  end
  # @show summary(net.Model(x))
  # @show x
  # @show net.Model(x)[1,1]
  net.Model(x)[1,1]
end
function ping()
end


function remNetMem(key::String)
  if haskey(nets, key)
    delete!(nets, key)
  end
end

function test(x::Array{Float32,1})
  @show x
end


end # module



# filename = "tmp2.jld2"
# @save filename dtrn dtst