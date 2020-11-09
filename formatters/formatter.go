package formatters

import "monitarda/tasks"

type Formatter interface {
	Format(result tasks.Result) (Result, error)
}
