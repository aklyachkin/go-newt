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

func WinMenu(title, text string, suggestedWidth, flexDown, flexUp, maxListHeight int, items []string, button1 ...string) (int, int) {
    textbox := TextboxReflowed(-1, -1, text, suggestedWidth, flexDown, flexUp, 0)
    if len(items) < maxListHeight {
        maxListHeight = len(items)
    }
    needScroll := 0
    if len(items) > maxListHeight {
        needScroll = FLAG_SCROLL
    }

    listbox := Listbox(-1, -1, maxListHeight, needScroll | FLAG_RETURNEXIT);
    for i, v := range items {
        u := uint(i)
        ListboxAddEntry(listbox, v, uintptr(unsafe.Pointer(&u)))
    }

    var listItem int
    ListboxSetCurrent(listbox, listItem)

    buttons := make([]Component, len(button1))
    buttonBar := CreateGrid(len(button1), 1)
    for i, v := range button1 {
        j := 1
        if i == 0 { j = 0 }
        buttons[i] = Button(-1, -1, v)
        GridSetField(buttonBar, i, 0, GRID_COMPONENT, buttons[i], j, 0, 0, 0, 0, 0)
    }

    grid := GridSimpleWindow(textbox, listbox, buttonBar)
    GridWrappedWindow(grid, title)

    form := Form(nil, "", 0)
    GridAddComponentsToForm(grid, form, 1)
    GridFree(grid, 1)

    result := RunForm(form)
    listItemPtr := ListboxGetCurrent(listbox)
    listItemC := *(*C.int)(unsafe.Pointer(listItemPtr))
    listItem = int(listItemC)

    rc := 0
    for _, v := range(buttons) {
        if result == v {
            break
        }
        rc++
    }
    if rc == len(buttons) { rc = 0 } else { rc++ }

    FormDestroy(form)
    PopWindow()

    return rc, int(listItem)

    //t1 := C.CString(title)
    //t2 := C.CString(text)
    //b1 := C.CString(button1)
    //defer func() {
    //    C.free(unsafe.Pointer(t1))
    //    C.free(unsafe.Pointer(t2))
    //    C.free(unsafe.Pointer(b1))
    //}()
    //Citems := make([]*C.char, len(items))
    //for i, v := range(items) {
    //    Citems[i] = C.CString(v)
   // }
    //defer func() {
    //    for i := range(Citems) {
    //        C.free(unsafe.Pointer(Citems[i]))
    //    }
    //}()

    //panic("not implemented")
    //var listItem C.int
    // r := C.newtWinMenu(t1, t2, C.int(suggestedWidth), C.int(flexDown), C.int(flexUp), C.int(maxListHeight), (**C.char)(unsafe.Pointer(&Citems[0])), &listItem, b1, nil)
    // return int(r), int(listItem)
}
