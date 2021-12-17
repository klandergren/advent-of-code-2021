package day04

import (
	"strconv"
	"strings"
)

type BingoBoard struct {
	gridYX [][]*BingoSquare
}

func (b *BingoBoard) String() string {
	var sb strings.Builder

	sb.WriteString("\n")
	for _, row := range b.gridYX {
		for _, square := range row {
			sb.WriteString(square.String())
			sb.WriteString(" ")
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func NewBingoBoard(rawData string) (*BingoBoard, error) {
	rows := strings.Split(rawData, "\n")

	gridYX := make([][]*BingoSquare, 5)
	for y, row := range rows {
		gridYX[y] = make([]*BingoSquare, 5)

		for x, stringValue := range strings.Fields(row) {
			intValue, err := strconv.Atoi(stringValue)

			if err != nil {
				return nil, err
			}

			gridYX[y][x] = &BingoSquare{
				intValue,
				false,
			}
		}
	}

	return &BingoBoard{
		gridYX,
	}, nil
}

func (b *BingoBoard) SumUnmarked() int {
	sum := 0
	for _, row := range b.gridYX {
		for _, square := range row {
			if !square.IsMarked {
				sum += square.Value
			}
		}
	}
	return sum
}

func (b *BingoBoard) Reset() {
	for _, row := range b.gridYX {
		for _, square := range row {
			square.IsMarked = false
		}
	}
}

func (b *BingoBoard) Play(num int) {
	for _, row := range b.gridYX {
		for _, square := range row {
			if square.Value == num {
				square.IsMarked = true
			}
		}
	}
}

func (b *BingoBoard) IsWinner() bool {
	// check rows
	for _, row := range b.gridYX {
		hasCompleteRow := true
		for _, square := range row {
			hasCompleteRow = hasCompleteRow && square.IsMarked
		}
		if hasCompleteRow {
			return true
		}
	}

	// check cols
	for x := 0; x < 5; x++ {
		hasCompleteCol := true
		for y := 0; y < 5; y++ {
			square := b.gridYX[y][x]
			hasCompleteCol = hasCompleteCol && square.IsMarked
		}
		if hasCompleteCol {
			return true
		}
	}
	return false
}
