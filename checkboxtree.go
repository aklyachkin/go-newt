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

func CheckboxTree(left, top, height, flags int) Component {
    var c Component
    c.c = C.newtCheckboxTree(C.int(left), C.int(top), C.int(height), C.int(flags))
    c.t = GRID_COMPONENT
    return c
}

func CheckboxTreeMulti(left, top, height int, seq string, flags int) Component {
    var c Component
    s := C.CString(seq)
    defer C.free(unsafe.Pointer(s))
    c.c = C.newtCheckboxTreeMulti(C.int(left), C.int(top), C.int(height), s, C.int(flags))
    c.t = GRID_COMPONENT
    return c
}

func CheckboxTreeGetSelection(c Component) (int, *uintptr) {
    var numitems C.int
    items := C.newtCheckboxTreeGetSelection(c.c, &numitems)
    i := uintptr(*items)
    return int(numitems), &i
}

func CheckboxTreeGetCurrent(c Component) uintptr {
    return uintptr(C.newtCheckboxTreeGetCurrent(c.c))
}

func CheckboxTreeSetCurrent(c Component, item uintptr) {
    C.newtCheckboxTreeSetCurrent(c.c, unsafe.Pointer(item))
}

func CheckboxTreeGetMultiSelection(c Component, seqnum int) (int, *uintptr) {
    var numitems C.int
    items := C.newtCheckboxTreeGetMultiSelection(c.c, &numitems, C.char(seqnum))
    i := uintptr(*items)
    return int(numitems), &i
}

func CheckboxTreeAddItem(c Component, text string, data uintptr, flags int, index int) int {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    return int(C.newtCheckboxTreeAddItem(c.c, t, unsafe.Pointer(data), C.int(flags), C.int(index)))
}

func CheckboxTreeAddArray(c Component, text string, data uintptr, flags int, indexes []int) int {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    var i []C.int
    for a := range indexes {
        i = append(i, C.int(a))
    }
    return int(C.newtCheckboxTreeAddArray(c.c, t, unsafe.Pointer(data), C.int(flags), &i[0]))
}

func CheckboxTreeFindItem(c Component, data uintptr) int {
    r := C.newtCheckboxTreeFindItem(c.c, unsafe.Pointer(data))
    return int(*r)
}

func CheckboxTreeSetEntry(c Component, data uintptr, text string) {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    C.newtCheckboxTreeSetEntry(c.c, unsafe.Pointer(data), t)
}

func CheckboxTreeSetWidth(c Component, width int) {
    C.newtCheckboxTreeSetWidth(c.c, C.int(width))
}

func CheckboxTreeGetEntryValue(c Component) (int, uintptr) {
    var p unsafe.Pointer
    r := C.newtCheckboxTreeGetEntryValue(c.c, p)
    return int(r), uintptr(p)
}

func CheckboxTreeSetEntryValue(c Component, data uintptr, value int) {
    C.newtCheckboxTreeSetEntryValue(c.c, unsafe.Pointer(data), C.char(value))
}

