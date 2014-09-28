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

func CheckboxTree(left, top, height, flags int) Component {
    var c Component
    c.c = C.newtCheckboxTree(C.int(left), C.int(top), C.int(height), C.int(flags))
    c.t = GRID_COMPONENT
    c.ct = CMP_CHECKBOXTREE
    return c
}

func CheckboxTreeMulti(left, top, height int, seq string, flags int) Component {
    var c Component
    s := C.CString(seq)
    defer C.free(unsafe.Pointer(s))
    c.c = C.newtCheckboxTreeMulti(C.int(left), C.int(top), C.int(height), s, C.int(flags))
    c.t = GRID_COMPONENT
    c.ct = CMP_CHECKBOXTREE
    return c
}

func CheckboxTreeGetSelection(c Component) (int, *uintptr) {
    if c.ct == CMP_CHECKBOXTREE {
        var numitems C.int
        items := C.newtCheckboxTreeGetSelection(c.c, &numitems)
        i := uintptr(*items)
        return int(numitems), &i
    } else {
        return -1, nil
    }
}

func CheckboxTreeGetCurrent(c Component) uintptr {
    if c.ct == CMP_CHECKBOXTREE {
        return uintptr(C.newtCheckboxTreeGetCurrent(c.c))
    } else {
        var i uintptr
        return i
    }
}

func CheckboxTreeSetCurrent(c Component, item uintptr) {
    if c.ct == CMP_CHECKBOXTREE {
        C.newtCheckboxTreeSetCurrent(c.c, unsafe.Pointer(item))
    }
}

func CheckboxTreeGetMultiSelection(c Component, seqnum int) (int, *uintptr) {
    if c.ct == CMP_CHECKBOXTREE {
        var numitems C.int
        items := C.newtCheckboxTreeGetMultiSelection(c.c, &numitems, C.char(seqnum))
        i := uintptr(*items)
        return int(numitems), &i
    } else {
        return -1, nil
    }
}

func CheckboxTreeAddItem(c Component, text string, data uintptr, flags int, index int) int {
    if c.ct == CMP_CHECKBOXTREE {
        t := C.CString(text)
        defer C.free(unsafe.Pointer(t))
        i := C.newtCheckboxTreeAddItem(c.c, t, unsafe.Pointer(data), C.int(flags), C.int(index))
        return int(i)
    } else {
        return -1
    }
}

func CheckboxTreeAddArray(c Component, text string, data uintptr, flags int, indexes []int) int {
    if c.ct == CMP_CHECKBOXTREE {
        t := C.CString(text)
        defer C.free(unsafe.Pointer(t))
        var i []C.int
        for a := range indexes {
            i = append(i, C.int(a))
        }
        return int(C.newtCheckboxTreeAddArray(c.c, t, unsafe.Pointer(data), C.int(flags), &i[0]))
    } else {
        return -1
    }
}

func CheckboxTreeFindItem(c Component, data uintptr) int {
    if c.ct == CMP_CHECKBOXTREE {
        r := C.newtCheckboxTreeFindItem(c.c, unsafe.Pointer(data))
        return int(*r)
    } else {
        return -1
    }
}

func CheckboxTreeSetEntry(c Component, data uintptr, text string) {
    if c.ct == CMP_CHECKBOXTREE {
        t := C.CString(text)
        defer C.free(unsafe.Pointer(t))
        C.newtCheckboxTreeSetEntry(c.c, unsafe.Pointer(data), t)
    }
}

func CheckboxTreeSetWidth(c Component, width int) {
    if c.ct == CMP_CHECKBOXTREE {
        C.newtCheckboxTreeSetWidth(c.c, C.int(width))
    }
}

func CheckboxTreeGetEntryValue(c Component) (int, uintptr) {
    if c.ct == CMP_CHECKBOXTREE {
        var p unsafe.Pointer
        r := C.newtCheckboxTreeGetEntryValue(c.c, p)
        return int(r), uintptr(p)
    } else {
        var i uintptr
        return -1, i
    }
}

func CheckboxTreeSetEntryValue(c Component, data uintptr, value int) {
    if c.ct == CMP_CHECKBOXTREE {
        C.newtCheckboxTreeSetEntryValue(c.c, unsafe.Pointer(data), C.char(value))
    }
}

