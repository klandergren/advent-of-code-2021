package day08

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func PartOne(reader *bufio.Reader) (int, error) {
	return -1, errors.New("not implemented yet")
}
func PartTwo(reader *bufio.Reader) (int, error) {
	return -1, errors.New("not implemented yet")
}

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

var BaseTransform = map[string]string{
	"a": "a",
	"b": "b",
	"c": "c",
	"d": "d",
	"e": "e",
	"f": "f",
	"g": "g",
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
	newWireSignal := make([]int, len(BaseTransform))

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

type Display struct {
	InputSignals []string
}

func (d *Display) OutputValue() int {
	// create digits
	digits := make([]*Digit, len(d.InputSignals))

	for i, s := range d.InputSignals {
		d := NewDigitFromInput(s)
		digits[i] = d
	}

	// create string of values
	var sb strings.Builder
	for _, d := range digits {
		sb.WriteString(fmt.Sprintf("%d", d.Value()))
	}

	v, err := strconv.Atoi(sb.String())

	if err != nil {
		log.Fatal(err)

	}

	return v
}

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

func GenerateTransform(r *Reading) map[string]string {
	transform := make(map[string]string, len(BaseTransform))

	d1 := NewDigitFromInput(r.InputSignal1())
	d4 := NewDigitFromInput(r.InputSignal4())
	d7 := NewDigitFromInput(r.InputSignal7())
	d8 := NewDigitFromInput(r.InputSignal8())

	// find segment A
	dSegmentA := d7.Subtract(d1)
	transform[dSegmentA.InputSignal] = "a"

	// find segment B and segment D
	dMaybeBD := d4.Subtract(d1)

	// looking at 0, 6, or 9, we can look for segments that have
	// a unique amount within the set
	sixes := r.SignalsOfLength(6)

	dSixChars := make([]*Digit, len(sixes))
	for i, s := range sixes {
		dSixChars[i] = NewDigitFromInput(s)
	}

	for _, maybe := range dMaybeBD.RawInputSignals() {
		count := 0
		for _, d := range dSixChars {
			if strings.Contains(d.InputSignal, maybe) {
				count++
			}
		}
		if count == 2 {
			transform[maybe] = "d"
		}
		if count == 3 {
			transform[maybe] = "b"
		}
	}

	// find segment C and segment F by the same manner
	dMaybeCF := d7.Intersect(d1)

	for _, maybe := range dMaybeCF.RawInputSignals() {
		count := 0
		for _, d := range dSixChars {
			if strings.Contains(d.InputSignal, maybe) {
				count++
			}
		}
		if count == 2 {
			transform[maybe] = "c"
		}
		if count == 3 {
			transform[maybe] = "f"
		}
	}

	// and finally segment E and segment G by creating a synthetic
	// digit of the signals corresponding to the discovered segments
	var sb strings.Builder
	for k, _ := range transform {
		sb.WriteString(k)
	}
	knownSegments := sb.String()

	dSynthetic := NewDigitFromInput(knownSegments)
	dMaybeEG := d8.Subtract(dSynthetic)

	for _, maybe := range dMaybeEG.RawInputSignals() {
		count := 0
		for _, d := range dSixChars {
			if strings.Contains(d.InputSignal, maybe) {
				count++
			}
		}
		if count == 2 {
			transform[maybe] = "e"
		}
		if count == 3 {
			transform[maybe] = "g"
		}
	}
	return transform
}

func Transform(outputSignals []string, transform map[string]string) (transformedSignals []string) {
	for _, os := range outputSignals {
		rawSignals := strings.Split(os, "")

		newSignal := make([]string, len(rawSignals))
		for _, rawSignal := range rawSignals {
			newSignal = append(newSignal, transform[rawSignal])
		}

		transformedSignals = append(transformedSignals, strings.Join(newSignal, ""))
	}

	return transformedSignals
}

func (r *Reading) String() string {
	return fmt.Sprintf("%s | %s\n", r.UniqueSignals, r.OutputSignals)
}

func partOne() {
	readings := load()

	count1478 := 0

	for _, r := range readings {
		for _, s := range r.OutputSignals {
			d := NewDigitFromInput(s)

			switch v := d.Value(); v {
			case 1:
				count1478++
			case 4:
				count1478++
			case 7:
				count1478++
			case 8:
				count1478++
			}
		}
	}

	fmt.Println("part one: ", count1478)
}

func partTwo() {
	readings := load()

	sum := 0

	for _, r := range readings {
		t := GenerateTransform(r)

		d := &Display{InputSignals: Transform(r.OutputSignals, t)}

		sum += d.OutputValue()
	}

	fmt.Println("part two: ", sum)
}

func load() (readings []*Reading) {
	//file, err := os.Open("./test-data/day08a.txt")
	//file, err := os.Open("./test-data/day08b.txt")
	file, err := os.Open("./input-data/day08.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		data := strings.Split(string(line), "|")

		uniqueSignals := strings.Fields(data[0])
		outputSignals := strings.Fields(data[1])

		readings = append(readings, NewReading(uniqueSignals, outputSignals))
	}

	return readings
}
