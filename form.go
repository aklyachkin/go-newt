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
    c.ct = CMP_FORM
    return c
}

func FormSetTimer(form Component, millisecs int) {
    if form.ct == CMP_FORM {
        C.newtFormSetTimer(form.c, C.int(millisecs))
    }
}

func FormWatchFd(form Component, fd, fdFlags int) {
    if form.ct == CMP_FORM {
        C.newtFormWatchFd(form.c, C.int(fd), C.int(fdFlags))
    }
}

func FormSetSize(c Component) {
    if c.ct == CMP_FORM {
        C.newtFormSetSize(c.c)
    }
}

func FormGetCurrent(form Component) Component {
    if form.ct == CMP_FORM {
        var c Component
        c.c = C.newtFormGetCurrent(form.c)
        c.t = GRID_COMPONENT
        // c.ct == ?????
        return c
    } else {
        var c Component
        return c
    }
}

func FormSetBackground(c Component, color int) {
    if c.ct == CMP_FORM {
        C.newtFormSetBackground(c.c, C.int(color))
    }
}

func FormSetCurrent(form, c Component) {
    if form.ct == CMP_FORM {
        C.newtFormSetCurrent(form.c, c.c)
    }
}

func FormAddComponent(form, c Component) {
    if form.ct == CMP_FORM {
        C.newtFormAddComponent(form.c, c.c)
    }
}

func FormAddComponents(form Component, args ...Component) {
    if form.ct == CMP_FORM {
        for _, v := range args {
            FormAddComponent(form, v)
        }
    }
}

func FormSetHeight(c Component, height int) {
    if c.ct == CMP_FORM {
        C.newtFormSetHeight(c.c, C.int(height))
    }
}

func FormSetWidth(c Component, width int) {
    if c.ct == CMP_FORM {
        C.newtFormSetWidth(c.c, C.int(width))
    }
}

func RunForm(form Component) Component {
    if form.ct == CMP_FORM {
        var c Component
        c.c = C.newtRunForm(form.c)
        c.t = GRID_COMPONENT
        // c.ct == ????
        return c
    } else {
        var c Component
        return c
    }
}

func FormRun(form Component) ExitStruct {
    var es ExitStruct
    if form.ct == CMP_FORM {
        C.newtFormRun(form.c, &es.es)
    }
    return es
}

func DrawForm(form Component) {
    if form.ct == CMP_FORM {
        C.newtDrawForm(form.c)
    }
}

func FormAddHotKey(c Component, key int) {
    if c.ct == CMP_FORM {
        C.newtFormAddHotKey(c.c, C.int(key))
    }
}

func FormGetScrollPosition(c Component) int {
    if c.ct == CMP_FORM {
        return int(C.newtFormGetScrollPosition(c.c))
    } else {
        return -1
    }
}

func FormSetScrollPosition(c Component, position int) {
    if c.ct == CMP_FORM {
        C.newtFormSetScrollPosition(c.c, C.int(position))
    }
}

func FormDestroy(c Component) {
    if c.ct == CMP_FORM {
        C.newtFormDestroy(c.c)
    }
}
