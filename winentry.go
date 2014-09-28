package newt

/*
#cgo pkg-config: libnewt
#cgo LDFLAGS: -lnewt

#include <newt.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"

import "unsafe"

type WinEntry struct {
    we C.struct_newtWinEntry
    dv ResultStr
}

func NewWinEntry(text string, flags int) WinEntry {
    var we WinEntry
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    we.we.text = (*C.char)(C.calloc(C.size_t(len(text) + 2), 1))
    C.strncpy(we.we.text, t, C.size_t(len(text)))
    we.we.value = nil
    we.we.flags = C.int(flags)
    we.dv = NewResultStr()
    return we
}

func NewWinEntryDef(text, defaultValue string, flags int) WinEntry {
    we := NewWinEntry(text, flags)
    we.dv.Set(defaultValue)
    we.we.value = &we.dv.value
    return we
}

func (we WinEntry) Text() string {
    return C.GoString(we.we.text)
}

func (we WinEntry) Value() string {
    return we.dv.String()
    //return uintptr(unsafe.Pointer(we.we.value))
}

func (we WinEntry) Flags() int {
    return int(we.we.flags)
}

func (we WinEntry) Destroy() {
    we.dv.Destroy()
    C.free(unsafe.Pointer(we.we.text))
}
