package ml

/*
#cgo CFLAGS: -fPIC -DJULIA_INIT_DIR="/usr/local/julia-1.0.0/lib" -I/usr/local/julia-1.0.0/include/julia -I.
#cgo LDFLAGS: -L/usr/local/julia-1.0.0/lib/julia  -L/usr/local/julia-1.0.0/lib -Wl,-rpath,/usr/local/julia-1.0.0/lib -ljulia
#define JULIA_ENABLE_THREADING
#include <stdio.h>
#include <julia.h>
JULIA_DEFINE_FAST_TLS() // only define this once, in an executable (not in a shared library) if you want fast code.
typedef struct {
	ssize_t rowCnt;
	ssize_t colCnt;
} ntuple2int;
static inline ntuple2int* Tuple(int colCnt, int rowCnt) {
	jl_value_t *types[] = {(jl_value_t*)jl_long_type, (jl_value_t*)jl_long_type};
	jl_tupletype_t *tt = jl_apply_tuple_type_v(types, 2);
	ntuple2int *tuple = (ntuple2int*)jl_new_struct_uninit(tt);
	JL_GC_PUSH1(&tuple);
	tuple->rowCnt = rowCnt;
	tuple->colCnt = colCnt;
	JL_GC_POP();
	return tuple;
}
static inline char* ArrData(jl_array_t *a) {
	return (char*)jl_array_data(a);
}
static inline size_t ArrLen(jl_array_t *a) {
	return jl_array_len(a);
}
static inline void PrintEx() {
	if (jl_exception_occurred())
    printf("%s \n", jl_typeof_str(jl_exception_occurred()));
}
*/
import "C"
import (
	"math"
	"runtime"
	"io/ioutil"
	"fmt"
	"reflect"
	"unsafe"
	"sys"
	"sys/err"
	"sys/trc"
)

const (
	Filename = "ml.jl"
)

type (
	Lrnr struct {
		Filepath string
		Opnd     bool
		JlModule *C.jl_module_t
		exitC    chan bool
		actC 		 chan sys.Act
	}
)

func NewLrnr(filepath string) (r *Lrnr) {
	r = &Lrnr{}
	r.Filepath = filepath
	r.exitC = make(chan bool)
	r.actC = make(chan sys.Act)
	r.Opn()
	return r
}
func (x *Lrnr) Opn() {
	if !x.Opnd {
		trcr := trc.New("ml.Lrnr.Opn")
		defer trcr.End()
		sys.Log(x.Filepath)

		started := make(chan bool)
		go x.loop(started, x.Filepath)
		<-started

		x.Opnd = true
	}
}
func (x *Lrnr) Cls() {
	sys.Log("Lrnr.Cls")
	x.Opnd = false
	x.exitC <- true // request loop exit
}

// OPERATE ALL JULIA FUNCTIONS ON SAME OS THREAD
// ISSUE IS DIFFERENT GOLANG OS THREAD/GO ROUTINE ENTERING INTO SINGLE-THREADED JULIA ENVIRONMENT CAUSES SEG FAULT
// AVOID BUG WITH HST TRAIN / RLT 'PREDICT' FROM DIFFERENT GO ROUTINES
func (x *Lrnr) loop(started chan bool, filepath string) {
	defer func() {
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
	}()

	// lock OS thread to avoid JULIA bugs (suspect, but not certain, thread-local storage within JULIA)
	runtime.LockOSThread()
	C.jl_init()
	defer C.jl_atexit_hook(0)
	content, er := ioutil.ReadFile(filepath)
	if er != nil {
		err.Panic(er)
	}
	// sys.Log(string(content))
	content_c := C.CString(string(content))
	defer C.free(unsafe.Pointer(content_c))
	valMod := C.jl_eval_string(content_c)
	x.JlModule = (*C.jl_module_t)(unsafe.Pointer(valMod))
	started <- true

	for {
		select {
		case <-x.exitC:
			return
		case a := <- x.actC:
			a.Act()
		}
	}
}

