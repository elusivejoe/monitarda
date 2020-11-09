package storage

import (
	"sync"
)

type FormatterDescriptor struct {
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

func newDescriptor() FormatterDescriptor {
	return FormatterDescriptor{writerId: generateId()}
}

func (wd FormatterDescriptor) WriterId() uint64 {
	return wd.writerId
}
