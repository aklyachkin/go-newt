package newt

/*
#cgo CFLAGS: -I/opt/local/include
#cgo LDFLAGS: -L/opt/local/lib

#include <newt.h>
*/
import "C"

import "unsafe"

const (     // newtExitStruct reason
    EXIT_HOTKEY = iota
    EXIT_COMPONENT
    EXIT_FDREADY
    EXIT_TIMER
    EXIT_ERROR
)

type ExitStruct struct {
    es C.struct_newtExitStruct
}

func (es ExitStruct) Reason() int {
    return int(es.es.reason)
}

func (es ExitStruct) Component() Component {
    var c Component
    c.c = *(*C.newtComponent)(unsafe.Pointer(&es.es.u[0]))
    return c
}

func (es ExitStruct) Watch() int {
    var i32 int32
    i32 = *(*int32)(unsafe.Pointer(&es.es.u[0]))
    return int(i32)
}

func (es ExitStruct) Key() int {
    var i32 int32
    i32 = *(*int32)(unsafe.Pointer(&es.es.u[0]))
    return int(i32)
}
