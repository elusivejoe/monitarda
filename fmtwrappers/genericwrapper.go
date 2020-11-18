package fmtwrappers

import (
	"fmt"
	"monitarda/tasks"
)

type GenericWrapper struct {
	task tasks.Task
}

func NewGenericWrapper(task tasks.Task) *GenericWrapper {
	return &GenericWrapper{task: task}
}

func (gw *GenericWrapper) Fire() (tasks.Result, error) {

	result, err := gw.task.Fire()

	if err != nil {
		return result, err
	}

	result.SetValue(fmt.Sprintf("Formatted: %s", result.Value()))

	return result, nil
}

func (gw *GenericWrapper) Description() string {
	return gw.task.Description()
}

func (gw *GenericWrapper) String() string {
	return gw.task.Description()
}
