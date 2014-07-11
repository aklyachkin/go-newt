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
func CompactButton(left, top int, text string) Component {
    var c Component
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    a := C.newtCompactButton(C.int(left), C.int(top), t)
    c.c = a
    return c
}

func Button(left, top int, text string) Component {
    var c Component
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    a := C.newtButton(C.int(left), C.int(top), t)
    c.c = a
    return c
}

func Checkbox(left, top int, text string, defValue, seq string, result ResultStr) Component {
    var c Component
    t := C.CString(text)
    var s *C.char
    if seq == "" {
        s = nil
    } else {
        s = C.CString(seq)
    }
    defer C.free(unsafe.Pointer(t))
    defer C.free(unsafe.Pointer(s))
    c.c = C.newtCheckbox(C.int(left), C.int(top), t, C.char(defValue[0]), s, result.value)
    return c
}

func CheckboxGetValue(c Component) int {
    a := C.newtCheckboxGetValue(c.c)
    return int(a)
}

func CheckboxSetValue(c Component, value int) {
    C.newtCheckboxSetValue(c.c, C.char(value))
}

func CheckboxSetFlags(c Component, flags int, sense uint32) {
    C.newtCheckboxSetFlags(c.c, C.int(flags), sense)
}

func Radiobutton(left, top int, text string, isDefault int, prevButton *Component) Component {
    var c Component
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    if prevButton == nil {
        c.c = C.newtRadiobutton(C.int(left), C.int(top), t, C.int(isDefault), nil)
    } else {
        c.c = C.newtRadiobutton(C.int(left), C.int(top), t, C.int(isDefault), prevButton.c)
    }
    return c
}

func RadioGetCurrent(setMember Component) Component {
    var c Component
    c.c = C.newtRadioGetCurrent(setMember.c)
    return c
}

func RadioSetCurrent(setMember Component) {
    C.newtRadioSetCurrent(setMember.c)
}

func GetScreenSize() (int, int) {
    var cols, rows C.int
    C.newtGetScreenSize(&cols, &rows)
    return int(cols), int(rows)
}

func Label(left, top int, text string) Component {
    var c Component
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    c.c = C.newtLabel(C.int(left), C.int(top), t)
    return c
}

func LabelSetText(c Component, text string) {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    C.newtLabelSetText(c.c, t)
}

func LabelSetColors(c Component, colorset int) {
    C.newtLabelSetColors(c.c, C.int(colorset))
}

func VerticalScrollbar(left, top, height, normalColorset, thumbColorset int) Component {
    var c Component
    c.c = C.newtVerticalScrollbar(C.int(left), C.int(top), C.int(height), C.int(normalColorset), C.int(thumbColorset))
    return c
}

func ScrollbarSet(c Component, where, total int) {
    C.newtScrollbarSet(c.c, C.int(where), C.int(total))
}

func ScrollbarSetColors(c Component, normal, thumb int) {
    C.newtScrollbarSetColors(c.c, C.int(normal), C.int(thumb))
}

func Listbox(left, top, height, flags int) Component {
    var c Component
    c.c = C.newtListbox(C.int(left), C.int(top), C.int(height), C.int(flags))
    return c
}

func ListboxGetCurrent(c Component) uintptr {
    bptr := C.newtListboxGetCurrent(c.c)
    return uintptr(bptr)
}

func ListboxSetCurrent(c Component, num int) {
    C.newtListboxSetCurrent(c.c, C.int(num))
}

func ListboxSetCurrentByKey(c Component, key uintptr) {
    C.newtListboxSetCurrentByKey(c.c, unsafe.Pointer(key))
}

func ListboxSetEntry(c Component, num int, text string) {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    C.newtListboxSetEntry(c.c, C.int(num), t)
}

func ListboxSetWidth(c Component, width int) {
    C.newtListboxSetWidth(c.c, C.int(width))
}

func ListboxSetData(c Component, num int, data uintptr) {
    C.newtListboxSetData(c.c, C.int(num), unsafe.Pointer(data))
}

