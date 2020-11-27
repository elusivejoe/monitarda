package polling

import (
	"fmt"
	"sync"
	"time"

	"github.com/elusivejoe/monitarda/tasks"
)

type TaskDescriptor struct {
	taskId      uint64
	repeat      Repeat
	duration    time.Duration
	resultsChan <-chan tasks.Result
}

//TODO: get rid of global mutable variable
var taskId uint64
var taskIdMutex sync.Mutex

func generateId() uint64 {
	taskIdMutex.Lock()
	defer taskIdMutex.Unlock()

	taskId++
	return taskId
}

func newDescriptor(repeat Repeat, duration time.Duration, resultsChan <-chan tasks.Result) TaskDescriptor {
	return TaskDescriptor{taskId: generateId(), repeat: repeat, duration: duration, resultsChan: resultsChan}
}

func (td TaskDescriptor) TaskId() uint64 {
	return td.taskId
}

func (td TaskDescriptor) ResultsChan() <-chan tasks.Result {
	return td.resultsChan
}

func (td TaskDescriptor) Repeat() Repeat {
	return td.repeat
}

func (td TaskDescriptor) Duration() time.Duration {
	return td.duration
}

func (td TaskDescriptor) String() string {
	return fmt.Sprintf("Id: %d Repeat: %s, Duration: %s", td.TaskId(), td.Repeat(), td.Duration())
}
