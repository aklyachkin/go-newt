package newt

/*
#cgo CFLAGS: -I/opt/local/include
#cgo LDFLAGS: -L/opt/local/lib -lnewt

#include <stdio.h>
#include <stdlib.h>
#include <newt.h>

newtCallback helpCallback;

void n_setSuspendCallback(void *cb, void *data)
{
    newtSetSuspendCallback(cb, data);
}

void n_HelpCallback(newtComponent c, void *data)
{
    printf("n_setHelpCallback()\n");
    if (helpCallback != NULL) {
        helpCallback(c, data);
    }
    return;
}
*/
import "C"

import (
    "unsafe"
    "reflect"
)

const (
    ColorsetRoot = C.NEWT_COLORSET_ROOT
    COLORSET_ROOT = C.NEWT_COLORSET_ROOT
    ColorsetBorder = C.NEWT_COLORSET_BORDER
    COLORSET_BORDER = C.NEWT_COLORSET_BORDER
    ColorsetWindow = C.NEWT_COLORSET_WINDOW
    COLORSET_WINDOW = C.NEWT_COLORSET_WINDOW
    ColorsetShadow = C.NEWT_COLORSET_SHADOW
    COLORSET_SHADOW = C.NEWT_COLORSET_SHADOW
    ColorsetTitle = C.NEWT_COLORSET_TITLE
    COLORSET_TITLE = C.NEWT_COLORSET_TITLE
    ColorsetButton = C.NEWT_COLORSET_BUTTON
    COLORSET_BUTTON = C.NEWT_COLORSET_BUTTON
    ColorsetActButton = C.NEWT_COLORSET_ACTBUTTON
    COLORSET_ACTBUTTON = C.NEWT_COLORSET_ACTBUTTON
    ColorsetCheckbox = C.NEWT_COLORSET_CHECKBOX
    COLORSET_CHECKBOX = C.NEWT_COLORSET_CHECKBOX
    ColorsetActCheckbox = C.NEWT_COLORSET_ACTCHECKBOX
    COLORSET_ACTCHECKBOX = C.NEWT_COLORSET_ACTCHECKBOX
    ColorsetEntry = C.NEWT_COLORSET_ENTRY
    COLORSET_ENTRY = C.NEWT_COLORSET_ENTRY
    ColorsetLabel = C.NEWT_COLORSET_LABEL
    COLORSET_LABEL = C.NEWT_COLORSET_LABEL
    ColorsetListbox = C.NEWT_COLORSET_LISTBOX
    COLORSET_LISTBOX = C.NEWT_COLORSET_LISTBOX
    ColorsetActListbox = C.NEWT_COLORSET_ACTLISTBOX
    COLORSET_ACTLISTBOX = C.NEWT_COLORSET_ACTLISTBOX
    ColorsetTextbox = C.NEWT_COLORSET_TEXTBOX
    COLORSET_TEXTBOX = C.NEWT_COLORSET_TEXTBOX
    ColorsetActTextbox = C.NEWT_COLORSET_ACTTEXTBOX
    COLORSET_ACTTEXTBOX = C.NEWT_COLORSET_ACTTEXTBOX
    ColorsetHelpLine = C.NEWT_COLORSET_HELPLINE
    COLORSET_HELPLINE = C.NEWT_COLORSET_HELPLINE
    ColorsetRootText = C.NEWT_COLORSET_ROOTTEXT
    COLORSET_ROOTTEXT = C.NEWT_COLORSET_ROOTTEXT
    ColorsetEmptyScale = C.NEWT_COLORSET_EMPTYSCALE
    COLORSET_EMPTYSCALE = C.NEWT_COLORSET_EMPTYSCALE
    ColorsetFullScale = C.NEWT_COLORSET_FULLSCALE
    COLORSET_FULLSCALE = C.NEWT_COLORSET_FULLSCALE
    ColorsetDisEntry = C.NEWT_COLORSET_DISENTRY
    COLORSET_DISENTRY = C.NEWT_COLORSET_DISENTRY
    ColorsetCompactButton = C.NEWT_COLORSET_COMPACTBUTTON
    COLORSET_COMPACTBUTTON = C.NEWT_COLORSET_COMPACTBUTTON
    ColorsetActSelListbox = C.NEWT_COLORSET_ACTSELLISTBOX
    COLORSET_ACTSELLISTBOX = C.NEWT_COLORSET_ACTSELLISTBOX
    ColorsetSelListbox = C.NEWT_COLORSET_SELLISTBOX
    COLORSET_SELLISTBOX = C.NEWT_COLORSET_SELLISTBOX

    ARG_LAST = C.NEWT_ARG_LAST
    ARG_APPEND = C.NEWT_ARG_APPEND

    FLAG_RETURNEXIT = C.NEWT_FLAG_RETURNEXIT
    FlagReturnExit = C.NEWT_FLAG_RETURNEXIT
    FLAG_HIDDEN = C.NEWT_FLAG_HIDDEN
    FlagHidden = C.NEWT_FLAG_HIDDEN
    FLAG_SCROLL = C.NEWT_FLAG_SCROLL
    FlagScroll = C.NEWT_FLAG_SCROLL
    FLAG_DISABLED = C.NEWT_FLAG_DISABLED
    FlagDisabled = C.NEWT_FLAG_DISABLED
    FLAG_BORDER = C.NEWT_FLAG_BORDER
    FlagBorder = C.NEWT_FLAG_BORDER
    FLAG_WRAP = C.NEWT_FLAG_WRAP
    FlagWrap = C.NEWT_FLAG_WRAP
    FLAG_NOF12 = C.NEWT_FLAG_NOF12
    FlagNoF12 = C.NEWT_FLAG_NOF12
    FLAG_MULTIPLE = C.NEWT_FLAG_MULTIPLE
    FlagMultiple = C.NEWT_FLAG_MULTIPLE
    FLAG_SELECTED = C.NEWT_FLAG_SELECTED
    FlagSelected = C.NEWT_FLAG_SELECTED
    FLAG_CHECKBOX = C.NEWT_FLAG_CHECKBOX
    FlagCheckbox = C.NEWT_FLAG_CHECKBOX
    FLAG_PASSWORD = C.NEWT_FLAG_PASSWORD
    FlagPassword = C.NEWT_FLAG_PASSWORD
    FLAG_SHOWCURSOR = C.NEWT_FLAG_SHOWCURSOR
    FlagShowCursor = C.NEWT_FLAG_SHOWCURSOR

    FD_READ = C.NEWT_FD_READ
    FD_WRITE = C.NEWT_FD_WRITE
    FD_EXCEPT = C.NEWT_FD_EXCEPT

    CHECKBOXTREE_UNSELECTABLE = C.NEWT_CHECKBOXTREE_UNSELECTABLE
    CHECKBOXTREE_HIDE_BOX = C.NEWT_CHECKBOXTREE_HIDE_BOX
    CHECKBOXTREE_COLLAPSED = C.NEWT_CHECKBOXTREE_COLLAPSED
    CHECKBOXTREE_EXPANDED = C.NEWT_CHECKBOXTREE_EXPANDED
    CHECKBOXTREE_UNSELECTED = C.NEWT_CHECKBOXTREE_UNSELECTED
    CHECKBOXTREE_SELECTED = C.NEWT_CHECKBOXTREE_SELECTED

    LISTBOX_RETURNEXIT = C.NEWT_LISTBOX_RETURNEXIT
    ENTRY_SCROLL = C.NEWT_ENTRY_SCROLL
    ENTRY_HIDDEN = C.NEWT_ENTRY_HIDDEN
    ENTRY_RETURNEXIT = C.NEWT_ENTRY_RETURNEXIT
    ENTRY_DISABLED = C.NEWT_ENTRY_DISABLED

    TEXTBOX_WRAP = C.NEWT_TEXTBOX_WRAP
    TEXTBOX_SCROLL = C.NEWT_TEXTBOX_SCROLL
    FORM_NOF12 = C.NEWT_FORM_NOF12

    ANCHOR_LEFT = C.NEWT_ANCHOR_LEFT
    ANCHOR_RIGHT = C.NEWT_ANCHOR_RIGHT
    ANCHOR_TOP = C.NEWT_ANCHOR_TOP
    ANCHOR_BOTTOM = C.NEWT_ANCHOR_BOTTOM

    GRID_FLAG_GROWX = C.NEWT_GRID_FLAG_GROWX
    GRID_FLAG_GROWY = C.NEWT_GRID_FLAG_GROWY
)

