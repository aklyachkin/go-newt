package main

import "github.com/aklyachkin/go-newt"
import "fmt"

func main() {
    newt.Init()
    newt.Cls()

    newt.OpenWindow(10, 5, 40, 11, "Checkboxes and Radio buttons")

    cbValue := newt.NewResultStr()

    checkbox := newt.Checkbox(1, 1, "A checkbox", " ", " *X", cbValue)

    rb := make([]newt.Component, 3)
    rb[0] = newt.Radiobutton(1, 3, "Choice 1", 1, nil)
    rb[1] = newt.Radiobutton(1, 4, "Choice 2", 0, &(rb[0]))
    rb[2] = newt.Radiobutton(1, 5, "Choice 3", 0, &(rb[1]))

    button := newt.Button(1, 7, "Ok")

    form := newt.Form(nil, nil, 0)
    newt.FormAddComponent(form, checkbox)
    for i := range rb {
        newt.FormAddComponent(form, rb[i])
    }
    newt.FormAddComponent(form, button)

    newt.RunForm(form)
    newt.Finished()

    for i := range rb {
        if newt.RadioGetCurrent(rb[0]) == rb[i] {
            fmt.Println("radio button picked: ", i)
        }
    }

    newt.FormDestroy(form)
    fmt.Println("Checkbox value: ", cbValue.String())
}
