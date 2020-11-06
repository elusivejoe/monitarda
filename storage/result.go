package storage

type Result struct {
	result string
}

func NewResult(str string) Result {
	return Result{result: str}
}

func (r *Result) Result() string {
	return r.result
}