const (
    KEY_TAB = C.NEWT_KEY_TAB
    KEY_ENTER = C.NEWT_KEY_ENTER
    KEY_SUSPEND = C.NEWT_KEY_SUSPEND
    KEY_ESCAPE = C.NEWT_KEY_ESCAPE
    KEY_RETURN = C.NEWT_KEY_RETURN
    KEY_EXTRA_BASE = C.NEWT_KEY_EXTRA_BASE
    KEY_UP = C.NEWT_KEY_UP
    KEY_DOWN = C.NEWT_KEY_DOWN
    KEY_LEFT = C.NEWT_KEY_LEFT
    KEY_RIGHT = C.NEWT_KEY_RIGHT
    KEY_BKSPC = C.NEWT_KEY_BKSPC
    KEY_DELETE = C.NEWT_KEY_DELETE
    KEY_HOME = C.NEWT_KEY_HOME
    KEY_END = C.NEWT_KEY_END
    KEY_UNTAB = C.NEWT_KEY_UNTAB
    KEY_PGUP = C.NEWT_KEY_PGUP
    KEY_PGDN = C.NEWT_KEY_PGDN
    KEY_INSERT = C.NEWT_KEY_INSERT
    KEY_F1 = C.NEWT_KEY_F1
    KEY_F2 = C.NEWT_KEY_F2
    KEY_F3 = C.NEWT_KEY_F3
    KEY_F4 = C.NEWT_KEY_F4
    KEY_F5 = C.NEWT_KEY_F5
    KEY_F6 = C.NEWT_KEY_F6
    KEY_F7 = C.NEWT_KEY_F7
    KEY_F8 = C.NEWT_KEY_F8
    KEY_F9 = C.NEWT_KEY_F9
    KEY_F10 = C.NEWT_KEY_F10
    KEY_F11 = C.NEWT_KEY_F11
    KEY_F12 = C.NEWT_KEY_F12
    KEY_RESIZE = C.NEWT_KEY_RESIZE
    KEY_ERROR = C.NEWT_KEY_ERROR
)

