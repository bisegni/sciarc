package services

import (
	"fmt"
)

// GetApi
type StopMonitorApiFactory struct {
}

func (smf StopMonitorApiFactory) Name() string {
	return "nmon"
}

func (smf StopMonitorApiFactory) BuildApiByString(result_channel chan<- *ApiResult, conn_id int, conf *string) interface{} {
	return &StopMonitorApi{
		conn_id:      conn_id,
		channel_name: conf,
		res_chan:     result_channel,
	}
}
func (smf StopMonitorApiFactory) BuildApiByBinary(result_channel chan<- *ApiResult, conn_id int, conf *[]byte) interface{} {
	return nil
}

// GetApi
type StopMonitorApi struct {
	conn_id      int
	channel_name *string
	res_chan     chan<- *ApiResult
}

func (sm *StopMonitorApi) Run() {
	// Do your task here

	StopMonitorChannel(*sm.channel_name, sm.conn_id)
	message := fmt.Sprintf("Monitor deactivated on %s", *sm.channel_name)
	sm.res_chan <- &ApiResult{sm.conn_id, String, &message, nil}
}
