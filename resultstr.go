package newt

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"

type ResultStr struct {
    value *C.char
}

func NewResultStr() ResultStr {
    var r ResultStr
    r.value = (*C.char)(C.malloc(1024))
    return r
}

func (r ResultStr) String() string {
    return C.GoString(r.value)
}

func (r ResultStr) Destroy() {
    C.free(unsafe.Pointer(r.value))
}
