package newt

/*
#cgo CFLAGS: -I/opt/local/include
#cgo LDFLAGS: -L/opt/local/lib -lnewt

#include <stdio.h>
#include <stdlib.h>
#include <newt.h>

*/
import "C"

type TButtonbar struct {
    text string
    button Component
}

func ButtonBar(bbar ...*TButtonbar) Component {
    grid := CreateGrid(len(bbar), 1)

    for i, _ := range bbar {
        bbar[i].button = Button(-1, -1, bbar[i].text)
        GridSetField(grid, i, 0, GRID_COMPONENT, bbar[i].button, 1, 0, 0, 0, 0, 0)
    }

    return grid
}

