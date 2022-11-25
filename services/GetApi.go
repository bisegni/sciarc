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

func (gaf GetApiFactory) BuildApiByString(conn_id int, conf *string) interface{} {
	return &GetApi{
		conn_id:      conn_id,
		channel_name: conf,
	}
}
func (gaf GetApiFactory) BuildApiByBinary(conn_id int, conf *[]byte) interface{} {
	return nil
}

// GetApi
type GetApi struct {
	conn_id      int
	channel_name *string
	res_chan     chan *ApiResult
}

func (ga *GetApi) Run() {
	var result ApiResult
	// Do your task here
	value, err := epics.GetChannelvalue(ga.channel_name)
	if err != nil {
		result.err = err
	} else {
		result.result_type = String
		result.result = value
	}
	ga.res_chan <- &result
}
