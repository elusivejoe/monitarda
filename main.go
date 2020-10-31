package main

import (
	"monitarda/polling"
	"monitarda/tasks"
	"time"
)

func main() {
	var poller = polling.NewPoller()

	poller.Poll(tasks.NewGenericTask(tasks.Once, time.Second*7))
	var task2 = poller.Poll(tasks.NewGenericTask(tasks.Infinite, time.Second*1))

	timer := time.NewTimer(time.Second * 5)

	go func() {
		<-timer.C
		poller.Unpoll(task2.TaskId())
	}()

	poller.WaitAllTasks()
}
