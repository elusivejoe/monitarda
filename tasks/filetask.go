package tasks

import (
	"io/ioutil"
)

type FileTask struct {
	path        string
	description string
}

func NewFileTask(path string, description string) *FileTask {
	return &FileTask{path: path, description: description}
}

func (t *FileTask) Description() string {
	return t.description
}

func (t *FileTask) String() string {
	return t.Description()
}

func (t *FileTask) Fire() (Result, error) {
	data, err := ioutil.ReadFile(t.path)
	return Result{value: string(data)}, err
}
