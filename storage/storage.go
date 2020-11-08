package storage

import (
	"fmt"
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

func (s *Storage) Register(writer Writer) WriterDescriptor {
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
			case result, ok := <-writer.Channel():
				{
					if !ok {
						break outerLoop
					}

					fmt.Printf("Store results: %s", result.Result())
				}
			case <-stopper:
				{
					break outerLoop
				}
			}
		}

		fmt.Printf("Writer %d has finished the job\n", descriptor.WriterId())
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
