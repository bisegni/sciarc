package services

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/shuLhan/share/lib/websocket"
)

// Node...
type Node struct {
	dispatcher *ApiDispatcher
}

// handleBin from websocket by echo-ing back the payload.
func (n *Node) handleBin(conn int, payload []byte) {
	err := n.dispatcher.SubmitRequest(
		&RequestSubmission{
			conn_id:      conn,
			request_type: Binary,
			data:         payload,
		},
	)
	if err != nil {
		log.Printf("Error submitting binary message: %s\n", err)
	}
}

// handleText from websocket by echo-ing back the payload.
func (n *Node) handleText(conn int, payload []byte) {
	var command = string(payload)
	err := n.dispatcher.SubmitRequest(
		&RequestSubmission{
			conn_id:      conn,
			request_type: String,
			data:         &command,
		},
	)
	if err != nil {
		log.Printf("Error submitting text message: %s\n", err)
	}
}

func (n *Node) stringAnswer(conn int, resutl *string) {
	var packet []byte = websocket.NewFrameText(false, []byte(*resutl))
	err := websocket.Send(conn, packet)
	if err != nil {
		log.Println("handleText: " + err.Error())
	}
}

func (n *Node) binaryAnswer(conn int, resutl *[]byte) {
	var packet []byte = websocket.NewFrameText(false, []byte(*resutl))
	err := websocket.Send(conn, packet)
	if err != nil {
		log.Println("handleText: " + err.Error())
	}
}

func (n *Node) Start() error {
	var err error
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	n.dispatcher, err = NewApiDispatcher(1)
	if err != nil {
		return err
	}
	// set the dispatcher
	n.dispatcher.SetStringAnswerHandler(n.stringAnswer)
	n.dispatcher.SetBinaryAnswerHandler(n.binaryAnswer)

	stringEventCallback = n.stringAnswer

	n.dispatcher.RegisterApi(GetApiFactory{})
	n.dispatcher.RegisterApi(StartMonitorApiFactory{})

	var opts = &websocket.ServerOptions{
		Address:    "127.0.0.1:8000",
		HandleBin:  n.handleBin,
		HandleText: n.handleText,
	}
	var srv = websocket.NewServer(opts)
	go func() {
		<-c
		srv.Stop()
	}()

	srv.Start()
	n.dispatcher.Stop()
	return nil
}
