package storage

import (
	"fmt"
	"sync"
)

type InputDescriptor struct {
	inputId uint64
}

//TODO: get rid of global mutable variable
var inputId uint64
var inputIdMutex sync.Mutex

func generateId() uint64 {
	inputIdMutex.Lock()
	defer inputIdMutex.Unlock()

	inputId++
	return inputId
}

func newDescriptor() InputDescriptor {
	return InputDescriptor{inputId: generateId()}
}

func (inDesc InputDescriptor) InputId() uint64 {
	return inDesc.inputId
}

func (inDesc InputDescriptor) String() string {
	return fmt.Sprintf("Id: %d", inDesc.InputId())
}