type (
	FitCmd struct {
		lrnr *Lrnr
		key string
		ftrNames []string
		ftrs [][]float32
		lbls []float32
		bits []byte
		done chan bool
	}
)
func (x *Lrnr) Fit(key string, ftrNames []string, ftrs [][]float32, lbls []float32) (bits []byte) {
	cmd := &FitCmd{}
	cmd.lrnr = x
	cmd.key = key
	cmd.ftrNames = ftrNames
	cmd.ftrs = ftrs
	cmd.lbls = lbls
	cmd.done = make(chan bool)
	x.actC <- cmd
	<- cmd.done
	return cmd.bits
}
func (x *FitCmd) Act() {
	defer func(){
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
		x.done <- true
	}()
	trcr := trc.Newf("ml.Lrnr.Fit: fitting (cols:%v rows:%v)", len(x.ftrNames), len(x.lbls))
	defer trcr.End()
	// sys.Logf("key: %v \n", key)

	// ftrs = make([][]float32, 2)
	// ftrs[0] = []float32{1,2,3,4}
	// ftrs[1] = []float32{5,6,7,8}
	// lbls = []float32{9,10,11,12}

	if len(x.lbls) == 0 {
		err.Panic("lbl cnt zero")
	}
	for n, ftr := range x.ftrs {
		if len(ftr) != len(x.lbls) {
			sys.Logf("not enough data: ftr cnt unequal to lbl cnt (ftrIdx:%v ftrCnt:%v lblCnt:%v)", n, len(ftr), len(x.lbls))
			return
		}
	}
	// for n, ftr := range ftrs {
	// 	sys.Log("GO: n:", n, "len:", len(ftr), "ftr:", ftr)
	// }
	// sys.Log("GO:", "len:", len(lbls), "lbls:", lbls)

	colCnt, rowCnt := uint64(len(x.ftrs[0])), uint64(len(x.ftrs))
	flt_elm_typ := (*C.jl_value_t)(unsafe.Pointer(C.jl_float32_type))

	// x
	var idx uint64
	x_go := make([]float32, colCnt*rowCnt) // TRANSFORM TO JULIA COLUMN-MAJOR 2D ARRAY
	for colIdx := uint64(0); colIdx < colCnt; colIdx++ { // col
		for rowIdx := uint64(0); rowIdx < rowCnt; rowIdx++ { // row
			x_go[idx]=x.ftrs[rowIdx][colIdx]
			idx++
		}
	}
	x_go_hdr := (*reflect.SliceHeader)(unsafe.Pointer(&x_go))
	x_go_data := unsafe.Pointer(x_go_hdr.Data)
	x_arr_type := C.jl_apply_array_type(flt_elm_typ, C.ulong(uint64(2)))
	x_dims_tuple := C.Tuple(C.int(int(colCnt)), C.int(int(rowCnt)))
	x_dims := (*C.jl_value_t)(unsafe.Pointer(x_dims_tuple))
	x_arr := C.jl_ptr_to_array(x_arr_type, x_go_data, x_dims, C.int(0))
	// y
	y_go_hdr := (*reflect.SliceHeader)(unsafe.Pointer(&x.lbls))
	y_go_data := unsafe.Pointer(y_go_hdr.Data)
	y_arr_type := C.jl_apply_array_type(flt_elm_typ, C.ulong(uint64(1)))
	y_arr := C.jl_ptr_to_array_1d(y_arr_type, y_go_data, C.size_t(uint64(len(x.lbls))), C.int(0))
	// call julia 'fit' function
	fit_c := C.CString(`fit`)
	defer C.free(unsafe.Pointer(fit_c))
	key_c := C.CString(x.key)
	defer C.free(unsafe.Pointer(key_c))
	fn := C.jl_get_function(x.lrnr.JlModule, fit_c)
	key_prm := C.jl_cstr_to_string(key_c)
	x_arr_prm := (*C.jl_value_t)(unsafe.Pointer(x_arr))
	y_arr_prm := (*C.jl_value_t)(unsafe.Pointer(y_arr))
	bits_val_jl := C.jl_call3(fn, key_prm, x_arr_prm, y_arr_prm)
	x.lrnr.JlErrChk()


	bits_arr_jl := (*C.jl_array_t)(unsafe.Pointer(bits_val_jl))
	bits_data_jl := C.ArrData(bits_arr_jl)
	bits_len := C.ArrLen(bits_arr_jl)
	// GoBytes: It is important to keep in mind that the Go garbage collector will not interact with this data, and that if it is freed from the C side of things, the behavior of any Go code using the slice is nondeterministic.
	bits_jl := C.GoBytes(unsafe.Pointer(bits_data_jl), C.int(bits_len))
	// fmt.Println("--- fit bits", len(bits_jl))
	// var sb strings.Builder
	// for _, b := range bits_jl {
	// 	sb.WriteString(fmt.Sprintf("0x%x, ", b))
	// }
	// fmt.Println(sb.String())
	// fmt.Println("--- fit bits", len(bits_jl))
	// TODO: COPY BITS; JL WILL GC BITS
	x.bits = make([]byte, C.int(bits_len))
	copy(x.bits, bits_jl)
	// sys.Logf("fit bits %x %x %x", bits[0], bits[1], bits[2])
	// sys.Logf("fit bits %x %x %x", bits[len(bits)-3], bits[len(bits)-2], bits[len(bits)-1])
}

