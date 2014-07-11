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

func WinMessage(title, buttonText, text string) {
    t1 := C.CString(title)
    b1 := C.CString(buttonText)
    t2 := C.CString(text)
    defer func() {
        C.free(unsafe.Pointer(t1))
        C.free(unsafe.Pointer(b1))
        C.free(unsafe.Pointer(t2))
    }()
    C.newtWinMessage(t1, b1, t2)
}

func WinMessageV(title, buttonText, text string) {
    panic("not implemented")
}

func WinChoice(title, button1, button2, text string) int {
    t1 := C.CString(title)
    b1 := C.CString(button1)
    b2 := C.CString(button2)
    t2 := C.CString(text)
    defer func() {
        C.free(unsafe.Pointer(t1))
        C.free(unsafe.Pointer(b1))
        C.free(unsafe.Pointer(b2))
        C.free(unsafe.Pointer(t2))
    }()
    return int(C.newtWinChoice(t1, b1, b2, t2))
}

func WinTernary(title, button1, button2, button3, message string) int {
    t1 := C.CString(title)
    b1 := C.CString(button1)
    b2 := C.CString(button2)
    b3 := C.CString(button3)
    m1 := C.CString(message)
    defer func() {
        C.free(unsafe.Pointer(t1))
        C.free(unsafe.Pointer(b1))
        C.free(unsafe.Pointer(b2))
        C.free(unsafe.Pointer(b3))
        C.free(unsafe.Pointer(m1))
    }()
    return int(C.newtWinTernary(t1, b1, b2, b3, m1))
}

func WinMenu(title, text string, suggestedWidth, flexDown, flexUp, maxListHeight int, items []string, button1 string) (int, int) {
    t1 := C.CString(title)
    t2 := C.CString(text)
    b1 := C.CString(button1)
    defer func() {
        C.free(unsafe.Pointer(t1))
        C.free(unsafe.Pointer(t2))
        C.free(unsafe.Pointer(b1))
    }()
    Citems := make([]*C.char, len(items))
    for i, v := range(items) {
        Citems[i] = C.CString(v)
    }
    defer func() {
        for i := range(Citems) {
            C.free(unsafe.Pointer(Citems[i]))
        }
    }()

    var listItem C.int
    // r := C.newtWinMenu(t1, t2, C.int(suggestedWidth), C.int(flexDown), C.int(flexUp), C.int(maxListHeight), (**C.char)(unsafe.Pointer(&Citems[0])), &listItem, b1, nil)
    return int(r), int(listItem)
}
