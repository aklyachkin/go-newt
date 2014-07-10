package main

import "fmt"
import "github.com/aklyachkin/go-newt"

func main() {
    newt.Init()
    newt.Cls()

    newt.OpenWindow(10, 5, 40, 8, "Entry and Label Sample")

    entryValue := newt.NewResultStr()
    label := newt.Label(1, 1, "Enter a string:")
    entry := newt.Entry(16, 1, "sample", 20, &entryValue, newt.FLAG_SCROLL | newt.FLAG_RETURNEXIT)
    button := newt.Button(17, 3, "Ok")
    form := newt.Form(nil, nil, 0)
    newt.FormAddComponents(form, label, entry, button)

    newt.RunForm(form)

    newt.Finished()
    fmt.Println("Final string was: ", entryValue.String())
    newt.FormDestroy(form)
}