const (         // newtGridElement
    GRID_EMPTY = iota
    GRID_COMPONENT
    GRID_SUBGRID
)

type Colors struct {
    rootFg, rootBg string
    borderFg, borderBg string
    windowFg, windowBg string
    shadowFg, shadowBg string
    titleFg, titleBg string
    buttonFg, buttonBg string
    actButtonFg, actButtonBg string
    checkboxFg, checkboxBg string
    actCheckboxFg, actCheckboxBg string
    entryFg, entryBg string
    labelFg, labelBg string
    listboxFg, listboxBg string
    actListboxFg, actListboxBg string
    textboxFg, textboxBg string
    helpLineFg, helpLineBg string
    rootTextFg, rootTextBg string
    emptyScale, fullScale string
    disabledEntryFg, disabledEntryBg string
    compactButtonFg, compactButtonBg string
    actSelListboxFg, actSelListboxBg string
    selListboxFg, selListboxBg string
}

type Component struct {
    c C.newtComponent
    g C.newtGrid
}

type WinEntry C.struct_newtWinEntry

type SuspendCallback func([]byte)
type Callback func(*Component, []byte)

func colors2newt(c Colors) C.struct_newtColors {
    var n C.struct_newtColors
    n.rootFg = C.CString(c.rootFg)
    n.rootBg = C.CString(c.rootBg)
    n.borderFg = C.CString(c.borderFg)
    n.borderBg = C.CString(c.borderBg)
    n.windowFg = C.CString(c.windowFg)
    n.windowBg = C.CString(c.windowBg)
    n.shadowFg = C.CString(c.shadowFg)
    n.shadowBg = C.CString(c.shadowBg)
    n.titleFg = C.CString(c.titleFg)
    n.titleBg = C.CString(c.titleBg)
    n.buttonFg = C.CString(c.buttonFg)
    n.buttonBg = C.CString(c.buttonBg)
    n.actButtonFg = C.CString(c.actButtonFg)
    n.actButtonBg = C.CString(c.actButtonBg)
    n.checkboxFg = C.CString(c.checkboxFg)
    n.checkboxBg = C.CString(c.checkboxBg)
    n.actCheckboxFg = C.CString(c.actCheckboxFg)
    n.actCheckboxBg = C.CString(c.actCheckboxBg)
    n.entryFg = C.CString(c.entryFg)
    n.entryBg = C.CString(c.entryBg)
    n.labelFg = C.CString(c.labelFg)
    n.labelBg = C.CString(c.labelBg)
    n.listboxFg = C.CString(c.listboxFg)
    n.listboxBg = C.CString(c.listboxBg)
    n.actListboxFg = C.CString(c.actListboxFg)
    n.actListboxBg = C.CString(c.actListboxBg)
    n.textboxFg = C.CString(c.textboxFg)
    n.textboxBg = C.CString(c.textboxBg)
    n.helpLineFg = C.CString(c.helpLineFg)
    n.helpLineBg = C.CString(c.helpLineBg)
    n.rootTextFg = C.CString(c.rootTextFg)
    n.rootTextBg = C.CString(c.rootTextBg)
    n.emptyScale = C.CString(c.emptyScale)
    n.fullScale = C.CString(c.fullScale)
    n.disabledEntryFg = C.CString(c.disabledEntryFg)
    n.disabledEntryBg = C.CString(c.disabledEntryBg)
    n.compactButtonFg = C.CString(c.compactButtonFg)
    n.compactButtonBg = C.CString(c.compactButtonBg)
    n.actSelListboxFg = C.CString(c.actSelListboxFg)
    n.actSelListboxBg = C.CString(c.actSelListboxBg)
    n.selListboxFg = C.CString(c.selListboxFg)
    n.selListboxBg = C.CString(c.selListboxBg)
    return n
}

