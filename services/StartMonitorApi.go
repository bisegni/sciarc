package services

import (
	"fmt"
)

// GetApi
type StartMonitorApiFactory struct {
}

func (smf StartMonitorApiFactory) Name() string {
	return "mon"
}

func (smf StartMonitorApiFactory) BuildApiByString(result_channel chan<- *ApiResult, conn_id int, conf *string) interface{} {
	return &StartMonitorApi{
		conn_id:      conn_id,
		channel_name: conf,
		res_chan:     result_channel,
	}
}
func (smf StartMonitorApiFactory) BuildApiByBinary(result_channel chan<- *ApiResult, conn_id int, conf *[]byte) interface{} {
	return nil
}

// GetApi
type StartMonitorApi struct {
	conn_id      int
	channel_name *string
	res_chan     chan<- *ApiResult
}

func (sm *StartMonitorApi) Run() {
	// Do your task here

	StartMonitorChannel(*sm.channel_name, sm.conn_id)
	message := fmt.Sprintf("Monitor activated on %s", *sm.channel_name)
	sm.res_chan <- &ApiResult{sm.conn_id, String, &message, nil}
}
