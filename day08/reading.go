package day08

import (
	"fmt"
	"sort"
	"strings"
)

type Reading struct {
	UniqueSignals []string
	OutputSignals []string
}

// makes it easier to sort the signals for debugging
func NewReading(u, o []string) *Reading {
	uniqueSignals := make([]string, len(u))

	for i, s := range u {
		n := strings.Split(s, "")
		sort.Strings(n)
		uniqueSignals[i] = strings.Join(n, "")
	}

	outputSignals := make([]string, len(o))

	for i, s := range o {
		n := strings.Split(s, "")
		sort.Strings(n)
		outputSignals[i] = strings.Join(n, "")
	}

	return &Reading{
		UniqueSignals: uniqueSignals,
		OutputSignals: outputSignals,
	}
}

func (r *Reading) SignalsOfLength(l int) (patterns []string) {
	for _, us := range r.UniqueSignals {
		if len(us) == l {
			patterns = append(patterns, us)
		}
	}
	return patterns
}

func (r *Reading) InputSignal1() string {
	return r.SignalsOfLength(2)[0]
}

func (r *Reading) InputSignal4() string {
	return r.SignalsOfLength(4)[0]
}

func (r *Reading) InputSignal7() string {
	return r.SignalsOfLength(3)[0]
}

func (r *Reading) InputSignal8() string {
	return r.SignalsOfLength(7)[0]
}

func (r *Reading) String() string {
	return fmt.Sprintf("%s | %s\n", r.UniqueSignals, r.OutputSignals)
}
