package storage

import (
	"sync"
)

type WriterDescriptor struct {
	writerId uint64
}

//TODO: get rid of global mutable variable
var writerId uint64
var writerIdMutex sync.Mutex

func generateId() uint64 {
	writerIdMutex.Lock()
	defer writerIdMutex.Unlock()

	writerId++
	return writerId
}

func newDescriptor() WriterDescriptor {
	return WriterDescriptor{writerId: generateId()}
}

func (wd WriterDescriptor) WriterId() uint64 {
	return wd.writerId
}
