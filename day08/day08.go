package day08

import (
	"bufio"
	"io"
	"strings"
)

func PartOne(reader *bufio.Reader) (int, error) {
	readings, err := load(reader)
	if err != nil {
		return -1, err
	}

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

	return count1478, nil
}

func PartTwo(reader *bufio.Reader) (int, error) {
	readings, err := load(reader)
	if err != nil {
		return -1, err
	}

	sum := 0

	for _, r := range readings {
		t := GenerateTransform(r)

		d := &Display{InputSignals: Transform(r.OutputSignals, t)}

		v, err := d.OutputValue()
		if err != nil {
			return -1, err
		}

		sum += v
	}

	return sum, nil
}

func GenerateTransform(r *Reading) map[string]string {
	transform := make(map[string]string, 7)

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

func load(reader *bufio.Reader) (readings []*Reading, err error) {
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		data := strings.Split(string(line), "|")

		uniqueSignals := strings.Fields(data[0])
		outputSignals := strings.Fields(data[1])

		readings = append(readings, NewReading(uniqueSignals, outputSignals))
	}

	return readings, nil
}
