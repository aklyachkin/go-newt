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

func Entry(left, top int, initialValue string, width int, result *ResultStr, flags int) Component {
    var c Component
    iv := C.CString(initialValue)
    defer C.free(unsafe.Pointer(iv))
    c.c = C.newtEntry(C.int(left), C.int(top), iv, C.int(width), (**C.char)(&result.value), C.int(flags))
    c.t = GRID_COMPONENT
    c.ct = CMP_ENTRY
    return c
}

func EntrySet(c Component, value string, cursorAtEnd int) {
    if c.ct == CMP_ENTRY {
        t := C.CString(value)
        defer C.free(unsafe.Pointer(t))
        C.newtEntrySet(c.c, t, C.int(cursorAtEnd))
    }
}

func EntrySetFilter(c Component, filter uintptr, data []byte) {
    panic("not implemented")
}

func EntryGetValue(c Component) string {
    if c.ct == CMP_ENTRY {
        r := C.newtEntryGetValue(c.c)
        defer C.free(unsafe.Pointer(r))
        return C.GoString(r)
    } else {
        return ""
    }
}

func EntrySetFlags(c Component, flags int, sense uint32) {
    if c.ct == CMP_ENTRY {
        C.newtEntrySetFlags(c.c, C.int(flags), sense)
    }
}

func EntrySetColors(c Component, normal, disabled int) {
    if c.ct == CMP_ENTRY {
        C.newtEntrySetColors(c.c, C.int(normal), C.int(disabled))
    }
}

func EntryGetCursorPosition(c Component) int {
    if c.ct == CMP_ENTRY {
        return int(C.newtEntryGetCursorPosition(c.c))
    } else {
        return -1
    }
}

func EntrySetCursorPosition(c Component, position int) {
    if c.ct == CMP_ENTRY {
        C.newtEntrySetCursorPosition(c.c, C.int(position))
    }
}

