package newt

/*
#cgo pkg-config: libnewt
#cgo LDFLAGS: -lnewt

#include <stdio.h>
#include <stdlib.h>
#include <newt.h>
*/
import "C"

import (
    "unsafe"
)

func CompactButton(left, top int, text string) Component {
    var c Component
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    a := C.newtCompactButton(C.int(left), C.int(top), t)
    c.c = a
    c.t = GRID_COMPONENT
    c.ct = CMP_COMPACTBUTTON
    return c
}

func Button(left, top int, text string) Component {
    var c Component
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    a := C.newtButton(C.int(left), C.int(top), t)
    c.c = a
    c.t = GRID_COMPONENT
    c.ct = CMP_BUTTON
    return c
}

