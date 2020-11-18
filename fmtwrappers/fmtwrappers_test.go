package fmtwrappers

import (
	"monitarda/tasks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericWrapper(t *testing.T) {
	description := "Test Generic Task"
	wrappedTask := NewGenericWrapper(tasks.NewGenericTask(description))

	assert.Equal(t, description, wrappedTask.Description())
	assert.Equal(t, description, wrappedTask.String())

	result, err := wrappedTask.Fire()
	assert.Nil(t, err)
	assert.Equal(t, "Formatted: Fired: "+description, result.Value())
}
