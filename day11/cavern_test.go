package day11

import (
	"strings"
	"testing"

	"github.com/klandergren/advent-of-code-2021/util"
)

var testInput_Flashes_Initial = `11111
19991
19191
19991
11111
`

var testInput_Flashes_AfterStep1 = `34543
40004
50005
40004
34543
`

var testInput_Flashes_AfterStep2 = `45654
51115
61116
51115
45654
`

var testInput_Counts = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`

func Test_FlashLogic(t *testing.T) {
	r := strings.NewReader(testInput_Flashes_Initial)
	lines, err := util.LoadLines(r)
	if err != nil {
		t.Error(err)
	}
	c, err := NewCavern(lines)
	if err != nil {
		t.Error(err)
	}

	if c.String() != testInput_Flashes_Initial {
		t.Errorf("initial state should look like\n: %s \n\n but got:\n %s", testInput_Flashes_Initial, c.String())
	}

	c.Step()

	if c.String() != testInput_Flashes_AfterStep1 {
		t.Errorf("after step 1 state should look like\n: %s \n\n but got:\n %s", testInput_Flashes_AfterStep1, c.String())
	}

	c.Step()

	if c.String() != testInput_Flashes_AfterStep2 {
		t.Errorf("after step 1 state should look like\n: %s \n\n but got:\n %s", testInput_Flashes_AfterStep2, c.String())
	}

}

func Test_Counts(t *testing.T) {
	r := strings.NewReader(testInput_Counts)

	lines, err := util.LoadLines(r)
	if err != nil {
		t.Error(err)
	}

	c, err := NewCavern(lines)
	if err != nil {
		t.Error(err)
	}

	flashCount := c.Advance(10)
	if c.StepCount != 10 {
		t.Errorf("expected c.StepCount return of 10 but got: %d", c.StepCount)
	}
	if flashCount != 204 {
		t.Errorf("expected flashCount return of 204 but got: %d", flashCount)
	}
	if c.TotalFlashCount != 204 {
		t.Errorf("expected c.TotalFlashCount return of 204 but got: %d", c.TotalFlashCount)
	}

	c.Advance(90)
	if c.StepCount != 100 {
		t.Errorf("expected c.StepCount return of 100 but got: %d", c.StepCount)
	}

	if c.TotalFlashCount != 1656 {
		t.Errorf("expected c.TotalFlashCount return of 1656 but got: %d", c.TotalFlashCount)
	}
}
