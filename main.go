package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	readSmart()
	runMonitoring(10 * time.Second)
	runServer()
}

var mutSmart sync.Mutex
var resultSmart string

func handlerSmart(out http.ResponseWriter, request *http.Request) {
	mutSmart.Lock()
	defer mutSmart.Unlock()

	fmt.Fprintf(out, "Smart: %s", resultSmart)
}

func readSmart() {
	mutSmart.Lock()
	defer mutSmart.Unlock()

	resultSmart = time.Now().String()
}

func runMonitoring(interval time.Duration) {
	tickChan := time.Tick(interval)

	go func() {
		for {
			<-tickChan
			readSmart()
		}
	}()
}

func runServer() {
	http.HandleFunc("/smart", handlerSmart)

	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}
