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

func TextboxReflowed(left, top int, text string, width, flexDown, flexUp, flags int) Component {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    var c Component
    c.c = C.newtTextboxReflowed(C.int(left), C.int(top), t, C.int(width), C.int(flexDown), C.int(flexUp), C.int(flags))
    c.t = GRID_COMPONENT
    return c
}

func Textbox(left, top, width, height, flags int) Component {
    var c Component
    c.c = C.newtTextbox(C.int(left), C.int(top), C.int(width), C.int(height), C.int(flags))
    c.t = GRID_COMPONENT
    return c
}

func TextboxSetText(c Component, text string) {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    C.newtTextboxSetText(c.c, t)
}

func TextboxSetHeight(c Component, height int) {
    C.newtTextboxSetHeight(c.c, C.int(height))
}

func TextboxGetNumLines(c Component) int {
    return int(C.newtTextboxGetNumLines(c.c))
}

func TextboxSetColors(c Component, normal, active int) {
    C.newtTextboxSetColors(c.c, C.int(normal), C.int(active))
}

func ReflowText(text string, width, flexDown, flexUp int) (string, int, int) {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    var actualWidth, actualHeight C.int
    r := C.newtReflowText(t, C.int(width), C.int(flexDown), C.int(flexUp), &actualWidth, &actualHeight)
    defer C.free(unsafe.Pointer(r))
    s := C.GoString(r)
    return s, int(actualWidth), int(actualHeight)
}

