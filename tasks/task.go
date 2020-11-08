package tasks

type Task interface {
	Fire() (Result, error)
	Description() string
}
