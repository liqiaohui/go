package cgoexample

/*
#include <stdio.h>
#include <stdlib.h>

void myprintf(char* s) {
	printf("%s\n", s)
}
*/
import "C"

import "unsafe"

func GoPrintf(s string) {
	cs := C.CString(s)
	C.myprintf(cs)
	C.free(unsafe.Pointer(cs))
}
