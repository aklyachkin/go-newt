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

func Scale(left, top, width int, fullValue int64) Component {
    var c Component
    c.c = C.newtScale(C.int(left), C.int(top), C.int(width), C.longlong(fullValue))
    return c
}

func ScaleSet(c Component, amount uint64) {
    C.newtScaleSet(c.c, C.ulonglong(amount))
}

func ScaleSetColors(c Component, empty, full int) {
    C.newtScaleSetColors(c.c, C.int(empty), C.int(full))
}

