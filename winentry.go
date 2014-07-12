package newt

/*
#cgo CFLAGS: -I/opt/local/include
#cgo LDFLAGS: -L/opt/local/lib

#include <newt.h>
*/
import "C"

import "unsafe"

type WinEntry struct {
    we C.struct_newtWinEntry
}

func (we WinEntry) Text() string {
    return C.GoString(we.we.text)
}

func (we WinEntry) Value() uintptr {
    return uintptr(unsafe.Pointer(we.we.value))
}

func (we WinEntry) Flags() int {
    return int(we.we.flags)
}
