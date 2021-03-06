package newt

/*
#cgo pkg-config: libnewt
#cgo LDFLAGS: -lnewt

#include <stdio.h>
#include <stdlib.h>
#include <newt.h>
*/
import "C"

func VerticalScrollbar(left, top, height, normalColorset, thumbColorset int) Component {
    var c Component
    c.c = C.newtVerticalScrollbar(C.int(left), C.int(top), C.int(height), C.int(normalColorset), C.int(thumbColorset))
    c.t = GRID_COMPONENT
    return c
}

func ScrollbarSet(c Component, where, total int) {
    C.newtScrollbarSet(c.c, C.int(where), C.int(total))
}

func ScrollbarSetColors(c Component, normal, thumb int) {
    C.newtScrollbarSetColors(c.c, C.int(normal), C.int(thumb))
}

