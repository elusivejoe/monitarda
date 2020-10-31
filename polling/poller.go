package polling

import (
	"fmt"
	"monitarda/tasks"
	"sync"
	"time"
)

type polledTask struct {
	task       tasks.Task
	descriptor TaskDescriptor
	stopper    chan bool
}

type Poller struct {
	tasks     map[uint64]*polledTask
	waitGroup sync.WaitGroup
}

func NewPoller() *Poller {
	return &Poller{tasks: make(map[uint64]*polledTask)}
}

var pollMutex sync.Mutex

func (p *Poller) Poll(task tasks.Task) TaskDescriptor {
	pollMutex.Lock()
	defer pollMutex.Unlock()

	descriptor := newDescriptor(task)
	stopper := make(chan bool)

	p.runRoutine(task, stopper, descriptor)
	p.tasks[descriptor.taskId] = &polledTask{task: task, descriptor: descriptor, stopper: stopper}

	return descriptor
}

func (p *Poller) runRoutine(t tasks.Task, stopper chan bool, descriptor TaskDescriptor) {
	p.waitGroup.Add(1)

	tickChannel := time.NewTicker(t.Interval()).C

	go func() {
	outerLoop:
		for {
			select {
			case <-tickChannel:
				t.Fire()

				if t.RepeatMode() == tasks.Once {
					break outerLoop
				}

			case <-stopper:
				break outerLoop
			}
		}

		fmt.Printf("Task %d finished\n", descriptor.TaskId())

		p.cleanUpTask(descriptor.TaskId())
		p.waitGroup.Done()
	}()
}

func (p *Poller) Unpoll(id uint64) {
	pollMutex.Lock()
	defer pollMutex.Unlock()

	p.tasks[id].stopper <- true
}

func (p *Poller) cleanUpTask(id uint64) {
	delete(p.tasks, id)
}

func (p *Poller) ListTasks() []TaskDescriptor {
	pollMutex.Lock()
	defer pollMutex.Unlock()

	var allTasks = make([]TaskDescriptor, len(p.tasks))

	for _, v := range p.tasks {
		allTasks = append(allTasks, v.descriptor)
	}

	return allTasks
}

func (p *Poller) WaitAll() {
	p.waitGroup.Wait()
}
