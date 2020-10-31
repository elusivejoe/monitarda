package main

import (
	"monitarda/polling"
	"monitarda/tasks"
	"time"
)

func main() {
	var poller = polling.NewPoller()

	task1 := poller.Poll(tasks.NewGenericTask(tasks.Once, time.Second*7))
	task2 := poller.Poll(tasks.NewGenericTask(tasks.Infinite, time.Second*1))
	poller.Poll(tasks.NewGenericTask(tasks.Once, time.Second*10))

	go func() {
		<-time.NewTimer(time.Second * 5).C
		poller.Unpoll(task1.TaskId())
		poller.Unpoll(task2.TaskId())
	}()

	poller.WaitAll()
}
