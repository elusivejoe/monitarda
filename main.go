package main

import (
	"fmt"
	"monitarda/polling"
	"time"
)

func main() {
	var task1 = polling.CreateTask(polling.Once, time.Second*10)
	var task2 = polling.CreateTask(polling.Infinite, time.Minute)

	fmt.Println(task1)
	fmt.Println(task2)

	var poller = polling.CreatePoller()

	var descriptor1 = poller.Poll(task1)
	var descriptor2 = poller.Poll(task2)

	fmt.Printf("%d %s\n", descriptor1.TaskId(), descriptor1.TaskString())
	fmt.Printf("%d %s\n", descriptor2.TaskId(), descriptor2.TaskString())
}
