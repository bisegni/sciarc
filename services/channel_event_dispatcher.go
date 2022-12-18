package services

import (
	"sync"

	"github.com/bisegni/sciarc/services/epics"
	"golang.org/x/exp/slices"
)

var stringEventCallback func(id int, resutl *string)
var dispacthChannel = make(chan *epics.EpicsEventData)
var mtxChnlConnection sync.RWMutex
var channelConnListener = make(map[string][]int)

func dispatch() {
	if stringEventCallback != nil {
		return
	}
	for {
		select {
		case eventData := <-dispacthChannel:
			{
				for _, conn := range channelConnListener[eventData.Channel] {
					str := string(eventData.Data)
					stringEventCallback(conn, &str)
				}
				break
			}
		}
	}
}

func init() {
	go dispatch()
}

func StartMonitorChannel(channel string, conn_id int) {
	mtxChnlConnection.Lock()
	defer mtxChnlConnection.Unlock()
	idx := slices.IndexFunc(channelConnListener[channel], func(conn int) bool { return conn == conn_id })
	if idx == -1 {
		channelConnListener[channel] = append(channelConnListener[channel], conn_id)
	}
	if len(channelConnListener[channel]) == 1 {
		epics.StartChannelMonitor(channel, dispacthChannel)
	}
}

func StopMonitorChannel(channel string, conn_id int) {
	mtxChnlConnection.Lock()
	defer mtxChnlConnection.Unlock()

	idx := slices.IndexFunc(channelConnListener[channel], func(conn int) bool { return conn == conn_id })
	if idx != -1 {
		channelConnListener[channel] = append(channelConnListener[channel][:idx], channelConnListener[channel][idx+1:]...)
	} else {
		return
	}
	if len(channelConnListener[channel]) == 0 {
		epics.StopChannelMonitor(channel)
	}

}
