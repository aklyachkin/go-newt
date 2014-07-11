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

func CreateGrid(cols, rows int) Component {
    var c Component
    c.g = C.newtCreateGrid(C.int(cols), C.int(rows))
    return c
}

func GridVStacked(typ uint32, what uintptr) Component {
    panic("not implemented")
}

func GridVCloseStacked(typ uint32, what uintptr) Component {
    panic("not implemented")
}

func GridHStacked(typ uint32, what uintptr) Component {
    panic("not implemented")
}

func GridHCloseStacked(typ uint32, what uintptr) Component {
    panic("not implemented")
}

func GridBasicWindow(text, middle, buttons Component) Component {
    var c Component
    c.g = C.newtGridBasicWindow(text.c, middle.g, buttons.g)
    return c
}

func GridSimpleWindow(text, middle, buttons Component) Component {
    var c Component
    c.g = C.newtGridSimpleWindow(text.c, middle.c, buttons.g)
    return c
}

func GridSetField(c Component, col, row int, typ uint32, val Component, padLeft, padTop, padRight, padBottom, anchor, flags int) {
    switch typ {
    case GRID_COMPONENT:
        C.newtGridSetField(c.g, C.int(col), C.int(row), typ, unsafe.Pointer(val.c), C.int(padLeft), C.int(padTop), C.int(padRight), C.int(padBottom), C.int(anchor), C.int(flags))
    case GRID_SUBGRID:
        C.newtGridSetField(c.g, C.int(col), C.int(row), typ, unsafe.Pointer(val.g), C.int(padLeft), C.int(padTop), C.int(padRight), C.int(padBottom), C.int(anchor), C.int(flags))
    case GRID_EMPTY:
        C.newtGridSetField(c.g, C.int(col), C.int(row), typ, nil, C.int(padLeft), C.int(padTop), C.int(padRight), C.int(padBottom), C.int(anchor), C.int(flags))
    }
}

func GridPlace(c Component, left, top int) {
    C.newtGridPlace(c.g, C.int(left), C.int(top))
}

func GridDestroy(c Component, recurse int) {
    C.newtGridFree(c.g, C.int(recurse))
}

func GridFree(c Component, recurse int) {
    C.newtGridFree(c.g, C.int(recurse))
}

func GridGetSize(c Component) (int, int) {
    var width, height C.int
    C.newtGridGetSize(c.g, &width, &height)
    return int(width), int(height)
}

func GridWrappedWindow(c Component, title string) {
    t := C.CString(title)
    defer C.free(unsafe.Pointer(t))
    C.newtGridWrappedWindow(c.g, t)
}

func GridWrappedWindowAt(c Component, title string, left, top int) {
    t := C.CString(title)
    defer C.free(unsafe.Pointer(t))
    C.newtGridWrappedWindowAt(c.g, t, C.int(left), C.int(top))
}

func GridAddComponentsToForm(grid, form Component, recurse int) {
    C.newtGridAddComponentsToForm(grid.g, form.c, C.int(recurse))
}