// func (x *Lrnr) Fit_PREV(key string, ftrNames []string, ftrs [][]float32, lbls []float32) (bits []byte) {
// 	x.Opn()
// 	trcr := trc.Newf("ml.Lrnr.Fit: fitting (cols:%v rows:%v)", len(ftrNames), len(lbls))
// 	defer trcr.End()
// 	// sys.Logf("key: %v \n", key)

// 	// ftrs = make([][]float32, 2)
// 	// ftrs[0] = []float32{1,2,3,4}
// 	// ftrs[1] = []float32{5,6,7,8}
// 	// lbls = []float32{9,10,11,12}

// 	if len(lbls) == 0 {
// 		err.Panic("lbl cnt zero")
// 	}
// 	for n, ftr := range ftrs {
// 		if len(ftr) != len(lbls) {
// 			err.Panicf("ftr cnt unequal to lbl cnt (ftrIdx:%v ftrCnt:%v lblCnt:%v)", n, len(ftr), len(lbls))
// 		}
// 	}
// 	// for n, ftr := range ftrs {
// 	// 	sys.Log("GO: n:", n, "len:", len(ftr), "ftr:", ftr)
// 	// }
// 	// sys.Log("GO:", "len:", len(lbls), "lbls:", lbls)

// 	colCnt, rowCnt := uint64(len(ftrs[0])), uint64(len(ftrs))
// 	flt_elm_typ := (*C.jl_value_t)(unsafe.Pointer(C.jl_float32_type))

// 	// x
// 	var idx uint64
// 	x_go := make([]float32, colCnt*rowCnt) // TRANSFORM TO JULIA COLUMN-MAJOR 2D ARRAY
// 	for colIdx := uint64(0); colIdx < colCnt; colIdx++ { // col
// 		for rowIdx := uint64(0); rowIdx < rowCnt; rowIdx++ { // row
// 			x_go[idx]=ftrs[rowIdx][colIdx]
// 			idx++
// 		}
// 	}
// 	x_go_hdr := (*reflect.SliceHeader)(unsafe.Pointer(&x_go))
// 	x_go_data := unsafe.Pointer(x_go_hdr.Data)
// 	x_arr_type := C.jl_apply_array_type(flt_elm_typ, C.ulong(uint64(2)))
// 	x_dims_tuple := C.Tuple(C.int(int(colCnt)), C.int(int(rowCnt)))
// 	x_dims := (*C.jl_value_t)(unsafe.Pointer(x_dims_tuple))
// 	x_arr := C.jl_ptr_to_array(x_arr_type, x_go_data, x_dims, C.int(0))
// 	// y
// 	y_go_hdr := (*reflect.SliceHeader)(unsafe.Pointer(&lbls))
// 	y_go_data := unsafe.Pointer(y_go_hdr.Data)
// 	y_arr_type := C.jl_apply_array_type(flt_elm_typ, C.ulong(uint64(1)))
// 	y_arr := C.jl_ptr_to_array_1d(y_arr_type, y_go_data, C.size_t(uint64(len(lbls))), C.int(0))
// 	// call julia 'fit' function
// 	fit_c := C.CString(`fit`)
// 	defer C.free(unsafe.Pointer(fit_c))
// 	key_c := C.CString(key)
// 	defer C.free(unsafe.Pointer(key_c))
// 	fn := C.jl_get_function(x.JlModule, fit_c)
// 	key_prm := C.jl_cstr_to_string(key_c)
// 	x_arr_prm := (*C.jl_value_t)(unsafe.Pointer(x_arr))
// 	y_arr_prm := (*C.jl_value_t)(unsafe.Pointer(y_arr))
// 	bits_val_jl := C.jl_call3(fn, key_prm, x_arr_prm, y_arr_prm)
// 	x.JlErrChk()


