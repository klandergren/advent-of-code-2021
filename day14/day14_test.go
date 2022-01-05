package day14

import (
	"strings"
	"testing"
)

var inputData = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

func Test_PartOne(t *testing.T) {
	r := strings.NewReader(inputData)

	ans, err := PartOne(r)
	if err != nil {
		t.Fatal(err)
	}

	if ans != 1588 {
		t.Errorf("expected 1588 but got: %d", ans)
	}

}
