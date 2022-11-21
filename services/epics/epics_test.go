package epics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	value, err := GetChannelvalue("variable:sum")
	assert.NotNil(t, err, "There not be an error")
	assert.Equal(t, len(value) > 0, true, "The value need to be filled")
}
