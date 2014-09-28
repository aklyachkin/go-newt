package newt

/*
#cgo pkg-config: libnewt
#cgo LDFLAGS: -lnewt

#include <stdio.h>
#include <stdlib.h>
#include <newt.h>
*/
import "C"

func Scale(left, top, width int, fullValue int64) Component {
    var c Component
    c.c = C.newtScale(C.int(left), C.int(top), C.int(width), C.longlong(fullValue))
    c.t = GRID_COMPONENT
    return c
}

func ScaleSet(c Component, amount uint64) {
    C.newtScaleSet(c.c, C.ulonglong(amount))
}

func ScaleSetColors(c Component, empty, full int) {
    C.newtScaleSetColors(c.c, C.int(empty), C.int(full))
}

