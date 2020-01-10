package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent_TableName(t *testing.T) {
	event := &Event{Location: "c47a85a64c"}
	tableName := event.TableName()

	assert.Equal(t, "events", tableName)
}
