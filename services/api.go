package services

import (
	"github.com/bisegni/sciarc/services/epics"
	"github.com/shettyh/threadpool"
)

var pool *threadpool.ThreadPool

func init() {
	pool = threadpool.NewThreadPool(200, 1000000)
}

type ResultType int32

const (
	Binary ResultType = iota
	String
)

// Api result
type ApiResult struct {
	conn_id     int32
	result_type ResultType
	result      interface{}
	err         error
}

// GetApi
type GetApi struct {
	conn_id      int
	channel_name string
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

func CallGet(conn_id int, channel_name string, res_chan chan *ApiResult) error {
	return pool.Execute(&GetApi{
		conn_id:      conn_id,
		channel_name: channel_name,
		res_chan:     res_chan,
	})
}
