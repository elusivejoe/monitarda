package logging

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once
var logger *logrus.Logger

//TODO: Find a better way to configure the logger (.conf file?); no need to allow to configure it from anywhere
func Configure(conf logConfig) {
	logger := GetLogger()
	logger.SetLevel(conf.level)
	logger.SetOutput(conf.output)
}

func GetLogger() *logrus.Logger {
	once.Do(func() {
		logger = &logrus.Logger{
			Formatter: &logrus.TextFormatter{},
			Level:     logrus.DebugLevel,
			Out:       os.Stdout,
		}
	})

	return logger
}
