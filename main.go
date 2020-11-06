package main

import (
	"monitarda/fmtwriters"
	"monitarda/polling"
	"monitarda/storage"
	"monitarda/tasks"
	"time"
)

func main() {
	poller := polling.NewPoller()
	storage := storage.NewStorage()

	t1 := tasks.NewGenericTask(tasks.Once, time.Second*7)
	t2 := tasks.NewGenericTask(tasks.Infinite, time.Second*1)
	t3 := tasks.NewGenericTask(tasks.Once, time.Second*10)

	storage.Register(fmtwriters.NewGenericWriter(t1.Channel()))
	storage.Register(fmtwriters.NewGenericWriter(t2.Channel()))
	storage.Register(fmtwriters.NewGenericWriter(t3.Channel()))

	td1 := poller.Poll(t1)
	td2 := poller.Poll(t2)
	poller.Poll(t3)

	go func() {
		<-time.Tick(time.Second * 5)
		poller.Unpoll(td1.TaskId())
		poller.Unpoll(td2.TaskId())
	}()

	poller.WaitAll()
}
