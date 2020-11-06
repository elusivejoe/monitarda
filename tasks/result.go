package tasks

type Result struct {
	result string
}

func (r *Result) Value() string {
	return r.result
}
