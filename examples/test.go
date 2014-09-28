package main

import "github.com/aklyachkin/go-newt"
import "fmt"
import "strconv"

func helpCallback(c *newt.Component, tag []byte) {
    //newt.WinMessage("Help", "Ok", "Here should be tag")
    fmt.Println("Hello world")
}

func main() {
    newt.Init()
    newt.Cls()

    defer newt.Finished()

    //var hcb *newt.Callback
    //hcb = (*newt.Callback)(&helpCallback)

    // newt.SetSuspendCallback(suspendCallback)
    //newt.SetHelpCallback(hcb)

    newt.DrawRootText(0, 0, "Newt test program")
    newt.PushHelpLine("")
    newt.DrawRootText(-50, 0, "More root text")

    newt.OpenWindow(2, 2, 30, 10, "first window")
    newt.OpenWindow(10, 5, 65, 16, "window 2")

    f := newt.Form(nil, "This is some help text", 0)

    chklist := newt.Form(nil, "", 0)

    b1 := newt.Button(3, 1, "Exit")
    b2 := newt.Button(18, 1, "Update")
    r1 := newt.Radiobutton(20, 10, "Choice 1", 0, nil)
    r2 := newt.Radiobutton(20, 11, "Choice 2", 0, &r1)
    r3 := newt.Radiobutton(20, 12, "Choice 3", 0, &r2)

    rsf := newt.Form(nil, "", 0)
    newt.FormAddComponents(rsf, r1, r2, r3)
    newt.FormSetBackground(rsf, newt.COLORSET_CHECKBOX)

    results := make([]newt.ResultStr, 10)
    cs := make([]newt.Component, 10)
    for i := range cs {
        buf := fmt.Sprintf("Check %d", i)
        cs[i] = newt.Checkbox(3, 10 + i, buf, " ", "", &(results[i]))
        newt.FormAddComponent(chklist, cs[i])
    }

    l1 := newt.Label(3, 6, "Scale:")
    l2 := newt.Label(3, 7, "Scrolls:")
    l3 := newt.Label(3, 8, "Hidden:")
    var scaleVal, enr2, enr3 newt.ResultStr
    e1 := newt.Entry(12, 6, "", 20, &scaleVal, 0)
    e2 := newt.Entry(12, 7, "Default", 20, &enr2, newt.FLAG_SCROLL)
    e3 := newt.Entry(12, 8, "", 20, &enr3, newt.FLAG_PASSWORD)

    /*
    cbis[0].state = &results[0];
    cbis[0].en = e1;
    newtComponentAddCallback(cs[0], disableCallback, &cbis[0]);
    */

    scale := newt.Scale(3, 14, 32, 100)

    newt.FormSetHeight(chklist, 3)
    newt.FormAddComponents(f, b1, b2, l1, l2, l3, e1, e2, e3, chklist)
    newt.FormAddComponents(f, rsf, scale)

    lb := newt.Listbox(45, 1, 6, newt.FLAG_MULTIPLE | newt.FLAG_BORDER | newt.FLAG_SCROLL | newt.FLAG_SHOWCURSOR)
    newt.ListboxAppendEntry(lb, "First", 1)
    newt.ListboxAppendEntry(lb, "Second", 2)
    newt.ListboxAppendEntry(lb, "Third", 3)
    newt.ListboxAppendEntry(lb, "Forth", 4)
    newt.ListboxAppendEntry(lb, "Sixth", 6)
    newt.ListboxAppendEntry(lb, "Seventh", 7)
    newt.ListboxAppendEntry(lb, "Eighth", 8)
    newt.ListboxAppendEntry(lb, "Ninth", 9)
    newt.ListboxAppendEntry(lb, "Tenth", 10)

    newt.ListboxInsertEntry(lb, "Fifth", 5, 4)
    newt.ListboxInsertEntry(lb, "Eleventh", 11, 10)
    newt.ListboxDeleteEntry(lb, 11)

    timeLabel := newt.Label(45, 8, "Spinner: -")

    t := newt.Textbox(45, 10, 17, 5, newt.FLAG_WRAP)
    newt.TextboxSetText(t, "This is some text. Does it look ok?\nThis should be alone.\nThis shouldn't be printed")

    newt.FormAddComponents(f, lb, timeLabel, t)
    newt.Refresh()
    newt.FormSetTimer(f, 200)

    spinner := "-\\|/"
    i := 0
    es := newt.FormRun(f)
    c := es.Component()
    for !c.Equals(&b1) {

        switch es.Reason() {
        case newt.EXIT_COMPONENT:
            if c.Equals(&b2) {
                sv, err := strconv.Atoi(scaleVal.String())
                if err == nil {
                    newt.ScaleSet(scale, uint64(sv))
                }
                newt.Refresh()
            }
        case newt.EXIT_TIMER:
            i++
            if i >= len(spinner) {
                i = 0
            }
            newt.LabelSetText(timeLabel, "Spinner: " + string(spinner[i]))
        }
        es = newt.FormRun(f)
        c = es.Component()
    }

    numsel, selectedList := newt.ListboxGetSelection(lb)

    newt.PopWindow()
    newt.PopWindow()
    newt.Finished()

    fmt.Println("got string 1: ", scaleVal.String())
    fmt.Println("got string 2: ", enr2.String())
    fmt.Println("got string 3: ", enr3.String())

    newt.FormDestroy(f)

    fmt.Println("Selected listbox items: ", numsel)
    if selectedList != nil {
        for i,v := range selectedList {
            fmt.Println("i: ", i, "\tv: ", v)
        }
    }
    // free(selectedList)
}
