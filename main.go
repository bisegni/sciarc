package main

/*
#cgo CFLAGS: -I . -stdlib=libc++
#cgo LDFLAGS: -L${SRCDIR}/build/depinstall/lib -ldcomp -lstdc++
#include "engine/src/dcomp.h"
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
*/
import "C"
import "fmt"

//export goCallbackHandler
func goCallbackHandler() {
	fmt.Print("goCallbackHandler")
}

func main() {
	C.ACFunction()
}
