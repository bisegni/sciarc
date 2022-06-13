package services

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/hashicorp/memberlist"
)

// Message cluster exange message
type Message struct {
	Action string // add, del
	Data   map[string]string
}

// NodeClusterMetadata
type NodeClusterMetadata struct {
	// memberlist instance
	m *memberlist.Memberlist
	// instance memberlist node information
	node       *memberlist.Node
	name       string
	mtx        sync.RWMutex
	items      map[string]string
	broadcasts *memberlist.TransmitLimitedQueue
	metadata   []byte `default:[]byte{}`
}

func (n *NodeClusterMetadata) NodeMeta(limit int) []byte {
	return n.metadata
}

func (n *NodeClusterMetadata) NotifyMsg(b []byte) {
	if len(b) == 0 {
		return
	}

	switch b[0] {
	case 'd': // data
		var messages []*Message
		if err := json.Unmarshal(b[1:], &messages); err != nil {
			return
		}
		n.mtx.Lock()
		for _, u := range messages {
			for k, v := range u.Data {
				switch u.Action {
				case "add":
					n.items[k] = v
				case "del":
					delete(n.items, k)
				}
			}
		}
		n.mtx.Unlock()
	}
}

func (n *NodeClusterMetadata) GetBroadcasts(overhead, limit int) [][]byte {
	return n.broadcasts.GetBroadcasts(overhead, limit)
}

func (n *NodeClusterMetadata) LocalState(join bool) []byte {
	n.mtx.RLock()
	m := n.items
	n.mtx.RUnlock()
	b, _ := json.Marshal(m)
	return b
}

func (n *NodeClusterMetadata) MergeRemoteState(buf []byte, join bool) {
	if len(buf) == 0 {
		return
	}
	if !join {
		return
	}
	var m map[string]string
	if err := json.Unmarshal(buf, &m); err != nil {
		return
	}
	n.mtx.Lock()
	for k, v := range m {
		n.items[k] = v
	}
	n.mtx.Unlock()
}

func (n *NodeClusterMetadata) NotifyJoin(node *memberlist.Node) {
	fmt.Println("A node has joined: " + node.String())
}

func (n *NodeClusterMetadata) NotifyLeave(node *memberlist.Node) {
	fmt.Println("A node has left: " + node.String())
}

func (n *NodeClusterMetadata) NotifyUpdate(node *memberlist.Node) {
	fmt.Println("A node was updated: " + node.String())
}

// CLuster engine configuration
type ClusterConfig struct {
	JoinList []string
	BindPort int
}

// Init cluster engine
func Init(config *ClusterConfig) (*NodeClusterMetadata, error) {
	var err error
	var clusteNode = NodeClusterMetadata{}

	hostname, _ := os.Hostname()
	clusteNode.name = hostname + "-" + uuid.NewString()

	c := memberlist.DefaultLANConfig()
	c.Events = &clusteNode
	c.Delegate = &clusteNode
	c.BindPort = config.BindPort
	c.Name = clusteNode.name
	clusteNode.m, err = memberlist.Create(c)
	if err != nil {
		return nil, err
	}

	if config != nil && len(config.JoinList) > 0 {
		_, err := clusteNode.m.Join(config.JoinList)
		if err != nil {
			panic("Failed to join cluster: " + err.Error())
		}
	}

	clusteNode.broadcasts = &memberlist.TransmitLimitedQueue{
		NumNodes: func() int {
			return clusteNode.m.NumMembers()
		},
		RetransmitMult: 3,
	}
	clusteNode.node = clusteNode.m.LocalNode()
	return &clusteNode, nil
}