func ListboxAppendEntry(c Component, text string, data uintptr) int {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    return int(C.newtListboxAppendEntry(c.c, t, unsafe.Pointer(data)))
}

func ListboxAddEntry(c Component, text string, data uintptr) int {
    return ListboxAppendEntry(c, text, data)
}

func ListboxInsertEntry(c Component, text string, data, key uintptr) int {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    return int(C.newtListboxInsertEntry(c.c, t, unsafe.Pointer(data), unsafe.Pointer(key)))
}

func ListboxDeleteEntry(c Component, data uintptr) int {
    return int(C.newtListboxDeleteEntry(c.c, unsafe.Pointer(data)))
}

func ListboxClear(c Component) {
    C.newtListboxClear(c.c)
}

func ListboxGetEntry(c Component, num int) (string, uintptr) {
    var text *C.char
    var data unsafe.Pointer
    C.newtListboxGetEntry(c.c, C.int(num), &text, &data)
    return C.GoString(text), uintptr(data)
}

func ListboxGetSelection(c Component) (int, []uintptr) {
    var numitems C.int
    var i []uintptr
    items := C.newtListboxGetSelection(c.c, &numitems)
    if items != nil {
        hdr := reflect.SliceHeader{
            Data: uintptr(unsafe.Pointer(items)),
            Len: int(numitems),
            Cap: int(numitems),
        }
        i = *(*[]uintptr)(unsafe.Pointer(&hdr))
    }
    return int(numitems), i
}

func ListboxClearSelection(c Component) {
    C.newtListboxClearSelection(c.c)
}

func ListboxSelectItem(c Component, key uintptr, sense uint32) {
    C.newtListboxSelectItem(c.c, unsafe.Pointer(key), sense)
}

func ListboxItemCount(c Component) int {
    n := C.newtListboxItemCount(c.c)
    return int(n)
}

func CheckboxTree(left, top, height, flags int) Component {
    var c Component
    c.c = C.newtCheckboxTree(C.int(left), C.int(top), C.int(height), C.int(flags))
    return c
}

func CheckboxTreeMulti(left, top, height int, seq string, flags int) Component {
    var c Component
    s := C.CString(seq)
    defer C.free(unsafe.Pointer(s))
    c.c = C.newtCheckboxTreeMulti(C.int(left), C.int(top), C.int(height), s, C.int(flags))
    return c
}

func CheckboxTreeGetSelection(c Component) (int, *uintptr) {
    var numitems C.int
    items := C.newtCheckboxTreeGetSelection(c.c, &numitems)
    i := uintptr(*items)
    return int(numitems), &i
}

func CheckboxTreeGetCurrent(c Component) uintptr {
    return uintptr(C.newtCheckboxTreeGetCurrent(c.c))
}

func CheckboxTreeSetCurrent(c Component, item uintptr) {
    C.newtCheckboxTreeSetCurrent(c.c, unsafe.Pointer(item))
}

func CheckboxTreeGetMultiSelection(c Component, seqnum int) (int, *uintptr) {
    var numitems C.int
    items := C.newtCheckboxTreeGetMultiSelection(c.c, &numitems, C.char(seqnum))
    i := uintptr(*items)
    return int(numitems), &i
}

func CheckboxTreeAddItem(c Component, text string, data uintptr, flags int, index int) int {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    return int(C.newtCheckboxTreeAddItem(c.c, t, unsafe.Pointer(data), C.int(flags), C.int(index)))
}

func CheckboxTreeAddArray(c Component, text string, data uintptr, flags int, indexes []int) int {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    var i []C.int
    for a := range indexes {
        i = append(i, C.int(a))
    }
    return int(C.newtCheckboxTreeAddArray(c.c, t, unsafe.Pointer(data), C.int(flags), &i[0]))
}

func CheckboxTreeFindItem(c Component, data uintptr) int {
    r := C.newtCheckboxTreeFindItem(c.c, unsafe.Pointer(data))
    return int(*r)
}

