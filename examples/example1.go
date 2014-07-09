package main

import "time"
import "github.com/aklyachkin/go-newt"

func main() {
    newt.Init()
    newt.Cls()
    newt.DrawRootText(0, 0, "Some root text")
    newt.DrawRootText(-25, -2, "Root text in the other corner")
    newt.PushHelpLine("")
    newt.Refresh()
    time.Sleep(1 * time.Second)
    newt.PushHelpLine("A help line")
    newt.Refresh()
    time.Sleep(1 * time.Second)
    newt.PopHelpLine()
    newt.Refresh()
    time.Sleep(1 * time.Second)
    newt.Finished()
}
