package services

import (
	"log"

	"github.com/shuLhan/share/lib/debug"
	"github.com/shuLhan/share/lib/websocket"
)

// Node...
type Node struct {
	name string
	age  int
}

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
		packet []byte = websocket.NewFrameText(false, payload)
		err    error
	)

	if debug.Value >= 3 {
		log.Printf("testdata/server: handleText: {payload.len:%d}\n", len(payload))
	}

	err = websocket.Send(conn, packet)
	if err != nil {
		log.Println("handleText: " + err.Error())
	}
}

func (n *Node) Start() {
	var opts = &websocket.ServerOptions{
		Address:    "127.0.0.1:8000",
		HandleBin:  handleBin,
		HandleText: handleText,
	}
	var srv = websocket.NewServer(opts)

	srv.Start()
}
