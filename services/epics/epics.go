package epics

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L../../build/local/lib -L../../build/local/lib/linux-x86_64 -lstdc++ -lsciarc -lpvaClient -lpvAccess -lCom -lpvData -lpvDatabase -lpvAccessCA -lpthread -lboost_system
#include <../../build/local/include/sciarc/sciarc.h>
#include <stdio.h>
#include <stdlib.h>
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

// StartChannelMonitor monitoring data for a specified channel
func StartChannelMonitor(channel string) error {
	return nil
}

// StopChannelMonitor
func StopChannelMonitor(channel string) error {
	return nil
}

// StopChannelMonitor
func GetChannelvalue(channel string) (string, error) {
	cstr := C.CString(channel)
	defer C.free(unsafe.Pointer(cstr))
	var out *C.char = C.getData(cstr)
	defer C.free(unsafe.Pointer(out))
	return C.GoString(out), nil
}
