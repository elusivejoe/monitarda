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
	ticker     *time.Ticker
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

func (p *Poller) Poll(t tasks.Task) TaskDescriptor {
	pollMutex.Lock()
	defer pollMutex.Unlock()

	p.waitGroup.Add(1)

	var descriptor = newDescriptor(t)

	ticker := time.NewTicker(t.Interval())
	stopper := make(chan bool)

	p.runRoutine(t, ticker, stopper, descriptor)

	p.tasks[descriptor.taskId] = &polledTask{
		task:       t,
		descriptor: descriptor,
		ticker:     ticker,
		stopper:    stopper}

	return descriptor
}

func (p *Poller) runRoutine(t tasks.Task, ticker *time.Ticker, stopper chan bool, descriptor TaskDescriptor) {
	go func() {
	outerLoop:
		for {
			select {
			case <-ticker.C:
				{
					t.Fire()

					if t.RepeatMode() == tasks.Once {
						break outerLoop
					}
				}
			case <-stopper:
				{
					break outerLoop
				}
			}
		}

		fmt.Printf("Task %d finished\n", descriptor.TaskId())

		p.waitGroup.Done()
	}()
}

func (p *Poller) Unpoll(id uint64) {
	pollMutex.Lock()
	defer pollMutex.Unlock()

	p.tasks[id].ticker.Stop()
	p.tasks[id].stopper <- true

	delete(p.tasks, id)

	fmt.Printf("Task %d unpolled\n", id)
}

func (p *Poller) ListTasks() []TaskDescriptor {
	pollMutex.Lock()
	defer pollMutex.Unlock()

	var allTasks []TaskDescriptor

	for _, v := range p.tasks {
		allTasks = append(allTasks, v.descriptor)
	}

	return allTasks
}

func (p *Poller) WaitAllTasks() {
	p.waitGroup.Wait()
}
