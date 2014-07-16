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
    for i, _ := range(buttons) {
        if result == buttons[i] {
            break
        }
        rc++
    }
    if rc == len(buttons) { rc = 0 } else { rc++ }

    FormDestroy(form)
    PopWindow()

    return rc, int(listItem)
}

func WinEntries(title, text string, suggestedWidth, flexDown, flexUp, dataWidth int, items *[]WinEntry, button1 ...string) int {
    textw := TextboxReflowed(-1, -1, text, suggestedWidth, flexDown, flexUp, 0)
    _ = textw
    buttons := make([]Component, len(button1))
    buttonBar := CreateGrid(len(button1), 1)

    for i, v := range button1 {
        j := 1
        if i == 0 { j = 0 }
        buttons[i] = Button(-1, -1, v)
        GridSetField(buttonBar, i, 0, GRID_COMPONENT, buttons[i], j, 0, 0, 0, 0, 0)
    }

    subgrid := CreateGrid(2, len(*items))
    a := make([]ResultStr, len(*items))
    for i, _ := range *items {
        l2 := Label(-1, -1, (*items)[i].Text())
        e2 := Entry(-1, -1, (*items)[i].Value(), dataWidth, &a[i], (*items)[i].Flags())
        GridSetField(subgrid, 0, i, GRID_COMPONENT, l2, 0, 0, 0, 0, ANCHOR_LEFT, 0)
        GridSetField(subgrid, 1, i, GRID_COMPONENT, e2, 1, 0, 0, 0, 0, 0)
    }

    grid := CreateGrid(1, 3)
    form := Form(nil, "", 0)
    GridSetField(grid, 0, 0, GRID_COMPONENT, textw, 0, 0, 0, 0, ANCHOR_LEFT, 0)
    GridSetField(grid, 0, 1, GRID_SUBGRID, subgrid, 0, 1, 0, 0, 0, 0)
    GridSetField(grid, 0, 2, GRID_SUBGRID, buttonBar, 0, 1, 0, 0, 0, GRID_FLAG_GROWX)
    GridAddComponentsToForm(grid, form, 1)
    GridWrappedWindow(grid, title)
    GridFree(grid, 1)

    result := RunForm(form)

    rc := 0
    for i, _ := range(buttons) {
        if result == buttons[i] {
            break
        }
        rc++
    }
    if rc == len(buttons) { rc = 0 } else { rc++ }

    for i, _ := range a {
        (*items)[i].dv.Set(a[i].String())
    }

    FormDestroy(form)
    PopWindow()

    return rc
}