// 	bits_arr_jl := (*C.jl_array_t)(unsafe.Pointer(bits_val_jl))
// 	bits_data_jl := C.ArrData(bits_arr_jl)
// 	bits_len := C.ArrLen(bits_arr_jl)
// 	// GoBytes: It is important to keep in mind that the Go garbage collector will not interact with this data, and that if it is freed from the C side of things, the behavior of any Go code using the slice is nondeterministic.
// 	bits_jl := C.GoBytes(unsafe.Pointer(bits_data_jl), C.int(bits_len))
// 	// fmt.Println("--- fit bits", len(bits_jl))
// 	// var sb strings.Builder
// 	// for _, b := range bits_jl {
// 	// 	sb.WriteString(fmt.Sprintf("0x%x, ", b))
// 	// }
// 	// fmt.Println(sb.String())
// 	// fmt.Println("--- fit bits", len(bits_jl))
// 	// TODO: COPY BITS; JL WILL GC BITS
// 	bits = make([]byte, C.int(bits_len))
// 	copy(bits, bits_jl)
// 	// sys.Logf("fit bits %x %x %x", bits[0], bits[1], bits[2])
// 	// sys.Logf("fit bits %x %x %x", bits[len(bits)-3], bits[len(bits)-2], bits[len(bits)-1])
// 	return bits
// }

type (
	PredictCmd struct {
		lrnr *Lrnr
		key string
		ftrs []float32
		y float32
		done chan bool
	}
)
func (x *Lrnr) Predict(key string, ftrs []float32) (y float32) {
	cmd := &PredictCmd{}
	cmd.lrnr = x
	cmd.key = key
	cmd.ftrs = ftrs
	cmd.done = make(chan bool)
	x.actC <- cmd
	<- cmd.done
	return cmd.y
}

func (x *PredictCmd) Act() {
	defer func(){
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
			x.y = float32(math.Inf(-1))
		}
		x.done <- true
	}()
	// trcr := trc.Newf("ml.Lrnr.Predict")
	// sys.Logf("ml.Lrnr.Predict: key: %v", key)
	// sys.Logf("ml.Lrnr.Predict: ftrs:%v", ftrs)
	flt_elm_typ := (*C.jl_value_t)(unsafe.Pointer(C.jl_float32_type))
	x_go_hdr := (*reflect.SliceHeader)(unsafe.Pointer(&x.ftrs))
	x_go_data := unsafe.Pointer(x_go_hdr.Data)
	x_arr_type := C.jl_apply_array_type(flt_elm_typ, C.ulong(uint64(1)))
	x_arr := C.jl_ptr_to_array_1d(x_arr_type, x_go_data, C.size_t(uint64(len(x.ftrs))), C.int(0))

	predict_c := C.CString(`predict`)
	defer C.free(unsafe.Pointer(predict_c))
	key_c := C.CString(x.key)
	defer C.free(unsafe.Pointer(key_c))

	fn := C.jl_get_function(x.lrnr.JlModule, predict_c)
	key_prm := C.jl_cstr_to_string(key_c)
	x_arr_prm := (*C.jl_value_t)(unsafe.Pointer(x_arr))
	ret_val_jl := C.jl_call2(fn, key_prm, x_arr_prm)
	x.lrnr.JlErrChk()
	// ret_val_c := C.jl_unbox_float32(ret_val_jl)
	// y = float32(ret_val_c)
	// trcr.End(fmt.Sprintf("y:%v", y))
	x.y = float32(C.jl_unbox_float32(ret_val_jl))
}
// func (x *Lrnr) Predict_PREV(key string, ftrs []float32) (y float32) {
// 	x.Opn() // try opn for fst rlt
// 	// trcr := trc.Newf("ml.Lrnr.Predict")
// 	// sys.Logf("ml.Lrnr.Predict: key: %v", key)
// 	// sys.Logf("ml.Lrnr.Predict: ftrs:%v", ftrs)
// 	flt_elm_typ := (*C.jl_value_t)(unsafe.Pointer(C.jl_float32_type))
// 	x_go_hdr := (*reflect.SliceHeader)(unsafe.Pointer(&ftrs))
// 	x_go_data := unsafe.Pointer(x_go_hdr.Data)
// 	x_arr_type := C.jl_apply_array_type(flt_elm_typ, C.ulong(uint64(1)))
// 	x_arr := C.jl_ptr_to_array_1d(x_arr_type, x_go_data, C.size_t(uint64(len(ftrs))), C.int(0))

