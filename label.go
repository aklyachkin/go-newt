package newt

/*
#cgo CFLAGS: -I/opt/local/include
#cgo LDFLAGS: -L/opt/local/lib -lnewt

#include <stdio.h>
#include <stdlib.h>
#include <newt.h>
*/
import "C"

import (
    "unsafe"
)

func Label(left, top int, text string) Component {
    var c Component
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    c.c = C.newtLabel(C.int(left), C.int(top), t)
    c.t = GRID_COMPONENT
    return c
}

func LabelSetText(c Component, text string) {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    C.newtLabelSetText(c.c, t)
}

func LabelSetColors(c Component, colorset int) {
    C.newtLabelSetColors(c.c, C.int(colorset))
}

