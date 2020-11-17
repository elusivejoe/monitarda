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

func TestFileTask(t *testing.T) {
	description := "Test File Task"
	task := NewFileTask("../testdata/mdstat/dev-md0", description)

	assert.Equal(t, description, task.Description())
	assert.Equal(t, description, task.String())

	result, err := task.Fire()
	assert.Nil(t, err)
	assert.Equal(t, "Personalities : [raid1] \u000A"+
		"md0 : active raid1 sda[0] sdb[1]\u000A"+
		"      13672250368 blocks super 1.2 [2/2] [UU]\u000A"+
		"      bitmap: 0/102 pages [0KB], 65536KB chunk\u000A"+
		"\u000Aunused devices: <none>\u000A", result.Value())
}
