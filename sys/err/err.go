package err

import (
	"fmt"
	"sys"
)

type (
	// Err is a message error.
	Err struct {
		Msg        string
		StackTrace string
	}
	XprErr struct {
		Err
		Ln  uint32
		Col uint32
		Ch  rune
	}
)

// // Msg creates an Err.
// func Msg(msg string) *Err {
// 	return &Err{
// 		Msg:        msg,
// 		StackTrace: sys.StackTrace(),
// 	}
// }

// New creates an Err.
func New(v interface{}) *Err {
	return &Err{
		Msg:        fmt.Sprintf("%v", v),
		StackTrace: sys.StackTrace(),
	}
}

// Fmt creates an Err with a formatted message.
func Fmt(format string, args ...interface{}) *Err {
	return &Err{
		Msg:        fmt.Sprintf(format, args...),
		StackTrace: sys.StackTrace(),
	}
}

// Panic panics with an Err and stack trace.
func Panic(v interface{}) {
	e := &Err{
		Msg:        fmt.Sprintf("%v", v),
		StackTrace: sys.StackTrace(),
	}
	sys.Log(e)
	panic(e)
}

// Panicf panics with an Err and stack trace.
func Panicf(format string, args ...interface{}) {
	e := &Err{
		Msg:        fmt.Sprintf(format, args...),
		StackTrace: sys.StackTrace(),
	}
	sys.Log(e)
	panic(e)
}

// PanicXprf panics with an XprErr and stack trace.
func PanicXprf(format string, ln, col uint32, ch rune, args ...interface{}) {
	e := &XprErr{
		Ln:  ln,
		Col: col,
		Ch:  ch,
		Err: Err{
			Msg:        fmt.Sprintf(format, args...),
			StackTrace: sys.StackTrace(),
		},
	}
	sys.Log(e)
	panic(e)
}

// Recover creates a new Err error if there is a recovery; otherwise, nil.
func Recover() *Err {
	if v := recover(); v != nil {
		return New(v)
	}
	return nil
}

// Full returns an error message and stack trace.
func (x *Err) Full() string { return fmt.Sprintf("%v\n%v", x.Msg, x.StackTrace) }

// Error returns an error message.
func (x *Err) Error() string { return x.Msg }

// String returns an error message.
func (x *Err) String() string { return x.Msg }
