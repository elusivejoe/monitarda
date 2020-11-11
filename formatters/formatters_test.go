package formatters

import (
	"monitarda/tasks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericFormatter(t *testing.T) {
	description := "Test Generic Task"
	task := tasks.NewGenericTask(description)
	formatter := NewGenericFormatter()

	result, err := task.Fire()
	assert.Nil(t, err)

	formatted, err := formatter.Format(result)
	assert.Nil(t, err)
	assert.Equal(t, "Formatted: Fired: "+description, formatted.Value())
}
