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

func TestPollerLogic(t *testing.T) {
	poller := NewPoller()

	t1 := tasks.NewGenericTask("Task 1")
	t2 := tasks.NewGenericTask("Task 2")

	td1 := poller.Poll(t1, Once, time.Second*3)
	td2 := poller.Poll(t2, Infinite, time.Second*2)

	t1Result := "No result"
	t2Result := "No result"

	watchDog := make(chan bool)

	go func() {
		<-time.Tick(time.Second * 5)
		watchDog <- true
	}()

outerLoop:
	for {
		select {
		case result, ok := <-td1.ResultsChan():
			{
				if !ok {
					t1Result = "Error: Task 1"
				} else {
					t1Result = result.Value()
				}

				break outerLoop
			}

		case result, ok := <-td2.ResultsChan():
			{
				if !ok {
					t2Result = "Error: Task 2"
				} else {
					t2Result = result.Value()
				}
			}
		case <-watchDog:
			{
				t1Result = "Poller hanged"
				t2Result = "Poller hanged"

				break outerLoop
			}
		}
	}

	assert.Equal(t, "Fired: Task 1", t1Result)
	assert.Equal(t, "Fired: Task 2", t2Result)
}
