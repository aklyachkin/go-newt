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

func Form(vertBar *Component, helpTag string, flags int) Component {
    var c Component
    t := C.CString(helpTag)
    defer C.free(unsafe.Pointer(t))
    if vertBar != nil {
        c.c = C.newtForm(vertBar.c, unsafe.Pointer(t), C.int(flags))
    } else {
        c.c = C.newtForm(nil, unsafe.Pointer(t), C.int(flags))
    }
    c.t = GRID_COMPONENT
    return c
}

func FormSetTimer(form Component, millisecs int) {
    C.newtFormSetTimer(form.c, C.int(millisecs))
}

func FormWatchFd(form Component, fd, fdFlags int) {
    C.newtFormWatchFd(form.c, C.int(fd), C.int(fdFlags))
}

func FormSetSize(c Component) {
    C.newtFormSetSize(c.c)
}

func FormGetCurrent(form Component) Component {
    var c Component
    c.c = C.newtFormGetCurrent(form.c)
    c.t = GRID_COMPONENT
    return c
}

func FormSetBackground(c Component, color int) {
    C.newtFormSetBackground(c.c, C.int(color))
}

func FormSetCurrent(form, c Component) {
    C.newtFormSetCurrent(form.c, c.c)
}

func FormAddComponent(form, c Component) {
    C.newtFormAddComponent(form.c, c.c)
}

func FormAddComponents(form Component, args ...Component) {
    for _, v := range args {
        FormAddComponent(form, v)
    }
    // C.newtFormAddComponents(form.c, args)
}

func FormSetHeight(c Component, height int) {
    C.newtFormSetHeight(c.c, C.int(height))
}

func FormSetWidth(c Component, width int) {
    C.newtFormSetWidth(c.c, C.int(width))
}

func RunForm(form Component) Component {
    var c Component
    c.c = C.newtRunForm(form.c)
    c.t = GRID_COMPONENT
    return c
}

func FormRun(form Component) ExitStruct {
    var es ExitStruct
    C.newtFormRun(form.c, &es.es)
    return es
}

func DrawForm(form Component) {
    C.newtDrawForm(form.c)
}

func FormAddHotKey(c Component, key int) {
    C.newtFormAddHotKey(c.c, C.int(key))
}

func FormGetScrollPosition(c Component) int {
    return int(C.newtFormGetScrollPosition(c.c))
}

func FormSetScrollPosition(c Component, position int) {
    C.newtFormSetScrollPosition(c.c, C.int(position))
}

func FormDestroy(c Component) {
    C.newtFormDestroy(c.c)
}
