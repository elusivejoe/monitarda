package formatters

type Result struct {
	value string
}

func (r Result) Value() string {
	return r.value
}
