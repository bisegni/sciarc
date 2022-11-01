package main

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L${SRCDIR}/build/local/lib -ldcomp
#include "${SRCDIR}/build/local/sciarc/sciarc.h"
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
*/
import "C"
import (
	"fmt"

	"github.com/bisegni/sciarc/cmd"
)

//export goCallbackHandler
func goCallbackHandler() {
	fmt.Print("goCallbackHandler")
}

func main() {
	cmd.Execute()
}