func CheckboxTreeSetEntry(c Component, data uintptr, text string) {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    C.newtCheckboxTreeSetEntry(c.c, unsafe.Pointer(data), t)
}

func CheckboxTreeSetWidth(c Component, width int) {
    C.newtCheckboxTreeSetWidth(c.c, C.int(width))
}

func CheckboxTreeGetEntryValue(c Component) (int, uintptr) {
    var p unsafe.Pointer
    r := C.newtCheckboxTreeGetEntryValue(c.c, p)
    return int(r), uintptr(p)
}

func CheckboxTreeSetEntryValue(c Component, data uintptr, value int) {
    C.newtCheckboxTreeSetEntryValue(c.c, unsafe.Pointer(data), C.char(value))
}

func TextboxReflowed(left, top int, text string, width, flexDown, flexUp, flags int) Component {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    var c Component
    c.c = C.newtTextboxReflowed(C.int(left), C.int(top), t, C.int(width), C.int(flexDown), C.int(flexUp), C.int(flags))
    return c
}

func Textbox(left, top, width, height, flags int) Component {
    var c Component
    c.c = C.newtTextbox(C.int(left), C.int(top), C.int(width), C.int(height), C.int(flags))
    return c
}

func TextboxSetText(c Component, text string) {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    C.newtTextboxSetText(c.c, t)
}

func TextboxSetHeight(c Component, height int) {
    C.newtTextboxSetHeight(c.c, C.int(height))
}

func TextboxGetNumLines(c Component) int {
    return int(C.newtTextboxGetNumLines(c.c))
}

func TextboxSetColors(c Component, normal, active int) {
    C.newtTextboxSetColors(c.c, C.int(normal), C.int(active))
}

func ReflowText(text string, width, flexDown, flexUp int) (string, int, int) {
    t := C.CString(text)
    defer C.free(unsafe.Pointer(t))
    var actualWidth, actualHeight *C.int
    r := C.newtReflowText(t, C.int(width), C.int(flexDown), C.int(flexUp), actualWidth, actualHeight)
    defer C.free(unsafe.Pointer(r))
    return C.GoString(r), int(*actualWidth), int(*actualHeight)
}

func Form(vertBar *Component, helpTag string, flags int) Component {
    var c Component
    t := C.CString(helpTag)
    defer C.free(unsafe.Pointer(t))
    if vertBar != nil {
        c.c = C.newtForm(vertBar.c, unsafe.Pointer(t), C.int(flags))
    } else {
        c.c = C.newtForm(nil, unsafe.Pointer(t), C.int(flags))
    }
    return c
}

func FormSetTimer(form Component, millisecs int) {
    C.newtFormSetTimer(form.c, C.int(millisecs))
}

func FormWatchFd(form Component, fd, fdFlags int) {
    C.newtFormWatchFd(form.c, C.int(fd), C.int(fdFlags))
}

func FormSetSize(c Component) {
    C.newtFormSetSize(c.c)
}

func FormGetCurrent(form Component) Component {
    var c Component
    c.c = C.newtFormGetCurrent(form.c)
    return c
}

func FormSetBackground(c Component, color int) {
    C.newtFormSetBackground(c.c, C.int(color))
}

func FormSetCurrent(form, c Component) {
    C.newtFormSetCurrent(form.c, c.c)
}

func FormAddComponent(form, c Component) {
    C.newtFormAddComponent(form.c, c.c)
}

func FormAddComponents(form Component, args ...Component) {
    for _, v := range args {
        FormAddComponent(form, v)
    }
    // C.newtFormAddComponents(form.c, args)
}

func FormSetHeight(c Component, height int) {
    C.newtFormSetHeight(c.c, C.int(height))
}

func FormSetWidth(c Component, width int) {
    C.newtFormSetWidth(c.c, C.int(width))
}

func RunForm(form Component) Component {
    var c Component
    c.c = C.newtRunForm(form.c)
    return c
}

