package newt

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

type ResultStr struct {
    value *C.char
}

func NewResultStr() ResultStr {
    var r ResultStr
    r.value = (*C.char)(C.malloc(1024))
    C.memset(unsafe.Pointer(r.value), 0, 1024)
    return r
}

func (r ResultStr) Set(s string) {
    if len(s) <= 1024 {
        s1 := C.CString(s)
        defer C.free(unsafe.Pointer(s1))
        C.strncpy(r.value, s1, C.size_t(len(s)))
    }
}

func (r ResultStr) String() string {
    return C.GoString(r.value)
}

func (r ResultStr) Destroy() {
    C.free(unsafe.Pointer(r.value))
}
