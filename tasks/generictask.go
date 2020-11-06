package tasks

import (
	"fmt"
	"time"
)

type GenericTask struct {
	repeat   Repeat
	interval time.Duration
	channel  chan Result
}

func NewGenericTask(repeat Repeat, interval time.Duration) *GenericTask {
	return &GenericTask{repeat: repeat, interval: interval, channel: make(chan Result)}
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
	t.channel <- Result{result: fmt.Sprintf("Fired: %s\n", t)}
}

func (t *GenericTask) Channel() <-chan Result {
	return t.channel
}