func FormRun(form Component) ExitStruct {
    // panic
    var es ExitStruct
    C.newtFormRun(form.c, &es.es)
    return es
}

func DrawForm(form Component) {
    C.newtDrawForm(form.c)
}

func FormAddHotKey(c Component, key int) {
    C.newtFormAddHotKey(c.c, C.int(key))
}

func FormGetScrollPosition(c Component) int {
    return int(C.newtFormGetScrollPosition(c.c))
}

func FormSetScrollPosition(c Component, position int) {
    C.newtFormSetScrollPosition(c.c, C.int(position))
}

func Entry(left, top int, initialValue string, width int, result *ResultStr, flags int) Component {
    var c Component
    iv := C.CString(initialValue)
    defer C.free(unsafe.Pointer(iv))
    c.c = C.newtEntry(C.int(left), C.int(top), iv, C.int(width), (**C.char)(&result.value), C.int(flags))
    return c
}

func EntrySet(c Component, value string, cursorAtEnd int) {
    t := C.CString(value)
    defer C.free(unsafe.Pointer(t))
    C.newtEntrySet(c.c, t, C.int(cursorAtEnd))
}

func EntrySetFilter(c Component, filter uintptr, data []byte) {
    panic("not implemented")
}

func EntryGetValue(c Component) string {
    r := C.newtEntryGetValue(c.c)
    defer C.free(unsafe.Pointer(r))
    return C.GoString(r)
}

func EntrySetFlags(c Component, flags int, sense uint32) {
    C.newtEntrySetFlags(c.c, C.int(flags), sense)
}

func EntrySetColors(c Component, normal, disabled int) {
    C.newtEntrySetColors(c.c, C.int(normal), C.int(disabled))
}

func EntryGetCursorPosition(c Component) int {
    return int(C.newtEntryGetCursorPosition(c.c))
}

func EntrySetCursorPosition(c Component, position int) {
    C.newtEntrySetCursorPosition(c.c, C.int(position))
}

func Scale(left, top, width int, fullValue int64) Component {
    var c Component
    c.c = C.newtScale(C.int(left), C.int(top), C.int(width), C.longlong(fullValue))
    return c
}

func ScaleSet(c Component, amount uint64) {
    C.newtScaleSet(c.c, C.ulonglong(amount))
}

