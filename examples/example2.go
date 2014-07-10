package main

import "github.com/aklyachkin/go-newt"

func main() {
    newt.Init()
    newt.Cls()

    newt.OpenWindow(10, 5, 40, 6, "Button sample")

    b1 := newt.Button(10, 1, "Ok")
    b2 := newt.CompactButton(22, 2, "Cancel")
    form := newt.Form(nil, nil, 0)
    newt.FormAddComponents(form, b1, b2)

    newt.RunForm(form)

    newt.FormDestroy(form)
    newt.Finished()
}
