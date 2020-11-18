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
	assert.Equal(t, "Personalities : [raid1] \n"+
		"md0 : active raid1 sda[0] sdb[1]\n"+
		"      13672250368 blocks super 1.2 [2/2] [UU]\n"+
		"      bitmap: 0/102 pages [0KB], 65536KB chunk\n"+
		"\nunused devices: <none>\n", result.Value())
}

func TestResult(t *testing.T) {
	result := Result{"Test Value"}
	assert.Equal(t, "Test Value", result.Value())

	result.SetValue("New Test Value")
	assert.Equal(t, "New Test Value", result.Value())
}
