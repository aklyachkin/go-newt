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

/* Components */
func Checkbox(left, top int, text string, defValue, seq string, result *ResultStr) Component {
    var c Component
    t := C.CString(text)
    var s *C.char
    if seq == "" {
        s = nil
    } else {
        s = C.CString(seq)
    }
    defer C.free(unsafe.Pointer(t))
    defer C.free(unsafe.Pointer(s))
    c.c = C.newtCheckbox(C.int(left), C.int(top), t, C.char(defValue[0]), s, result.value)
    return c
}

func CheckboxGetValue(c Component) int {
    a := C.newtCheckboxGetValue(c.c)
    return int(a)
}

func CheckboxSetValue(c Component, value int) {
    C.newtCheckboxSetValue(c.c, C.char(value))
}

func CheckboxSetFlags(c Component, flags int, sense uint32) {
    C.newtCheckboxSetFlags(c.c, C.int(flags), sense)
}

func Radiobutton(left, top int, text string, isDefault int, prevButton *Component) Component {
    var c Component
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    if prevButton == nil {
        c.c = C.newtRadiobutton(C.int(left), C.int(top), t, C.int(isDefault), nil)
    } else {
        c.c = C.newtRadiobutton(C.int(left), C.int(top), t, C.int(isDefault), prevButton.c)
    }
    return c
}

func RadioGetCurrent(setMember Component) Component {
    var c Component
    c.c = C.newtRadioGetCurrent(setMember.c)
    return c
}

func RadioSetCurrent(setMember Component) {
    C.newtRadioSetCurrent(setMember.c)
}

