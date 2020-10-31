package tasks

import (
	"fmt"
	"time"
)

type GenericTask struct {
	repeat   Repeat
	interval time.Duration
}

func NewGenericTask(repeat Repeat, interval time.Duration) *GenericTask {
	return &GenericTask{repeat: repeat, interval: interval}
}

func (t *GenericTask) RepeatMode() Repeat {
	return t.repeat
}

func (t *GenericTask) Interval() time.Duration {
	return t.interval
}

func (t *GenericTask) String() string {
	return fmt.Sprintf("%d %s", t.RepeatMode(), t.Interval())
}

func (t *GenericTask) Fire() {
	fmt.Printf("Fired: %s\n", t)
}
