package day05

import "fmt"

type Reading struct {
	NumVents int
}

func (r *Reading) Inc() {
	r.NumVents++
}

// assumes 4 digit max
func (r *Reading) String() string {
	switch {
	case r.NumVents == 0:
		return "   . "
	default:
		return fmt.Sprintf("%4d ", r.NumVents)
	}
}
