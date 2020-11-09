package tasks

import "testing"

func TestGenericTask(t *testing.T) {
	description := "Test Generic Task"
	task := NewGenericTask(description)

	if task.Description() != description {
		t.Errorf("task.Description(): \"%s\"; expected: \"%s\"", task.Description(), description)
	}

	if task.Description() != task.String() {
		t.Errorf("task.String(): \"%s\"; expected: \"%s\"", task.String(), description)
	}

	result, err := task.Fire()

	if err != nil {
		t.Errorf("task.Fire(): unexpected error: %s", err)
	}

	expectedResult := "Fired: " + description

	if result.Value() != expectedResult {
		t.Errorf("task.Fire(): result: \"%s\"; expected: \"%s\"", result.Value(), expectedResult)
	}
}
