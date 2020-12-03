package logging

import (
	"fmt"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once
var logger *logrus.Logger

func GetLogger() *logrus.Logger {
	once.Do(func() {
		logger = &logrus.Logger{
			Hooks:        nil,
			ReportCaller: false,
			ExitFunc:     nil,
		}

		logger.SetLevel(logrus.InfoLevel)
		logger.SetFormatter(&logrus.TextFormatter{})

		logfile, err := os.OpenFile("monitarda.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

		if err != nil {
			fmt.Printf("error opening log file: %v", err)
			return
		}

		logger.SetOutput(logfile)
	})

	return logger
}
