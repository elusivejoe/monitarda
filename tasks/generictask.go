package tasks

import (
	"fmt"
)

type GenericTask struct {
	description string
}

func NewGenericTask(description string) *GenericTask {
	return &GenericTask{description: description}
}

func (t *GenericTask) Description() string {
	return t.description
}

func (t *GenericTask) String() string {
	return t.Description()
}

func (t *GenericTask) Fire() (Result, error) {
	return Result{value: fmt.Sprintf("Fired: %s\n", t)}, nil
}
