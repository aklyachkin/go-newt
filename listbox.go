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
    "reflect"
)

func Listbox(left, top, height, flags int) Component {
    var c Component
    c.c = C.newtListbox(C.int(left), C.int(top), C.int(height), C.int(flags))
    c.t = GRID_COMPONENT
    return c
}

func ListboxGetCurrent(c Component) uintptr {
    data := uintptr(C.newtListboxGetCurrent(c.c))
    return data
}

func ListboxSetCurrent(c Component, num int) {
    C.newtListboxSetCurrent(c.c, C.int(num))
}

func ListboxSetCurrentByKey(c Component, key uintptr) {
    C.newtListboxSetCurrentByKey(c.c, unsafe.Pointer(key))
}

func ListboxSetEntry(c Component, num int, text string) {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    C.newtListboxSetEntry(c.c, C.int(num), t)
}

func ListboxSetWidth(c Component, width int) {
    C.newtListboxSetWidth(c.c, C.int(width))
}

func ListboxSetData(c Component, num int, data uintptr) {
    C.newtListboxSetData(c.c, C.int(num), unsafe.Pointer(data))
}

func ListboxAppendEntry(c Component, text string, data uintptr) int {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    return int(C.newtListboxAppendEntry(c.c, t, unsafe.Pointer(data)))
}

func ListboxAddEntry(c Component, text string, data uintptr) int {
    return ListboxAppendEntry(c, text, data)
}

func ListboxInsertEntry(c Component, text string, data, key uintptr) int {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    return int(C.newtListboxInsertEntry(c.c, t, unsafe.Pointer(data), unsafe.Pointer(key)))
}

func ListboxDeleteEntry(c Component, data uintptr) int {
    return int(C.newtListboxDeleteEntry(c.c, unsafe.Pointer(data)))
}

func ListboxClear(c Component) {
    C.newtListboxClear(c.c)
}

func ListboxGetEntry(c Component, num int) (string, uintptr) {
    var text *C.char
    var data unsafe.Pointer
    C.newtListboxGetEntry(c.c, C.int(num), &text, &data)
    return C.GoString(text), uintptr(data)
}

func ListboxGetSelection(c Component) (int, []uintptr) {
    var numitems C.int
    var i []uintptr
    items := C.newtListboxGetSelection(c.c, &numitems)
    if items != nil {
        hdr := reflect.SliceHeader{
            Data: uintptr(unsafe.Pointer(items)),
            Len: int(numitems),
            Cap: int(numitems),
        }
        i = *(*[]uintptr)(unsafe.Pointer(&hdr))
    }
    return int(numitems), i
}

func ListboxClearSelection(c Component) {
    C.newtListboxClearSelection(c.c)
}

func ListboxSelectItem(c Component, key uintptr, sense uint32) {
    C.newtListboxSelectItem(c.c, unsafe.Pointer(key), sense)
}

func ListboxItemCount(c Component) int {
    n := C.newtListboxItemCount(c.c)
    return int(n)
}

