package services

import (
	"github.com/bisegni/sciarc/services/epics"
)

// GetApi
type GetApiFactory struct {
}

func (gaf GetApiFactory) Name() string {
	return "get"
}

func (gaf GetApiFactory) BuildApiByString(result_channel chan<- *ApiResult, conn_id int, conf *string) interface{} {
	return &GetApi{
		conn_id:      conn_id,
		channel_name: conf,
		res_chan:     result_channel,
	}
}
func (gaf GetApiFactory) BuildApiByBinary(result_channel chan<- *ApiResult, conn_id int, conf *[]byte) interface{} {
	return nil
}

// GetApi
type GetApi struct {
	conn_id      int
	channel_name *string
	res_chan     chan<- *ApiResult
}

func (ga *GetApi) Run() {
	var result *ApiResult
	// Do your task here
	value, err := epics.GetChannelvalue(*ga.channel_name)
	if err != nil {
		result = &ApiResult{ga.conn_id, String, nil, err}
	} else {
		result = &ApiResult{ga.conn_id, String, &value, nil}

	}
	ga.res_chan <- result
}
