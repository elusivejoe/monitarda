package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type logConfig struct {
	level  logrus.Level
	output io.Writer
}

func NewFileLogConfig(level logrus.Level, filename string) (logConfig, error) {
	logfile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return logConfig{}, err
	}

	return logConfig{level: level, output: logfile}, nil
}
