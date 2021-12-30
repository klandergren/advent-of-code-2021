package day13

import (
	"strconv"
	"strings"
)

type TransparentPaper [][]rune

func LoadTransparentPaper(lines []string) (*TransparentPaper, error) {

	// load coordinates
	coordinates := make([]Coordinate, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		coordinates = append(coordinates, Coordinate{x, y})
	}

	// figure out bounds
	var maxX int
	var maxY int
	for _, coord := range coordinates {
		if maxX < coord.X {
			maxX = coord.X
		}
		if maxY < coord.Y {
			maxY = coord.Y
		}
	}

	// make blank grid
	tp := make(TransparentPaper, maxY+1)
	for y, _ := range tp {
		row := make([]rune, maxX+1)

		for x, _ := range row {
			row[x] = '.'
		}

		tp[y] = row
	}

	// populate grid according to coordinates
	for _, c := range coordinates {
		tp[c.Y][c.X] = '#'
	}

	return &tp, nil
}

func (tp *TransparentPaper) DotsVisible() (count int) {
	for _, row := range *tp {
		for _, r := range row {
			if r == '#' {
				count += 1
			}
		}
	}
	return count
}

func (tp *TransparentPaper) Apply(instruction Instruction) {
	if instruction.Direction == "left" {
		tp.FoldLeftAt(instruction.Value)
	} else {
		tp.FoldUpAt(instruction.Value)
	}
}

func (tp *TransparentPaper) MaxX() int {
	return len((*tp)[0])
}

func (tp *TransparentPaper) MaxY() int {
	return len(*tp)
}

func (tp *TransparentPaper) FoldUpAt(y int) {
	// move everything after y into its own data
	bottom := (*tp)[y+1 : len(*tp)]

	// remove after the y line
	*tp = (*tp)[0:y]

	// reverse the bottom and apply to tp
	for i, _ := range bottom {
		k := len(bottom) - 1 - i

		row := bottom[k]

		topY := len(*tp) - k - 1

		for x, r := range row {
			if r == '#' {
				(*tp)[topY][x] = r
			}
		}
	}
}

func (tp *TransparentPaper) FoldLeftAt(x int) {
	// remove everything after x into its own data
	right := make([][]rune, 0)
	for y, row := range *tp {
		(*tp)[y] = row[0:x]

		right = append(right, row[x+1:len(row)])
	}

	// reverse each row of right and apply it to tp
	for y, row := range right {
		for i, _ := range row {
			x := len(row) - 1 - i
			r := right[y][x]

			xTp := len((*tp)[y]) - x - 1
			if r == '#' {
				(*tp)[y][xTp] = r
			}

		}
	}
}

func (tp *TransparentPaper) String() string {
	var sb strings.Builder
	for _, row := range *tp {
		for _, r := range row {
			sb.WriteString(string(r))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
