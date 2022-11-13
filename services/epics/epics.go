package epics

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L../../build/local/lib -L../../build/local/lib/linux-x86_64 -lstdc++ -lsciarc -lpvaClient -lpvAccess -lpthread -lboost_system
#include <../../build/local/include/sciarc/sciarc.h>
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export goCallbackHandler
func goCallbackHandler(channel_name *C.char, buff unsafe.Pointer, len C.int) {
	fmt.Print("goCallbackHandler for channel" + C.GoString(channel_name))
}

// Start monitoring data for a specified channel
func MonitorChannel(channle string) error {
	C.ACFunction()
	return nil
}
