package polling

import (
	"monitarda/tasks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTaskDescriptor(t *testing.T) {
	poller := NewPoller()

	t1 := tasks.NewGenericTask("Task 1")
	t2 := tasks.NewGenericTask("Task 2")

	assert.Equal(t, "Task 1", t1.Description())
	assert.Equal(t, "Task 2", t2.Description())

	td1 := poller.Poll(t1, Once, time.Second*5)
	td2 := poller.Poll(t2, Infinite, time.Minute*10)

	assert.Equal(t, uint64(1), td1.TaskId())
	assert.Equal(t, uint64(2), td2.TaskId())

	assert.Equal(t, Once, td1.Repeat())
	assert.Equal(t, Infinite, td2.Repeat())

	assert.Equal(t, time.Second*5, td1.Duration())
	assert.Equal(t, time.Minute*10, td2.Duration())
}
