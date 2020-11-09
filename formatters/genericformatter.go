package formatters

import (
	"monitarda/tasks"
)

type GenericFormatter struct{}

func NewGenericFormatter() *GenericFormatter {
	return &GenericFormatter{}
}

func (gf *GenericFormatter) Format(taskResult tasks.Result) (Result, error) {
	return Result{value: "Formatted: " + taskResult.Value()}, nil
}
