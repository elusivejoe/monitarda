package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/elusivejoe/monitarda/logging"
	"github.com/elusivejoe/monitarda/polling"
	"github.com/elusivejoe/monitarda/storage"
	"github.com/elusivejoe/monitarda/tasks"
)

func init() {
	logFileName := "monitarda.log"
	logConfig, err := logging.NewFileLogConfig(logrus.InfoLevel, logFileName)

	if err != nil {
		fmt.Printf("Failed to create log file: %s\n", err)
		return
	}

	logging.Configure(logConfig)
}

var logger = logging.GetLogger()

func main() {
	logger.Info("The program has started")

	poller := polling.NewPoller()
	resultsStorage := storage.NewStorage()

	t1 := tasks.NewGenericTask("Task 1")
	t2 := tasks.NewFileTask("testdata/mdstat/dev-md0", "Task 2")

	td1 := poller.Poll(t1, polling.Once, time.Second*5)
	td2 := poller.Poll(t2, polling.Infinite, time.Second*2)

	resultsStorage.AddInput(td1.ResultsChan())
	resultsStorage.AddInput(td2.ResultsChan())

	go func() {
		<-time.Tick(time.Second * 7)
		poller.Unpoll(td2.TaskId())
	}()

	poller.WaitAll()
	resultsStorage.WaitAll()

	logger.Info("The program has finished")
}
