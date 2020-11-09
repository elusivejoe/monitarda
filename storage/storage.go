package storage

import (
	"fmt"
	"monitarda/formatters"
	"monitarda/tasks"
	"sync"
)

type Storage struct {
	stoppers  map[uint64]chan bool
	waitGroup sync.WaitGroup
}

var storageMutex sync.Mutex

func NewStorage() *Storage {
	return &Storage{stoppers: make(map[uint64]chan bool)}
}

func (s *Storage) Register(formatter formatters.Formatter, inputChan <-chan tasks.Result) FormatterDescriptor {
	storageMutex.Lock()
	defer storageMutex.Unlock()

	s.waitGroup.Add(1)
	descriptor := newDescriptor()

	stopper := make(chan bool)
	s.stoppers[descriptor.WriterId()] = stopper

	go func() {
	outerLoop:
		for {
			select {
			case result, ok := <-inputChan:
				{
					if !ok {
						break outerLoop
					}

					fmt.Printf("Store results: %s", formatter.Format(result).Value())
				}
			case <-stopper:
				{
					break outerLoop
				}
			}
		}

		fmt.Printf("Formatter %d has finished the job\n", descriptor.WriterId())
		delete(s.stoppers, descriptor.WriterId())
		s.waitGroup.Done()
	}()

	return descriptor
}

func (s *Storage) Unregister(writerId uint64) {
	storageMutex.Lock()
	defer storageMutex.Unlock()

	stopper, ok := s.stoppers[writerId]

	if !ok {
		fmt.Printf("WriterId %d not found", writerId)
		return
	}

	stopper <- true
	delete(s.stoppers, writerId)
}

func (s *Storage) WaitAll() {
	s.waitGroup.Wait()
}
