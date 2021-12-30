package day13

import (
	"strings"
	"testing"

	"github.com/klandergren/advent-of-code-2021/util"
)

var tpData = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0
`

func Test_LoadTransparentPaper(t *testing.T) {
	r := strings.NewReader(tpData)
	lines, err := util.LoadLines(r)
	if err != nil {
		t.Error(err)
	}

	tp, err := LoadTransparentPaper(lines)
	if err != nil {
		t.Error(err)
	}

	if len(*tp) != 15 {
		t.Errorf("transparent paper should have 15 rows, got: %d\n", len(*tp))
	}

	if len((*tp)[0]) != 11 {
		t.Errorf("transparent paper should have 11 columns, got: %d\n", len((*tp)[0]))
	}

	expectedPattern := `...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
...........
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........
`

	if tp.String() != expectedPattern {
		t.Errorf("transparent paper should look like:\n%s\n but got:\n%s\n", expectedPattern, tp)
	}

}

func Test_FoldTransparentPaper(t *testing.T) {
	r := strings.NewReader(tpData)
	lines, err := util.LoadLines(r)
	if err != nil {
		t.Error(err)
	}

	tp, err := LoadTransparentPaper(lines)
	if err != nil {
		t.Error(err)
	}

	tp.FoldUpAt(7)

	expectedFoldY7 := `#.##..#..#.
#...#......
......#...#
#...#......
.#.#..#.###
...........
...........
`
	if tp.String() != expectedFoldY7 {
		t.Errorf("after folding at y=7, transparent paper should look like:\n%s\n but got:\n%s\n", expectedFoldY7, tp)
	}

	dotsVisible := tp.DotsVisible()
	if dotsVisible != 17 {
		t.Errorf("after folding at y=7, dots visible should be %d but got %d", 17, dotsVisible)
	}

	tp.FoldLeftAt(5)

	expectedFoldX5 := `#####
#...#
#...#
#...#
#####
.....
.....
`
	if tp.String() != expectedFoldX5 {
		t.Errorf("after folding at x=5, transparent paper should look like:\n%s\n but got:\n%s\n", expectedFoldX5, tp)
	}

	dotsVisible = tp.DotsVisible()
	if dotsVisible != 16 {
		t.Errorf("after folding at y=7, dots visible should be %d but got %d", 16, dotsVisible)
	}

}