func ScaleSetColors(c Component, empty, full int) {
    C.newtScaleSetColors(c.c, C.int(empty), C.int(full))
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

func FormDestroy(c Component) {
    C.newtFormDestroy(c.c)
}

func ComponentDestroy(c Component) {
    C.newtComponentDestroy(c.c)
}

func CreateGrid(cols, rows int) Component {
    var c Component
    c.g = C.newtCreateGrid(C.int(cols), C.int(rows))
    return c
}

func GridVStacked(typ uint32, what uintptr) Component {
    panic("not implemented")
}

func GridVCloseStacked(typ uint32, what uintptr) Component {
    panic("not implemented")
}

func GridHStacked(typ uint32, what uintptr) Component {
    panic("not implemented")
}

func GridHCloseStacked(typ uint32, what uintptr) Component {
    panic("not implemented")
}

func GridBasicWindow(text, middle, buttons Component) Component {
    var c Component
    c.g = C.newtGridBasicWindow(text.c, middle.g, buttons.g)
    return c
}

func GridSimpleWindow(text, middle, buttons Component) Component {
    var c Component
    c.g = C.newtGridSimpleWindow(text.c, middle.c, buttons.g)
    return c
}

func GridSetField(c Component, col, row int, typ uint32, val Component, padLeft, padTop, padRight, padBottom, anchor, flags int) {
    switch typ {
    case GRID_COMPONENT:
        C.newtGridSetField(c.g, C.int(col), C.int(row), typ, unsafe.Pointer(val.c), C.int(padLeft), C.int(padTop), C.int(padRight), C.int(padBottom), C.int(anchor), C.int(flags))
    case GRID_SUBGRID:
        C.newtGridSetField(c.g, C.int(col), C.int(row), typ, unsafe.Pointer(val.g), C.int(padLeft), C.int(padTop), C.int(padRight), C.int(padBottom), C.int(anchor), C.int(flags))
    case GRID_EMPTY:
        C.newtGridSetField(c.g, C.int(col), C.int(row), typ, nil, C.int(padLeft), C.int(padTop), C.int(padRight), C.int(padBottom), C.int(anchor), C.int(flags))
    }
}

func GridPlace(c Component, left, top int) {
    C.newtGridPlace(c.g, C.int(left), C.int(top))
}

func GridDestroy(c Component, recurse int) {
    C.newtGridFree(c.g, C.int(recurse))
}

func GridFree(c Component, recurse int) {
    C.newtGridFree(c.g, C.int(recurse))
}

func GridGetSize(c Component) (int, int) {
    var width, height C.int
    C.newtGridGetSize(c.g, &width, &height)
    return int(width), int(height)
}

func GridWrappedWindow(c Component, title string) {
    t := C.CString(title)
    defer C.free(unsafe.Pointer(t))
    C.newtGridWrappedWindow(c.g, t)
}

func GridWrappedWindowAt(c Component, title string, left, top int) {
    t := C.CString(title)
    defer C.free(unsafe.Pointer(t))
    C.newtGridWrappedWindowAt(c.g, t, C.int(left), C.int(top))
}

func GridAddComponentsToForm(grid, form Component, recurse int) {
    C.newtGridAddComponentsToForm(grid.g, form.c, C.int(recurse))
}

func ButtonBarV(button1 string, b1 Component) {
    panic("not implemented")
}

func ButtonBar(button1 string, b1 Component) {
    panic("not implemented")
}

func WinMessage(title, buttonText, text string) {
    t1 := C.CString(title)
    b1 := C.CString(buttonText)
    t2 := C.CString(text)
    defer func() {
        C.free(unsafe.Pointer(t1))
        C.free(unsafe.Pointer(b1))
        C.free(unsafe.Pointer(t2))
    }()
    C.newtWinMessage(t1, b1, t2)
}

func WinMessageV(title, buttonText, text string) {
    panic("not implemented")
}

func WinChoice(title, button1, button2, text string) int {
    t1 := C.CString(title)
    b1 := C.CString(button1)
    b2 := C.CString(button2)
    t2 := C.CString(text)
    defer func() {
        C.free(unsafe.Pointer(t1))
        C.free(unsafe.Pointer(b1))
        C.free(unsafe.Pointer(b2))
        C.free(unsafe.Pointer(t2))
    }()
    return int(C.newtWinChoice(t1, b1, b2, t2))
}

func WinTernary(title, button1, button2, button3, message string) int {
    t1 := C.CString(title)
    b1 := C.CString(button1)
    b2 := C.CString(button2)
    b3 := C.CString(button3)
    m1 := C.CString(message)
    defer func() {
        C.free(unsafe.Pointer(t1))
        C.free(unsafe.Pointer(b1))
        C.free(unsafe.Pointer(b2))
        C.free(unsafe.Pointer(b3))
        C.free(unsafe.Pointer(m1))
    }()
    return int(C.newtWinTernary(t1, b1, b2, b3, m1))
}

func WinMenu(title, text string, suggestedWidth, flexDown, flexUp, maxListHeight int, items []string, button1 string) (int, int) {
    t1 := C.CString(title)
    t2 := C.CString(text)
    b1 := C.CString(button1)
    defer func() {
        C.free(unsafe.Pointer(t1))
        C.free(unsafe.Pointer(t2))
        C.free(unsafe.Pointer(b1))
    }()
    panic("not implemented")
    // var listItem C.int
    // r := C.newtWinMenu(t1, t2, C.int(suggestedWidth), C.int(flexDown), C.int(flexUp), C.int(maxListHeight), items, &listItem, b1)
    // return int(r), int(listItem)
}
