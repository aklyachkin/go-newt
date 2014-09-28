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
    c.t = GRID_COMPONENT
    c.ct = CMP_CHECKBOX
    return c
}

func CheckboxGetValue(c Component) int {
    if c.ct == CMP_CHECKBOX {
        return int(C.newtCheckboxGetValue(c.c))
    } else {
        return -1
    }
}

func CheckboxSetValue(c Component, value int) {
    if c.ct == CMP_CHECKBOX {
        C.newtCheckboxSetValue(c.c, C.char(value))
    }
}

func CheckboxSetFlags(c Component, flags int, sense uint32) {
    if c.ct == CMP_CHECKBOX {
        C.newtCheckboxSetFlags(c.c, C.int(flags), sense)
    }
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
    c.t = GRID_COMPONENT
    c.ct = CMP_RADIOBUTTON
    return c
}

func RadioGetCurrent(setMember Component) Component {
    if setMember.ct == CMP_RADIOBUTTON {
        var c Component
        c.c = C.newtRadioGetCurrent(setMember.c)
        c.t = GRID_COMPONENT
        c.ct = CMP_RADIOBUTTON
        return c
    } else {
        var c Component
        return c
    }
}

func RadioSetCurrent(setMember Component) {
    if setMember.ct == CMP_RADIOBUTTON {
        C.newtRadioSetCurrent(setMember.c)
    }
}

