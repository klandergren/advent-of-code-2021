package day08

import (
	"sort"
	"strings"
)

var SignalValueLookup = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

var SegmentLookup = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
}

type Digit struct {
	InputSignal string
	WireSignal  []int
}

func NewDigitFromInput(inputSignal string) *Digit {
	wireSignal := make([]int, len(SegmentLookup))

	for s, i := range SegmentLookup {
		if strings.Contains(inputSignal, s) {
			wireSignal[i] = 1
		}
	}

	n := strings.Split(inputSignal, "")
	sort.Strings(n)
	sorted := strings.Join(n, "")

	return &Digit{
		InputSignal: sorted,
		WireSignal:  wireSignal,
	}
}

func NewDigitFromWire(wireSignal []int) *Digit {
	var sb strings.Builder

	for s, v := range SegmentLookup {
		if wireSignal[v] == 1 {
			sb.WriteString(string(s))
		}
	}
	return &Digit{
		InputSignal: sb.String(),
		WireSignal:  wireSignal,
	}
}

func (d *Digit) RawInputSignals() []string {
	return strings.Split(d.InputSignal, "")
}

// return a *Digit containing all wire signals in `d`, zero'ing out any non-zero signals also found in `other`
func (d *Digit) Subtract(other *Digit) *Digit {
	newWireSignal := make([]int, len(SegmentLookup))
	copy(newWireSignal, d.WireSignal)

	// for every signal value that `other` also has, zero out
	for i, sv := range other.WireSignal {
		// don't care about zero values
		if sv == 0 {
			continue
		}

		// also don't care when we don't have a value
		if newWireSignal[i] == 0 {
			continue
		}

		// we have a value, other has a value, so zero out
		newWireSignal[i] = 0
	}

	return NewDigitFromWire(newWireSignal)
}

// return a *Digit containing all signals present in both `d` and `other`
func (d *Digit) Intersect(other *Digit) *Digit {
	newWireSignal := make([]int, 7)

	for i, sv := range other.WireSignal {
		if sv == 1 && d.WireSignal[i] == 1 {
			newWireSignal[i] = 1
		}
	}

	return NewDigitFromWire(newWireSignal)
}

func (d *Digit) Value() int {
	if len(d.InputSignal) == 2 {
		return 1
	}
	if len(d.InputSignal) == 3 {
		return 7
	}
	if len(d.InputSignal) == 4 {
		return 4
	}
	if len(d.InputSignal) == 7 {
		return 8
	}

	return SignalValueLookup[d.InputSignal]
}
