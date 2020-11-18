package tasks

type Result struct {
	value string
}

func (r *Result) Value() string {
	return r.value
}

func (r *Result) SetValue(newVal string) {
	r.value = newVal
}