// 	predict_c := C.CString(`predict`)
// 	defer C.free(unsafe.Pointer(predict_c))
// 	key_c := C.CString(key)
// 	defer C.free(unsafe.Pointer(key_c))

// 	fn := C.jl_get_function(x.JlModule, predict_c)
// 	key_prm := C.jl_cstr_to_string(key_c)
// 	x_arr_prm := (*C.jl_value_t)(unsafe.Pointer(x_arr))
// 	ret_val_jl := C.jl_call2(fn, key_prm, x_arr_prm)
// 	x.JlErrChk()
// 	// ret_val_c := C.jl_unbox_float32(ret_val_jl)
// 	// y = float32(ret_val_c)
// 	// trcr.End(fmt.Sprintf("y:%v", y))
// 	return float32(C.jl_unbox_float32(ret_val_jl))
// }

type (
	SaveNetToDskCmd struct {
		lrnr *Lrnr
		key string
		net []byte
		done chan bool
	}
)
func (x *Lrnr) SaveNetToDsk(key string, net []byte) {
	cmd := &SaveNetToDskCmd{}
	cmd.lrnr = x
	cmd.key = key
	cmd.net = net
	cmd.done = make(chan bool)
	x.actC <- cmd
	<- cmd.done
}
func (x *SaveNetToDskCmd) Act() {
	defer func(){
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
		x.done <- true
	}()
	sys.Dsk().SavMl([]byte(x.key), x.net)
}
// func (x *Lrnr) SaveNetToDsk(key string, net []byte) {
// 	sys.Dsk().SavMl([]byte(key), net)
// }

type (
	LoadNetFromDskCmd struct {
		lrnr *Lrnr
		key string
		done chan bool
	}
)
func (x *Lrnr) LoadNetFromDsk(key string) {
	cmd := &LoadNetFromDskCmd{}
	cmd.lrnr = x
	cmd.key = key
	cmd.done = make(chan bool)
	x.actC <- cmd
	<- cmd.done
}
func (x *LoadNetFromDskCmd) Act() {
	defer func(){
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
		x.done <- true
	}()
	net := sys.Dsk().LoadMl([]byte(x.key))
	loadNet_c := C.CString(`loadNet`)
	defer C.free(unsafe.Pointer(loadNet_c))
	key_c := C.CString(x.key)
	defer C.free(unsafe.Pointer(key_c))

	uint_elm_typ := (*C.jl_value_t)(unsafe.Pointer(C.jl_uint8_type))
	net_go_hdr := (*reflect.SliceHeader)(unsafe.Pointer(&net))
	net_go_data := unsafe.Pointer(net_go_hdr.Data)
	net_arr_type := C.jl_apply_array_type(uint_elm_typ, C.ulong(uint64(1)))
	net_arr := C.jl_ptr_to_array_1d(net_arr_type, net_go_data, C.size_t(uint64(len(net))), C.int(0))

	fn := C.jl_get_function(x.lrnr.JlModule, loadNet_c)
	key_prm := C.jl_cstr_to_string(key_c)
	net_prm := (*C.jl_value_t)(unsafe.Pointer(net_arr))
	C.jl_call2(fn, key_prm, net_prm)
	x.lrnr.JlErrChk()
}
// func (x *Lrnr) LoadNetFromDsk(key string) {
// 	net := sys.Dsk().LoadMl([]byte(key))
// 	x.LoadNet(key, net)
// }

func (x *Lrnr) JlErrChk() {
	if C.jl_exception_occurred() != nil {
		fmt.Println("-- JULIA ERROR")
		// C.PrintEx()
		jl_ex := C.jl_exception_occurred()
		jl_err_msg_c := C.jl_typeof_str(jl_ex)
		fmt.Println("  ", C.GoString(jl_err_msg_c))
		show_c := C.CString("show")
		defer C.free(unsafe.Pointer(show_c))
		C.jl_call2(C.jl_get_function(C.jl_base_module, show_c), C.jl_stderr_obj(), jl_ex)
	}
}