package polling

type Repeat uint8

const (
	Once     Repeat = iota
	Infinite Repeat = iota
)

func (r Repeat) String() string {
	switch r {
	case Once:
		return "Once"

	case Infinite:
		return "Infinite"

	default:
		return string(r)
	}
}
