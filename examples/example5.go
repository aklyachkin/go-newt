package main

import "github.com/aklyachkin/go-newt"

func main() {
    message := "This is a pretty long message. It will be displayed " +
        "in a newt textbox, and illustrates how to construct " +
        "a textbox from arbitrary text which may not have " +
        "very good line breaks.\n\n" +
        "Notice how literal \\n characters are respected, and " +
        "may be used to force line breaks and blank lines.";

    newt.Init()
    newt.Cls()

    text := newt.TextboxReflowed(1, 1, message, 30, 5, 5, 0)
    button := newt.Button(12, newt.TextboxGetNumLines(text) + 2, "Ok")

    newt.OpenWindow(10, 5, 37, uint(newt.TextboxGetNumLines(text) + 7), "Textboxes")

    form := newt.Form(nil, "", 0)
    newt.FormAddComponents(form, text, button)

    newt.RunForm(form)
    newt.FormDestroy(form)
    newt.Finished()
}
