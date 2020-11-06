package storage

import "fmt"

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Register(writer Writer) {
	go func() {
		for {
			result, ok := <-writer.Channel()

			if !ok {
				break
			}

			fmt.Printf("Store results: %s", result.Result())
		}
	}()
}
