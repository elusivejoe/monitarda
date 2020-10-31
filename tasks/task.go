package tasks

import (
	"time"
)

type Repeat uint8

const (
	Once     Repeat = iota
	Infinite Repeat = iota
)

type Task interface {
	Fire()
	Interval() time.Duration
	RepeatMode() Repeat
	String() string
}
