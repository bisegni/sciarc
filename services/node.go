package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/centrifugal/centrifuge"
)

func waitExitSignal(n *centrifuge.Node) {
	sigCh := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		_ = n.Shutdown(context.Background())
		done <- true
	}()
	<-done
}

// Node...
type Node struct {
	name string
	age  int
}

func (n *Node) Start() {
	// Node is the core object in Centrifuge library responsible for
	// many useful things. For example Node allows publishing messages
	// into channels with its Publish method. Here we initialize Node
	// with Config which has reasonable defaults for zero values.
	node, err := centrifuge.New(centrifuge.Config{
		LogLevel: centrifuge.LogLevelDebug,
	})
	if err != nil {
		log.Fatal(err)
	}

	broker, _ := centrifuge.NewMemoryBroker(node, centrifuge.MemoryBrokerConfig{
		HistoryMetaTTL: 120 * time.Second,
	})
	node.SetBroker(broker)

	node.OnConnecting(func(ctx context.Context, e centrifuge.ConnectEvent) (centrifuge.ConnectReply, error) {
		// We need to apply token parsing logic here and return connection credentials.
		if !strings.HasPrefix(e.Token, "I am ") {
			return centrifuge.ConnectReply{}, centrifuge.DisconnectInvalidToken
		}
		userID := strings.TrimPrefix(e.Token, "I am ")
		credentials := &centrifuge.Credentials{
			UserID:   userID,
			ExpireAt: time.Now().Unix() + 5, // Expire in 5 seconds.
		}

		return centrifuge.ConnectReply{
			ClientSideRefresh: true, // This is required to use client-side refresh.
			Credentials:       credentials,
		}, nil
	})
	// Set ConnectHandler called when client successfully connected to Node.
	// Your code inside a handler must be synchronized since it will be called
	// concurrently from different goroutines (belonging to different client
	// connections). See information about connection life cycle in library readme.
	// This handler should not block – so do minimal work here, set required
	// connection event handlers and return.
	node.OnConnect(func(client *centrifuge.Client) {
		// In our example transport will always be Websocket but it can also be SockJS.
		transportName := client.Transport().Name()
		// In our example clients connect with JSON protocol but it can also be Protobuf.
		transportProto := client.Transport().Protocol()
		fmt.Printf("client connected via %s (%s)", transportName, transportProto)

		// Set SubscribeHandler to react on every channel subscription attempt
		// initiated by a client. Here you can theoretically return an error or
		// disconnect a client from a server if needed. But here we just accept
		// all subscriptions to all channels. In real life you may use a more
		// complex permission check here. The reason why we use callback style
		// inside client event handlers is that it gives a possibility to control
		// operation concurrency to developer and still control order of events.
		client.OnSubscribe(func(e centrifuge.SubscribeEvent, cb centrifuge.SubscribeCallback) {
			fmt.Printf("client subscribes on channel %s", e.Channel)
			cb(centrifuge.SubscribeReply{}, nil)
		})

		// By default, clients can not publish messages into channels. By setting
		// PublishHandler we tell Centrifuge that publish from a client-side is
		// possible. Now each time client calls publish method this handler will be
		// called and you have a possibility to validate publication request. After
		// returning from this handler Publication will be published to a channel and
		// reach active subscribers with at most once delivery guarantee. In our simple
		// chat app we allow everyone to publish into any channel but in real case
		// you may have more validation.
		client.OnPublish(func(e centrifuge.PublishEvent, cb centrifuge.PublishCallback) {
			fmt.Printf("client publishes into channel %s: %s", e.Channel, string(e.Data))
			cb(centrifuge.PublishReply{}, nil)
		})

		// Set Disconnect handler to react on client disconnect events.
		client.OnDisconnect(func(e centrifuge.DisconnectEvent) {
			fmt.Printf("client disconnected")
		})
	})

	// Run node. This method does not block. See also node.Shutdown method
	// to finish application gracefully.
	if err := node.Run(); err != nil {
		log.Fatal(err)
	}

	// Now configure HTTP routes.

	// Serve Websocket connections using WebsocketHandler.
	websocketHandler := centrifuge.NewWebsocketHandler(node, centrifuge.WebsocketConfig{
		ReadBufferSize:     1024,
		UseWriteBufferPool: true,
	})
	http.Handle("/connection/websocket", websocketHandler)

	// The second route is for serving index.html file.
	//http.Handle("/", http.FileServer(http.Dir("./")))

	go func() {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			fmt.Print(err)
		}
	}()

	waitExitSignal(node)
}
