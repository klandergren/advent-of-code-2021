package day04

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
)

func PartOne(reader io.Reader) (int, error) {
	game, boards, err := load(reader)
	if err != nil {
		return -1, err
	}

	winningScore := FindFirstWinner(game, boards)

	return winningScore, nil
}

func PartTwo(reader io.Reader) (int, error) {
	game, boards, err := load(reader)
	if err != nil {
		return -1, err
	}

	winningScore := FindLastWinner(game, boards)

	return winningScore, nil
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

func load(reader io.Reader) (game []int, boards []*BingoBoard, err error) {
	s := bufio.NewScanner(reader)
	s.Split(ScanEmptyLines)

	// load data
	isFirstLine := true
	for s.Scan() {
		if isFirstLine {
			// load data: create game
			for _, x := range strings.Split(s.Text(), ",") {
				num, err := strconv.Atoi(x)
				if err != nil {
					return nil, nil, err
				}
				game = append(game, num)
			}
			isFirstLine = false
			continue
		}

		// load data: create boards
		bb, err := NewBingoBoard(strings.TrimSpace(s.Text()))

		if err != nil {
			return nil, nil, err
		}
		boards = append(boards, bb)
	}

	return game, boards, nil
}
