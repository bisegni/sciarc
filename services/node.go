package services

import (
	"log"

	"github.com/shuLhan/share/lib/websocket"
)

// Node...
type Node struct {
	dispatcher *ApiDispatcher
}

// handleBin from websocket by echo-ing back the payload.
func (n *Node) handleBin(conn int, payload []byte) {
	err := n.dispatcher.SubmitBinaryRequest(payload)
	if err != nil {
		log.Printf("Error submitting binary message: @", err)
	}
}

// handleText from websocket by echo-ing back the payload.
func (n *Node) handleText(conn int, payload []byte) {
	err := n.dispatcher.SubmitStringRequest(string(payload))
	if err != nil {
		log.Printf("Error submitting text message: @", err)
	}

	//var value, _ = epics.GetChannelvalue(string(payload))
	// var packet []byte = websocket.NewFrameText(false, []byte(value))
	// err = websocket.Send(conn, packet)
	// if err != nil {
	// 	log.Println("handleText: " + err.Error())
	// }
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
	n.dispatcher, err = NewApiDispatcher(4)
	if err != nil {
		return err
	}
	// set the dispatcher
	n.dispatcher.SetStringAnswerHandler(n.stringAnswer)
	n.dispatcher.SetBinaryAnswerHandler(n.binaryAnswer)

	var opts = &websocket.ServerOptions{
		Address:    "127.0.0.1:8000",
		HandleBin:  n.handleBin,
		HandleText: n.handleText,
	}
	var srv = websocket.NewServer(opts)

	srv.Start()
	return nil
}
