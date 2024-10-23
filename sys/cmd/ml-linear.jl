module Ml
using Statistics
using Random
using Knet
using AutoGrad
using Random
using LinearAlgebra: axpy! # axpy!(x,y) sets y[:]=a*x+y
using ProgressMeter: @showprogress
using Plots#; default(fmt = :png)
# using Plots
# using StatPlots
# gr()

imgpth = "/home/rana/Pictures/sys/"

# used by fit and predict. stored to disk and retrieved during rlt prediction
struct Stat
  AvgsX::Array{Float32,1}
  StdsX::Array{Float32,1}
end

# We will use a callable object to define our linear model
struct Linear; w; b; end
(model::Linear)(x) = model.w * x .+ model.b



function fit(key::String, xsrc::Array{Float32,2}, ysrc::Array{Float32,1})
  println("JULIA: fit: start")

  Random.seed!(9)

  s = Stat(zeros(size(xsrc, 1)), zeros(size(xsrc, 1)))
  for i = 1:size(xsrc, 1)  # normalize all x-features in-place
    row = view(xsrc, i, :)
    s.AvgsX[i] = mean(row)              # record AvgX for use by rlt
    s.StdsX[i] = stdm(row, s.AvgsX[i])  # record StdX for use by rlt
    for j in 1:size(xsrc, 2)
      xsrc[i,j] = (xsrc[i,j]-s.AvgsX[i])/s.StdsX[i] # normalize x-feature in-place
    end
  end
  @show s

  spltPct = 0.8
  batchsize = 1
  lim = convert(Int64, floor(spltPct*length(ysrc)))
  dtrn = minibatch(xsrc[:,1:lim], ysrc[1:lim], batchsize)
  dtst = minibatch(xsrc[:,lim+1:size(xsrc, 2)], ysrc[lim+1:length(ysrc)], batchsize)
  # @show first(dtrn)
  # @show first(dtst)

  # Initialize a random Linear model
  inCnt = size(xsrc, 1)
  outCnt = 1
  w = randn(outCnt, inCnt)*0.01
  b = zeros(outCnt)
  model = Linear(w, b)
  @show model

  x, y = first(dtst)
  @show x, y
  ypred = model(x)
  @show ypred

  # We can calculate the accuracy of our model for the first minibatch
  accuracy(model, x, y) = mean(y' .== map(i->i[1], findmax(Array(model(x)),dims=1)[2]))
  @show accuracy(model, x, y)

  # We can calculate the accuracy of our model for the whole test set
  accuracy(model,data::Data) = mean(accuracy(model,x,y) for (x,y) in data)
  @show accuracy(model,dtst)

  # ZeroOne loss (or error) is defined as 1 - accuracy
  zeroone(x...) = 1 - accuracy(x...)
  @show zeroone(model, dtst)

  # Calculate loss (mean square error) of our model for the first minibatch
  # loss(w,x,y)=(sum(abs2,y-predict(w,x)) / size(x,2))
  loss(model, x, y) = (sum(abs2,y-model(x)) / size(x,2))
  @show loss(model, x, y)

  # per-instance average loss (mean square error) for the whole test set
  loss(model, data::Data) = mean(loss(model, x, y) for (x, y) in data)
  @show loss(model, dtst)

  # To compute gradients we need to mark fields of f as Params:
  w = Param(randn(outCnt, inCnt)*0.01)
  b = Param(zeros(outCnt))
  model = Linear(w, b)
  @show model

  # We can still do predictions with f and calculate loss:
  @show loss(model, x, y)

  # And we can do the same loss calculation also computing gradients:
  J = @diff loss(model, x, y)
  @show value(J)

  # To get the gradient of a parameter from J:
  ∇w = grad(J, model.w)
  @show ∇w
  ∇b = grad(J, model.b)
  @show ∇b

  function train2!(model, data, loss)
    for (x,y) in data
        lossV = @diff loss(model, x, y)
        for param in (model.w, model.b)
            ∇param = grad(lossV, param)
            axpy!(-0.1, ∇param, value(param))
        end
    end
  end

  # Let's collect some data to draw training curves and visualizing weights:
  function trainresults(file, loss, epochs, inCnt, outCnt)
    results = []
    pa(x) = Knet.gpu() >= 0 ? Param(KnetArray{Float32}(x)) : Param(Array{Float32}(x))
    model = Linear(pa(randn(outCnt,inCnt)*0.01), pa(zeros(outCnt)))
    @showprogress for epoch in 1:epochs  # 100ep 77s (0.2668, 0.0744)
        push!(results, deepcopy(model), loss(model, dtrn), loss(model, dtst), accuracy(model, dtrn), accuracy(model, dtst))
        train2!(model, dtrn, loss)
    end
    results = reshape(results, (5, :))
    # Knet.save(file,"results",results)

  end

  #train with epochs
  lin = trainresults("lin.jld2", loss, 100, inCnt, outCnt)
  # lin = Knet.load("lin.jld2","results")
  @show minimum(lin[3,:]), minimum(lin[5,:])

  # Demonstrates underfitting: training loss not close to 0
  # Also slight overfitting: test loss higher than train
  # p = plot([lin[2,:], lin[3,:]],ylim=(.0,.4),labels=[:trnloss :tstloss],xlabel="Epochs",ylabel="Loss")
  p = plot([lin[2,:], lin[3,:]],labels=[:trnloss :tstloss],xlabel="Epochs",ylabel="Loss")
  png(p, imgpth*"linloss.png")

  # this is the error plot, we get to about 7.5% test error, i.e. 92.5% accuracy
  p = plot([lin[4,:], lin[5,:]],labels=[:trnerr :tsterr],xlabel="Epochs",ylabel="Error")
  png(p, imgpth*"linerr.png")

  println("JULIA: fit: end")
end


end # module



# struct Stat # used by fit and predict. stored to disk and retrieved during realtime prediction
#   AvgsX::Array{Float32,1}
#   StdsX::Array{Float32,1}
# end
# struct Set
#   x::Array{Float32,2}
#   y::Array{Float32,1}
# end
# function Base.iterate(v::Set, i=1)
#   if i > length(v.y)
#     return nothing
#   end
#   ((view(v.x, :, i), v.y[i]), i+1)
# end
# function print(v::Set)
#   for (x, y) in v
#     println("x:$x  y:$y")
#   end
# end
# struct Dat
#   trn::Set
#   tst::Set
# end
# function print(v::Dat)
#   println("-- trn ", length(v.trn.y))
#   print(v.trn)
#   println("-- tst ", length(v.tst.y))
#   print(v.tst)
#   println("-- trn:", length(v.trn.y), " tst:", length(v.tst.y), " input:", size(v.trn.x, 1), " output:1")
# end
# function splt(x::Array{Float32,2}, y::Array{Float32,1}, spltPct::Float64)::Dat
#   lim = convert(Int64, floor(spltPct*length(y)))
#   # TODO: SHUFFLE TRN?
#   Dat(
#     Set(x[:,1:lim], y[1:lim]), #trn
#     Set(x[:,lim+1:size(x, 2)], y[lim+1:length(y)]) # tst
#   )
# end




# function fit_PREV(key::String, x::Array{Float32,2}, y::Array{Float32,1})
#   println("JULIA: fit: start")
#   # println("key: ", key)
#   # println("x ", x)
#   # println("y ", y)

#   # s = Stat(zeros(size(x, 1)), zeros(size(x, 1)))
#   # for i = 1:size(x, 1)  # normalize all x-features in-place
#   #   row = view(x, i, :)
#   #   # println("x[$i,:] ", row)
#   #   s.AvgsX[i] = mean(row)              # record AvgX for use by rlt
#   #   s.StdsX[i] = stdm(row, s.AvgsX[i])  # record StdX for use by rlt
#   #   # println("s.AvgsX[$i] ", s.AvgsX[i])
#   #   # println("s.StdsX[$i] ", s.StdsX[i])
#   #   for j in 1:size(x, 2)
#   #     # println("$i,$j: pre ", x[i,j])
#   #     x[i,j] = (x[i,j]-s.AvgsX[i])/s.StdsX[i]
#   #     # println("$i,$j: pst ", x[i,j])
#   #   end
#   # end
#   # dat = splt(x, y, 0.8)
#   # # print(dat)
#   # α = 2
#   # nineurons = size(x, 1)
#   # noneurons = 1
#   # nhneurons = α * (nineurons + noneurons)
#   # nlayers = [nineurons, nhneurons, convert(Int, nhneurons/2), convert(Int, nhneurons/3)]
#   # @show nlayers
#   # # model = MLP(nlayers...)
#   # # @which summary.(l.w for l in model.layers)
#   # # println(summary.(l.w for l in model.layers))
#   # # @time train!(model, dat.trn)

#   # xtrn, ytrn, xtst, ytst = housing()
#   # println("- xtrn")
#   # println(xtrn)
#   # println("- ytrn")
#   # println(ytrn)
#   # println("xtrn:", typeof(xtrn), " ytrn:", typeof(ytrn))
#   # println("xtrn:", size(xtrn), " ytrn:", size(ytrn))


#   # dtrn,dtst = mnistdata(xtype=Array{Float32})
#   # @show length.((dtrn,dtst))
#   # (x,y) = first(dtst)
#   # @show summary.((x,y))

#   # @show first(dat.trn)
#   # @show summary.(first(dat.trn))

#   b = minibatch(x, y, 1)
#   @show first(b)
#   @show summary.(first(b))

#   println("JULIA: fit: end")
# end