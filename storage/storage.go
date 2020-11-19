package storage

import (
	"fmt"
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

func (s *Storage) AddInput(inputChan <-chan tasks.Result) InputDescriptor {
	storageMutex.Lock()
	defer storageMutex.Unlock()

	s.waitGroup.Add(1)
	descriptor := newDescriptor()

	stopper := make(chan bool)
	s.stoppers[descriptor.InputId()] = stopper

	go func() {
	outerLoop:
		for {
			select {
			case result, ok := <-inputChan:
				{
					if !ok {
						fmt.Println("Input channel has been closed")
						break outerLoop
					}

					if err := s.storeResult(result); err != nil {
						fmt.Printf("Failed to store result: %s", result.Value())
						break outerLoop
					}
				}
			case <-stopper:
				{
					break outerLoop
				}
			}
		}

		fmt.Printf("Input %d has finished the job\n", descriptor.InputId())
		delete(s.stoppers, descriptor.InputId())
		s.waitGroup.Done()
	}()

	return descriptor
}

func (s *Storage) storeResult(result tasks.Result) error {
	fmt.Printf("Store: %s\n", result.Value())
	return nil
}

func (s *Storage) RemoveInput(inputId uint64) {
	storageMutex.Lock()
	defer storageMutex.Unlock()

	stopper, ok := s.stoppers[inputId]

	if !ok {
		fmt.Printf("InputId %d not found", inputId)
		return
	}

	stopper <- true
	delete(s.stoppers, inputId)
}

func (s *Storage) WaitAll() {
	s.waitGroup.Wait()
}
