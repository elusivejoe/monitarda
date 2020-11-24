package storage

import (
	"sync"

	"github.com/elusivejoe/monitarda/tasks"
	log "github.com/sirupsen/logrus"
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
						log.Debugf("Input %d has been closed", descriptor.InputId())
						break outerLoop
					}

					if err := s.storeResult(result); err != nil {
						log.Errorf("Failed to store result: %s", result.Value())
						break outerLoop
					}
				}
			case <-stopper:
				{
					break outerLoop
				}
			}
		}

		log.Infof("Input %d has finished the job", descriptor.InputId())
		delete(s.stoppers, descriptor.InputId())
		s.waitGroup.Done()
	}()

	return descriptor
}

func (s *Storage) storeResult(result tasks.Result) error {
	log.Debugf("Store: %s", result.Value())
	return nil
}

func (s *Storage) RemoveInput(inputId uint64) {
	storageMutex.Lock()
	defer storageMutex.Unlock()

	stopper, ok := s.stoppers[inputId]

	if !ok {
		log.Warnf("InputId %d not found", inputId)
		return
	}

	stopper <- true
	delete(s.stoppers, inputId)
}

func (s *Storage) WaitAll() {
	s.waitGroup.Wait()
}
