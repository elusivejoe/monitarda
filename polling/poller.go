package polling

import "sync"

type polledTask struct {
	task       Task
	descriptor TaskDescriptor
}

type Poller struct {
	tasks map[uint64]polledTask
}

func CreatePoller() Poller {
	return Poller{tasks: make(map[uint64]polledTask)}
}

var pollMutex sync.Mutex

func (p Poller) Poll(t Task) TaskDescriptor {
	pollMutex.Lock()
	defer pollMutex.Unlock()

	var descriptor = createDescriptor(t)

	p.tasks[descriptor.taskId] = polledTask{task: t, descriptor: descriptor}

	return descriptor
}

func (p Poller) Unpoll(id uint64) {
	pollMutex.Lock()
	defer pollMutex.Unlock()

	delete(p.tasks, id)
}

func (p Poller) ListTasks() []TaskDescriptor {
	pollMutex.Lock()
	defer pollMutex.Unlock()

	var allTasks []TaskDescriptor

	for _, v := range p.tasks {
		allTasks = append(allTasks, v.descriptor)
	}

	return allTasks
}
