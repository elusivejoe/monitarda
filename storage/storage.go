package storage

import (
	"fmt"
	"sync"

	"github.com/elusivejoe/monitarda/logging"
	"github.com/elusivejoe/monitarda/tasks"
)

type Storage struct {
	stoppers  map[uint64]chan bool
	waitGroup sync.WaitGroup
}

var storageMutex sync.Mutex
var logger = logging.GetLogger()

func NewStorage() *Storage {
	return &Storage{stoppers: make(map[uint64]chan bool)}
}

func (s *Storage) AddInput(inputChan <-chan tasks.Result) InputDescriptor {
	storageMutex.Lock()
	defer storageMutex.Unlock()

	s.waitGroup.Add(1)
	descriptor := newDescriptor()

	logger.Infof("Add input: %s", descriptor)

	stopper := make(chan bool)
	s.stoppers[descriptor.InputId()] = stopper

	go func() {
	outerLoop:
		for {
			select {
			case result, ok := <-inputChan:
				{
					if !ok {
						logger.Debugf("Input has been closed: %s", descriptor)
						break outerLoop
					}

					if err := s.storeResult(result); err != nil {
						logger.Errorf("Input: %s Failed to store result: %s", descriptor, err)
						break outerLoop
					}
				}
			case <-stopper:
				{
					break outerLoop
				}
			}
		}

		logger.Infof("Input %s has finished the job", descriptor)
		delete(s.stoppers, descriptor.InputId())
		s.waitGroup.Done()
	}()

	return descriptor
}

func (s *Storage) storeResult(result tasks.Result) error {
	logger.Debugf("Store: %s", result)
	return nil
}

func (s *Storage) RemoveInput(inputId uint64) error {
	storageMutex.Lock()
	defer storageMutex.Unlock()

	stopper, ok := s.stoppers[inputId]

	if !ok {
		return fmt.Errorf("input with Id %d not found", inputId)
	}

	stopper <- true
	delete(s.stoppers, inputId)

	return nil
}

func (s *Storage) WaitAll() {
	s.waitGroup.Wait()
}
