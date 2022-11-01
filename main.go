package main

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -Lbuild/local/lib -lsciarc
#include "build/local/include/sciarc/sciarc.h"
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
