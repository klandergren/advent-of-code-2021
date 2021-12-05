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

type BingoBoard struct {
	gridYX [][]*BingoSquare
}

func NewBingoBoard(rawData string) *BingoBoard {
	fmt.Println(rawData)
	rows := strings.Split(rawData, "\n")

	gridYX := make([][]*BingoSquare, 5)
	for y, row := range rows {
		fmt.Println("y:", y)
		fmt.Println("row:", row)
		gridYX[y] = make([]*BingoSquare, 5)

		for x, stringValue := range strings.Fields(row) {
			fmt.Println("x:", x)
			fmt.Println("stringValue:", stringValue)
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

	fmt.Println(gridYX)
	return &BingoBoard{
		gridYX,
	}
}

func (b *BingoBoard) Play(game []int) (isWinner bool, score int) {
	fmt.Printf("playing game: %+v\n", game)

	// mark drawn numbers
	for _, n := range game {
		for _, row := range b.gridYX {
			for _, square := range row {
				fmt.Printf("square before: %+v\n", square)
				square.IsMarked = (square.Value == n)
				fmt.Printf("square after: %+v\n", square)
			}
		}
	}
	fmt.Printf("board: %+v\n", b)

	// check rows
	for _, row := range b.gridYX {
		for _, square := range row {
		}
	}

	// check cols

	return false, 0
}

func part1() {
	file, err := os.Open("./test-data/day04.txt")
	//file, err := os.Open("./input-data/day03.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)
	s := bufio.NewScanner(r)
	s.Split(ScanEmptyLines)

	var game []int
	var boards []*BingoBoard

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
	fmt.Printf("game: %+v\n", game)
	fmt.Printf("boards: %+v\n", boards)

	// runner: run game
	winningScore := 0
	for i := 1; i < len(game); i++ {
		for _, b := range boards {
			if isWinner, score := b.Play(game[:i]); isWinner {
				winningScore = score
				fmt.Println("we win. score:", score)
			}
		}
	}

	fmt.Println("part 1: ", winningScore)
}

func part2() {
}
