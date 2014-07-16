package main

import "github.com/aklyachkin/go-newt"

import "fmt"

func main() {
    newt.Init()
    newt.Cls()

    r := newt.NewResultStr()

    b1 := newt.Checkbox(-1, -1, "A pretty long checkbox for testing", " ", "", r)
    b2 := newt.Button(-1, -1, "Another button")
    b3 := newt.Button(-1, -1, "But, but")
    b4 := newt.Button(-1, -1, "But what?")

    f := newt.Form(nil, "", 0)

    grid := newt.CreateGrid(2, 2)
    newt.GridSetField(grid, 0, 0, newt.GRID_COMPONENT, b1, 0, 0, 0, 0, newt.ANCHOR_RIGHT, 0)
    newt.GridSetField(grid, 0, 1, newt.GRID_COMPONENT, b2, 0, 0, 0, 0, 0, 0)
    newt.GridSetField(grid, 1, 0, newt.GRID_COMPONENT, b3, 0, 0, 0, 0, 0, 0)
    newt.GridSetField(grid, 1, 1, newt.GRID_COMPONENT, b4, 0, 0, 0, 0, 0, 0)

    newt.FormAddComponents(f, b1, b2, b3, b4)
    newt.GridWrappedWindow(grid, "first window")
    newt.GridFree(grid, 1)

    newt.RunForm(f)
    newt.FormDestroy(f)

    newt.PopWindow()

    flowedText, textWidth, textHeight := newt.ReflowText(
        "This is quite a bit of text. It is 40 " +
        "columns long, so some wrapping should be " +
        "done. Did you know that the quick, brown " +
        "fox jumped over the lazy dog?\n\n" +
        "In other news, it's pretty important that we " +
        "can properly force a line break.",
        40, 5, 5)
    t := newt.Textbox(-1, -1, textWidth, textHeight, newt.FLAG_WRAP)
    newt.TextboxSetText(t, flowedText)

    b1 = newt.Button(-1, -1, "Okay")
    b2 = newt.Button(-1, -1, "Cancel")

    grid = newt.CreateGrid(1, 2)
    subgrid := newt.CreateGrid(2, 1)

    newt.GridSetField(subgrid, 0, 0, newt.GRID_COMPONENT, b1, 0, 0, 0, 0, 0, 0)
    newt.GridSetField(subgrid, 1, 0, newt.GRID_COMPONENT, b2, 0, 0, 0, 0, 0, 0)

    newt.GridSetField(grid, 0, 0, newt.GRID_COMPONENT, t, 0, 0, 0, 1, 0, 0)
    newt.GridSetField(grid, 0, 1, newt.GRID_SUBGRID, subgrid, 0, 0, 0, 0, 0, newt.GRID_FLAG_GROWX)
    newt.GridWrappedWindow(grid, "another example")
    newt.GridDestroy(grid, 1)

    f = newt.Form(nil, "", 0)
    newt.FormAddComponents(f, b1, t, b2)
    newt.RunForm(f)

    newt.PopWindow()
    newt.FormDestroy(f)

    newt.WinMessage("Simple", "Ok", "This is a simple message window")
    newt.WinChoice("Simple", "Ok", "Cancel", "This is a simple choice window")

    textWidth = 0
    menuContents := []string{"One", "Two", "Three", "Four", "Five"}
    rc, textWidth := newt.WinMenu("Test Menu", "This is a simple involvation of the " +
        "newtWinMenu() call. It may or may not have a scrollbar, " +
        "depending on the need for one.", 50, 5, 5, 3, menuContents, "Ok", "Cancel")

    autoEntries := make([]newt.WinEntry, 4)
    autoEntries[0] = newt.NewWinEntry("An entry", 0)
    autoEntries[1] = newt.NewWinEntry("Another entry", 0)
    autoEntries[2] = newt.NewWinEntry("Third entry", 0)
    autoEntries[3] = newt.NewWinEntry("Forth entry", 0)

    rc1 := newt.WinEntries("Test newtWinEntries()", "This is a simple invovation of " +
        "newtWinEntries() call. It lets you get a lot of input " +
        "quite easily.", 50, 5, 5, 20, &autoEntries, "Ok", "Cancel")

    newt.Finished()
    fmt.Println("WinMenu():",rc, textWidth)
    fmt.Println("WinEntries():",rc1)

    for i, v := range autoEntries {
        fmt.Printf("%d) Value: %s\n", i, v.Value())
        v.Destroy()
    }
}
