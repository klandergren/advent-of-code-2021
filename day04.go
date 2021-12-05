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

type BingoBoard struct {
	grid [][]int
}

func NewBingoBoard(rawdata string) BingoBoard {
	return BingoBoard{
		nil,
	}
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
	var boards []BingoBoard

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
		boards = append(boards, NewBingoBoard(s.Text()))
	}
	fmt.Println(game)
	fmt.Println(boards)

	// runner: run game
	fmt.Println("part 1: ", 0)
}

func part2() {
}
