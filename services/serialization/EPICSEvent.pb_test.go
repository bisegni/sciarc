package serialization

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestSerialization(t *testing.T) {
	value := &PayloadInfo{
		Type:   PayloadType_SCALAR_STRING,
		Pvname: "test",
		Headers: []*FieldValue{
			&FieldValue{
				Name: "test",
				Val:  "10",
			},
			&FieldValue{
				Name: "test2",
				Val:  "100",
			},
		},
	}
	out, err := proto.Marshal(value)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, len(out), 0)
}
