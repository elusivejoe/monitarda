package polling

import (
	"fmt"
	"time"
)

type Repeat uint8

const (
	Once     Repeat = iota
	Infinite Repeat = iota
)

type Task struct {
	repeat   Repeat
	interval time.Duration
}

func CreateTask(repeat Repeat, interval time.Duration) Task {
	return Task{repeat: repeat, interval: interval}
}

func (t Task) Repeat() Repeat {
	return t.repeat
}

func (t Task) Interval() time.Duration {
	return t.interval
}

func (t Task) String() string {
	return fmt.Sprintf("%d %s", t.Repeat(), t.Interval())
}
