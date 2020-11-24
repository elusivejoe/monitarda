package polling

import (
	"sync"
	"time"

	"github.com/elusivejoe/monitarda/tasks"
	log "github.com/sirupsen/logrus"
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

func (p *Poller) Poll(task tasks.Task, repeat Repeat, duration time.Duration) TaskDescriptor {
	pollMutex.Lock()
	defer pollMutex.Unlock()

	stopper := make(chan bool)
	results := make(chan tasks.Result)

	descriptor := newDescriptor(repeat, duration, results)

	p.runRoutine(task, descriptor, stopper, results)
	p.tasks[descriptor.taskId] = &polledTask{task: task, descriptor: descriptor, stopper: stopper}

	return descriptor
}

func (p *Poller) runRoutine(t tasks.Task, descriptor TaskDescriptor, stopper chan bool, results chan tasks.Result) {
	p.waitGroup.Add(1)

	ticker := time.NewTicker(descriptor.Duration())

	go func() {
	outerLoop:
		for {
			select {
			case <-ticker.C:
				if result, err := t.Fire(); err != nil {
					log.Errorf("Task %d returned an error: %s", descriptor.taskId, err)
					break outerLoop
				} else {
					select {
					case results <- result:
					default:
					}
				}

				if descriptor.Repeat() == Once {
					break outerLoop
				}
			case <-stopper:
				break outerLoop
			}
		}

		log.Infof("Task %d finished", descriptor.taskId)

		ticker.Stop()
		close(results)

		p.cleanUpTask(descriptor.taskId)
		p.waitGroup.Done()
	}()
}

func (p *Poller) Unpoll(id uint64) {
	pollMutex.Lock()
	defer pollMutex.Unlock()

	task, ok := p.tasks[id]

	if !ok {
		log.Warnf("TaskId %d not found", id)
		return
	}

	task.stopper <- true
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
