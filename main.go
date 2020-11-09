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
	t2 := tasks.NewGenericTask("Task 2")
	t3 := tasks.NewGenericTask("Task 3")

	td1 := poller.Poll(t1, polling.Once, time.Second*10)
	td2 := poller.Poll(t2, polling.Infinite, time.Second*1)
	td3 := poller.Poll(t3, polling.Once, time.Second*15)

	resultsStorage.Register(formatters.NewGenericFormatter(), td1.ResultsChan())
	wd2 := resultsStorage.Register(formatters.NewGenericFormatter(), td2.ResultsChan())
	resultsStorage.Register(formatters.NewGenericFormatter(), td3.ResultsChan())

	go func() {
		<-time.Tick(time.Second * 5)
		resultsStorage.Unregister(wd2.WriterId())
		poller.Unpoll(td1.TaskId())
	}()

	go func() {
		<-time.Tick(time.Second * 10)
		poller.Unpoll(td2.TaskId())
	}()

	poller.WaitAll()
	resultsStorage.WaitAll()
}
