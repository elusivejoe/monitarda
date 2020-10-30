package polling

import "sync"

type TaskDescriptor struct {
	taskId     uint64
	taskString string
}

var taskId uint64
var taskIdMutex sync.Mutex

func generateId() uint64 {
	taskIdMutex.Lock()
	defer taskIdMutex.Unlock()

	taskId++
	return taskId
}

func createDescriptor(t Task) TaskDescriptor {
	return TaskDescriptor{taskId: generateId(), taskString: t.String()}
}

func (td TaskDescriptor) TaskId() uint64 {
	return td.taskId
}

func (td TaskDescriptor) TaskString() string {
	return td.taskString
}
