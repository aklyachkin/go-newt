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

func Entry(left, top int, initialValue string, width int, result *ResultStr, flags int) Component {
    var c Component
    iv := C.CString(initialValue)
    defer C.free(unsafe.Pointer(iv))
    c.c = C.newtEntry(C.int(left), C.int(top), iv, C.int(width), (**C.char)(&result.value), C.int(flags))
    return c
}

func EntrySet(c Component, value string, cursorAtEnd int) {
    t := C.CString(value)
    defer C.free(unsafe.Pointer(t))
    C.newtEntrySet(c.c, t, C.int(cursorAtEnd))
}

func EntrySetFilter(c Component, filter uintptr, data []byte) {
    panic("not implemented")
}

func EntryGetValue(c Component) string {
    r := C.newtEntryGetValue(c.c)
    defer C.free(unsafe.Pointer(r))
    return C.GoString(r)
}

func EntrySetFlags(c Component, flags int, sense uint32) {
    C.newtEntrySetFlags(c.c, C.int(flags), sense)
}

func EntrySetColors(c Component, normal, disabled int) {
    C.newtEntrySetColors(c.c, C.int(normal), C.int(disabled))
}

func EntryGetCursorPosition(c Component) int {
    return int(C.newtEntryGetCursorPosition(c.c))
}

func EntrySetCursorPosition(c Component, position int) {
    C.newtEntrySetCursorPosition(c.c, C.int(position))
}

