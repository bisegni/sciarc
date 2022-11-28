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

// EventData the data emitted by event monitor
type EventData struct {
	channel string
	data    []byte
	len     int32
}

//export goCallbackHandler
func goCallbackHandler(channel_name *C.char, buf unsafe.Pointer, len C.int) {
	defer C.free(unsafe.Pointer(channel_name))
	var event = &EventData{
		channel: C.GoString(channel_name),
		data:    unsafe.Slice((*byte)(buf), int32(len)),
		len:     int32(len),
	}

	fmt.Printf("goCallbackHandler for channel %s\n", event.channel)
}

// StartChannelMonitor monitoring data for a specified channel
func StartChannelMonitor(channel string, data_channel chan<- EventData) error {
	cstr := C.CString(channel)
	defer C.free(unsafe.Pointer(cstr))
	C.startMonitor(cstr)
	return nil
}

// StopChannelMonitor
func StopChannelMonitor(channel string) error {
	cstr := C.CString(channel)
	defer C.free(unsafe.Pointer(cstr))
	C.stopMonitor(cstr)
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
