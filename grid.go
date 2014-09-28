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

func CreateGrid(cols, rows int) Component {
    var c Component
    c.g = C.newtCreateGrid(C.int(cols), C.int(rows))
    c.t = GRID_SUBGRID
    return c
}

func stackem(isVert, close1 bool, comps ...Component) Component {
    num := len(comps)
    var grid Component
    if isVert {
        grid = CreateGrid(1, num)
    } else {
        grid = CreateGrid(num, 1)
    }

    for i, v := range comps {
        var row, col, padLeft, padRight int
        if isVert {
            row = i
            col = 0
            padLeft = 0
            padRight = 0
            if !close1 && i != 0 {
                padRight = 1
            }
        } else {
            row = 0
            col = i
            padLeft = 0
            padRight = 0
            if !close1 && i != 0 {
                padLeft = 1
            }
        }
        GridSetField(grid, col, row, v.t, v, padLeft, padRight, 0, 0, 0, 0)
    }

    return grid
}

func GridVStacked(comps ...Component) Component {
    return stackem(true, false, comps...)
}

func GridVCloseStacked(comps ...Component) Component {
    return stackem(true, true, comps...)
}

func GridHStacked(comps ...Component) Component {
    return stackem(false, false, comps...)
}

func GridHCloseStacked(comps ...Component) Component {
    return stackem(false, true, comps...)
}

func GridBasicWindow(text, middle, buttons Component) Component {
    var c Component
    c.g = C.newtGridBasicWindow(text.c, middle.g, buttons.g)
    c.t = GRID_SUBGRID
    return c
}

func GridSimpleWindow(text, middle, buttons Component) Component {
    var c Component
    c.g = C.newtGridSimpleWindow(text.c, middle.c, buttons.g)
    c.t = GRID_SUBGRID
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

