package fmtwriters

import (
	"monitarda/storage"
	"monitarda/tasks"
)

type GenericWriter struct {
	inChannel  <-chan tasks.Result
	outChannel <-chan storage.Result
}

func NewGenericWriter(inChan <-chan tasks.Result) *GenericWriter {
	outChannel := make(chan storage.Result)

	go func() {
		for {
			res, ok := <-inChan

			if !ok {
				close(outChannel)
				break
			}

			outChannel <- storage.NewResult("Formatted: " + res.Value())
		}
	}()
	return &GenericWriter{inChannel: inChan, outChannel: outChannel}
}

func (gw *GenericWriter) Channel() <-chan storage.Result {
	return gw.outChannel
}
