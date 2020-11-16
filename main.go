package main

import (
	"monitarda/formatters"
	"monitarda/polling"
	"monitarda/storage"
	"monitarda/tasks"
	"time"
)

func main() {
	poller := polling.NewPoller()
	resultsStorage := storage.NewStorage()

	t1 := tasks.NewGenericTask("Task 1")
	t2 := tasks.NewFileTask("testdata/mdstat/dev-md0", "Task 2")

	td1 := poller.Poll(t1, polling.Once, time.Second*5)
	td2 := poller.Poll(t2, polling.Infinite, time.Second*2)

	resultsStorage.Register(formatters.NewGenericFormatter(), td1.ResultsChan())
	resultsStorage.Register(formatters.NewGenericFormatter(), td2.ResultsChan())

	go func() {
		<-time.Tick(time.Second * 7)
		poller.Unpoll(td2.TaskId())
	}()

	poller.WaitAll()
	resultsStorage.WaitAll()
}