func arr2bytes(array uintptr, size int) []byte {
    var bytes = make([]byte, size)

    for i := 0; i < len(bytes); i++ {
        bytes[i] = byte(*(*C.uchar)(unsafe.Pointer(array)))
        array ++
    }

    return bytes
}

func bytes2arr(bytes []byte) uintptr {
    var array = unsafe.Pointer(C.calloc(C.size_t(len(bytes)), 1))
    var arrayptr = uintptr(array)

    for i := 0; i < len(bytes); i++ {
        *(*C.uchar)(unsafe.Pointer(arrayptr)) = C.uchar(bytes[i])
        arrayptr ++
    }

    return uintptr(array)
}

func Init() int {
    return int(C.newtInit())
}

func Finished() int {
    return int(C.newtFinished())
}

func Cls() {
    C.newtCls()
}

func ResizeScreen(redraw bool) {
    if redraw {
        C.newtResizeScreen(1)
    } else {
        C.newtResizeScreen(0)
    }
}

func WaitForKey() {
    C.newtWaitForKey()
}

func ClearKeyBuffer() {
    C.newtClearKeyBuffer()
}

func Delay(usecs uint) {
    C.newtDelay(C.uint(usecs))
}

func OpenWindow(left, top int, width, height uint, title string) int {
    t := C.CString(title)
    defer C.free(unsafe.Pointer(t))
    return int(C.newtOpenWindow(C.int(left), C.int(top), C.uint(width), C.uint(height), t))
}

func CenteredWindow(width, height uint, title string) int {
    t := C.CString(title)
    defer C.free(unsafe.Pointer(t))
    return int(C.newtCenteredWindow(C.uint(width), C.uint(height), t))
}

func PopWindow() {
    C.newtPopWindow()
}

func PopWindowNoRefresh() {
    C.newtPopWindowNoRefresh()
}

func SetColors(colors Colors) {
    n := colors2newt(colors)
    C.newtSetColors(n)
}

func SetColor(colorSet int, fg, bg string) {
    f := C.CString(fg)
    b := C.CString(bg)
    defer C.free(unsafe.Pointer(f))
    defer C.free(unsafe.Pointer(b))
    C.newtSetColor(C.int(colorSet), f, b)
}

func Refresh() {
    C.newtRefresh()
}

func Suspend() {
    C.newtSuspend()
}

func SetSuspendCallback(cb *SuspendCallback, data *[]byte) {
    C.n_setSuspendCallback(unsafe.Pointer(cb), unsafe.Pointer(data))
}

func SetHelpCallback(cb *Callback) {
    C.newtSetHelpCallback(C.newtCallback(unsafe.Pointer(C.n_HelpCallback)))
    C.helpCallback = *(*C.newtCallback)(unsafe.Pointer(cb))
}

func Resume() int {
    return int(C.newtResume())
}

func PushHelpLine(help string) {
    h := C.CString(help)
    defer C.free(unsafe.Pointer(h))
    C.newtPushHelpLine(h)
}

func RedrawHelpLine() {
    C.newtRedrawHelpLine()
}

func PopHelpLine() {
    C.newtPopHelpLine()
}

func DrawRootText(col, row int, text string) {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    C.newtDrawRootText(C.int(col), C.int(row), t)
}

func Bell() {
    C.newtBell()
}

func CursorOff() {
    C.newtCursorOff()
}

func CursorOn() {
    C.newtCursorOn()
}

/* Components */
func GetScreenSize() (int, int) {
    var cols, rows C.int
    C.newtGetScreenSize(&cols, &rows)
    return int(cols), int(rows)
}

func ComponentAddCallback(c Component, f Callback, data []byte) {
    panic("not implemented")
}

func ComponentTakesFocus(c Component, val int) {
    C.newtComponentTakesFocus(c.c, C.int(val))
}

func ComponentGetPosition(c Component) (int, int) {
    var left, top C.int
    C.newtComponentGetPosition(c.c, &left, &top)
    return int(left), int(top)
}

func ComponentGetSize(c Component) (int, int) {
    var width, height C.int
    C.newtComponentGetSize(c.c, &width, &height)
    return int(width), int(height)
}

func ComponentAddDestroyCallback(c Component, f Callback, data []byte) {
    panic("not implemented")
}

func ComponentDestroy(c Component) {
    C.newtComponentDestroy(c.c)
}

