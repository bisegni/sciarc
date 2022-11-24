package services

import (
	"log"

	"github.com/bisegni/sciarc/services/epics"
	"github.com/shuLhan/share/lib/debug"
	"github.com/shuLhan/share/lib/websocket"
)

// Node...
type Node struct {
	name string
	age  int
}

// RequestSubmission ...
type RequestSubmission struct {
	conn_id     int32
	result_type ResultType
	result      interface{}
	err         error
}

//Chasnnel for the request submission
var submission_channel chan *RequestSubmission = make(chan *RequestSubmission)
var result_channel chan *ApiResult = make(chan *ApiResult)

// handleBin from websocket by echo-ing back the payload.
func handleBin(conn int, payload []byte) {
	var (
		packet []byte = websocket.NewFrameBin(false, payload)
		err    error
	)

	err = websocket.Send(conn, packet)
	if err != nil {
		log.Println("handleBin: " + err.Error())
	}
}

// handleText from websocket by echo-ing back the payload.
func handleText(conn int, payload []byte) {
	var (
		err error
	)

	if debug.Value >= 3 {
		log.Printf("testdata/server: handleText: {payload.len:%d}\n", len(payload))
	}

	var value, _ = epics.GetChannelvalue(string(payload))
	var packet []byte = websocket.NewFrameText(false, []byte(value))
	err = websocket.Send(conn, packet)
	if err != nil {
		log.Println("handleText: " + err.Error())
	}
}

func request_handler(submission_channel chan *RequestSubmission, result_channel chan *ApiResult) {
	for {
		select {
		case <-quit:
			println("stopping f1")
			return
		case request := <-submission_channel:
			break
		case answer := <-result_channel:
			break
		}
	}
}

func (n *Node) Start() {
	//òancu request goroutine
	for i := 1; i <= 5; i++ {
		request_handler(submission_channel, result_channel)
	}

	var opts = &websocket.ServerOptions{
		Address:    "127.0.0.1:8000",
		HandleBin:  handleBin,
		HandleText: handleText,
	}
	var srv = websocket.NewServer(opts)

	srv.Start()
}
