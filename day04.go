package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()

	fmt.Println("done")
}

// cribbed from ScanLines implementation
func ScanEmptyLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

type BingoSquare struct {
	Value    int
	IsMarked bool
}

func (sq *BingoSquare) String() string {
	if sq.IsMarked {
		return fmt.Sprintf("%2d'", sq.Value)
	} else {
		return fmt.Sprintf("%2d ", sq.Value)
	}
}

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

func NewBingoBoard(rawData string) *BingoBoard {
	rows := strings.Split(rawData, "\n")

	gridYX := make([][]*BingoSquare, 5)
	for y, row := range rows {
		gridYX[y] = make([]*BingoSquare, 5)

		for x, stringValue := range strings.Fields(row) {
			intValue, err := strconv.Atoi(stringValue)

			if err != nil {
				log.Fatal(err)
			}

			gridYX[y][x] = &BingoSquare{
				intValue,
				false,
			}
		}
	}

	return &BingoBoard{
		gridYX,
	}
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

func FindFirstWinner(game []int, boards []*BingoBoard) int {
	for _, n := range game {
		for _, b := range boards {
			b.Play(n)
			if b.IsWinner() {
				return b.SumUnmarked() * n
			}
		}
	}
	return 0
}

func FindLastWinner(game []int, boards []*BingoBoard) int {
	winnerCount := 0
	for _, n := range game {
		for _, b := range boards {
			if b.IsWinner() {
				continue
			}
			b.Play(n)
			if b.IsWinner() {
				winnerCount++
				if winnerCount == len(boards) {
					return b.SumUnmarked() * n
				}
			}
		}
	}
	return 0
}

func part1() {
	game, boards := load()

	winningScore := FindFirstWinner(game, boards)

	fmt.Println("part 1: ", winningScore)
}

func part2() {
	game, boards := load()

	winningScore := FindLastWinner(game, boards)

	fmt.Println("part 2: ", winningScore)
}

func load() (game []int, boards []*BingoBoard) {
	//file, err := os.Open("./test-data/day04.txt")
	file, err := os.Open("./input-data/day04.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)
	s := bufio.NewScanner(r)
	s.Split(ScanEmptyLines)

	// load data
	isFirstLine := true
	for s.Scan() {
		if isFirstLine {
			// load data: create game
			for _, x := range strings.Split(s.Text(), ",") {
				num, err := strconv.Atoi(x)
				if err != nil {
					log.Fatal(err)
				}
				game = append(game, num)
			}
			isFirstLine = false
			continue
		}

		// load data: create boards
		boards = append(boards, NewBingoBoard(strings.TrimSpace(s.Text())))
	}

	return game, boards
}
