package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInitCluster(t *testing.T) {
	var config1 = ClusterConfig{
		JoinList: nil,
		BindPort: 8000,
	}

	var config2 = ClusterConfig{
		JoinList: []string{"localhost:8000"},
		BindPort: 8001,
	}

	clusterNode1, err := Init(&config1)
	assert.Nil(t, err)
	assert.NotNil(t, clusterNode1)
	fmt.Printf("Local member %s:%d\n", clusterNode1.node.Addr, clusterNode1.node.Port)

	time.Sleep(time.Duration(1) * time.Second)

	clusterNode2, err := Init(&config2)
	assert.Nil(t, err)
	assert.NotNil(t, clusterNode2)
	fmt.Printf("Local member %s:%d\n", clusterNode2.node.Addr, clusterNode2.node.Port)

	time.Sleep(time.Duration(1) * time.Second)

	joinedNode1 := clusterNode1.GetJoinedNode()
	joinedNode2 := clusterNode2.GetJoinedNode()
	assert.Equal(t, len(joinedNode1), len(joinedNode2))
	assert.NotEqual(t, joinedNode1, joinedNode2)
}
