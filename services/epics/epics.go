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
	"sync"
	"unsafe"
)

var mtx sync.RWMutex
var epicsMonitorChannel = make(map[string]chan<- *EpicsEventData)
var ErrorChannelAlreadyRegistered = fmt.Errorf("channel alredy registered")

// EventData the data emitted by event monitor
type EpicsEventData struct {
	Channel string
	Data    []byte
	Len     int32
}

func init() {
	C.init()
}

//export goCallbackHandler
func goCallbackHandler(channel_name *C.char, buf unsafe.Pointer, len C.int) {
	//defer C.free(unsafe.Pointer(channel_name))
	mtx.Lock()
	defer mtx.Unlock()
	//defer C.free(buf)
	var event = &EpicsEventData{
		Channel: C.GoString(channel_name),
		Data:    unsafe.Slice((*byte)(buf), int32(len)),
		Len:     int32(len),
	}
	if _, ok := epicsMonitorChannel[event.Channel]; !ok {
		//no channel present
		return
	}
	// push event into channel
	epicsMonitorChannel[event.Channel] <- event
	fmt.Printf("goCallbackHandler for channel %s\n", event.Channel)
}

// StartChannelMonitor monitoring data for a specified channel
func StartChannelMonitor(channel string, data_channel chan<- *EpicsEventData) error {
	mtx.Lock()
	defer mtx.Unlock()

	if _, ok := epicsMonitorChannel[channel]; ok {
		//already present
		return ErrorChannelAlreadyRegistered
	}

	// add new channel
	epicsMonitorChannel[channel] = data_channel

	cstr := C.CString(channel)
	//defer C.free(unsafe.Pointer(cstr))
	C.startMonitor(cstr)

	return nil
}

// StopChannelMonitor
func StopChannelMonitor(channel string) error {
	cstr := C.CString(channel)
	//defer C.free(unsafe.Pointer(cstr))
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
