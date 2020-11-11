package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericTask(t *testing.T) {
	description := "Test Generic Task"
	task := NewGenericTask(description)
	assert.Equal(t, description, task.Description())
	assert.Equal(t, description, task.String())

	result, err := task.Fire()
	assert.Nil(t, err)
	assert.Equal(t, "Fired: "+description, result.Value())
}
