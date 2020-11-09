package formatters

import (
	"monitarda/tasks"
	"testing"
)

func TestGenericFormatter(t *testing.T) {
	description := "Test Generic Task"
	task := tasks.NewGenericTask(description)

	formatter := NewGenericFormatter()

	result, err := task.Fire()

	if err != nil {
		t.Errorf("task.Fire(): unexpected error: %s", err)
	}

	formatted, err := formatter.Format(result)

	if err != nil {
		t.Errorf("formatter.Format(): unexpected error: %s", err)
	}

	expectedFormatted := "Formatted: Fired: " + description

	if formatted.Value() != expectedFormatted {
		t.Errorf("formatter.Format(): result: \"%s\"; expected: \"%s\"", formatted.Value(), expectedFormatted)
	}
}
