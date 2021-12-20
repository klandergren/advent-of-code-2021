package day10

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/klandergren/advent-of-code-2021/util"
)

func PartOne(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	illegal := make([]string, 0)
	for _, line := range lines {

		chars := strings.Split(line, "")

		_, err := Parse(chars)
		if err == nil {
			// incomplete line
		} else {
			//			fmt.Println("illegal character:", err.Error(), "found in line:", line)
			illegal = append(illegal, err.Error())
		}
	}

	score := 0
	for _, c := range illegal {
		switch c {
		case ")":
			score += 3
		case "]":
			score += 57
		case "}":
			score += 1197
		case ">":
			score += 25137
		default:
			return -1, fmt.Errorf("day10: unknown character: %s", c)
		}
	}

	return score, nil
}

func PartTwo(reader io.Reader) (int, error) {
	lines, err := util.LoadLines(reader)
	if err != nil {
		return -1, err
	}

	scores := make([]int, 0)
	for _, line := range lines {

		chars := strings.Split(line, "")

		chunks, err := Parse(chars)
		if err == nil {
			completions := make([]string, 0)
			for _, c := range *chunks {
				var closingChar string
				switch c.OpenChar {
				case "(":
					closingChar = ")"
				case "[":
					closingChar = "]"
				case "{":
					closingChar = "}"
				case "<":
					closingChar = ">"
				}
				completions = append(completions, closingChar)
			}

			// reverse the list
			lastIndex := len(completions) - 1
			for i := 0; i < len(completions)/2; i++ {
				completions[i], completions[lastIndex-i] = completions[lastIndex-i], completions[i]
			}

			score := 0
			for _, c := range completions {
				switch c {
				case ")":
					score = (score * 5) + 1
				case "]":
					score = (score * 5) + 2
				case "}":
					score = (score * 5) + 3
				case ">":
					score = (score * 5) + 4
				default:
					return -1, fmt.Errorf("day10: unknown character: %s", c)
				}
			}
			scores = append(scores, score)

		} else {
			//fmt.Println("illegal character:", err.Error(), "found in line:", line)
		}
	}

	// sort
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	// take the middle score
	score := scores[len(scores)/2]

	return score, nil
}

func Parse(chars []string) (*[]*Chunk, error) {
	if len(chars) == 0 {
		// TODO better error
		return nil, errors.New("!") // cannot parse empty
	}

	chunks := make([]*Chunk, 0)

	pos := -1
	for pos < len(chars) {
		pos = pos + 1

		if len(chars) <= pos {
			// TODO better chunk return and error
			return &chunks, nil
		}

		char := chars[pos]

		if IsOpen(char) {
			// start new chunk
			c := &Chunk{char, false, make([]Chunk, 0)}
			chunks = append(chunks, c)
		} else {
			// we should be closing whatever the last value of chunks is

			// check scenario where we have been given a line where the position of a
			// new chunk is actually a closing character
			if len(chunks) == 0 {
				return nil, errors.New(char)
			}

			// grab the last chunk off the slice
			lastPosition := len(chunks) - 1
			lastChunk := chunks[lastPosition]
			chunks = chunks[:lastPosition]
			lastOpenChar := lastChunk.OpenChar

			// check for errors
			switch char {
			case ">":
				if lastOpenChar != "<" {
					return nil, errors.New(char)
				}
			case "}":
				if lastOpenChar != "{" {
					return nil, errors.New(char)
				}
			case "]":
				if lastOpenChar != "[" {
					return nil, errors.New(char)
				}
			case ")":
				if lastOpenChar != "(" {
					return nil, errors.New(char)
				}
			default:
				return nil, errors.New("?")
			}

			// no errors; we have completed the last chunk
			lastChunk.Complete = true

			// is there a parent?
			if 0 <= len(chunks)-1 {
				// add it to its parent
				parentChunk := chunks[len(chunks)-1]
				parentChunk.Chunks = append(parentChunk.Chunks, *lastChunk)
			} else {
				// no parent
			}

		}
	}

	return nil, errors.New("!") // how'd we get here?
}

func IsOpen(c string) bool {
	return c == "(" || c == "[" || c == "{" || c == "<"
}
