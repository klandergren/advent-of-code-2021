package day12

import (
	"fmt"
	"strings"
	"testing"

	"github.com/klandergren/advent-of-code-2021/util"
)

var testInputA = `start-A
start-b
A-c
A-b
b-d
A-end
b-end
`

func Test_PartOne(t *testing.T) {
	t.SkipNow()
	r := strings.NewReader(testInputA)
	lines, err := util.LoadLines(r)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(lines)
}

func Test_PartTwo(t *testing.T) {
	t.SkipNow()
	r := strings.NewReader(testInputA)
	lines, err := util.LoadLines(r)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(lines)

}
