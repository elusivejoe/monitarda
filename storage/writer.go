package storage

type Writer interface {
	Channel() <-chan Result
}
