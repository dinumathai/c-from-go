package main

/*
#cgo LDFLAGS: -L. -lmylib
#include "c-lib/mylib.h"
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {

	fmt.Println("--------------START-----------------")

	// C Library
	name := C.CString("User")
	cMsg := C.getGreetingMsg(name)
	C.free(unsafe.Pointer(name))

	fmt.Println(C.GoString(cMsg))
	C.free(unsafe.Pointer(cMsg))

	fmt.Println("--------------END-----------------")
}
