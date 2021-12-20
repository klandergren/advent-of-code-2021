package day10

import (
	"errors"
	"fmt"
	"io"
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

		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("parsing line:", line)
		_, err := Parse(chars)
		if err == nil {
			fmt.Println("incomplete line:", line)
		} else {
			fmt.Println("illegal character:", err.Error(), "found in line:", line)
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

func Parse(chars []string) (*[]Chunk, error) {
	fmt.Println("chars:", chars)
	if len(chars) == 0 {
		// TODO better error
		return nil, errors.New("!") // cannot parse empty
	}

	chunks := make([]*Chunk, 0)

	pos := -1
	for pos < len(chars) {
		pos = pos + 1
		fmt.Println("start pos:", pos)

		if len(chars) <= pos {
			// TODO better chunk return and error
			return nil, nil
		}

		char := chars[pos]
		fmt.Println("start char:", char)
		fmt.Println("start chunks:", chunks)

		if IsOpen(char) {
			fmt.Println("creating new")
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
			fmt.Println("closing chunk:", lastChunk)
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
				fmt.Println("parentChunk:", parentChunk)
				parentChunk.Chunks = append(parentChunk.Chunks, *lastChunk)
			} else {
				fmt.Println("no parent")
			}

		}

		fmt.Println("end pos:", pos)
		fmt.Println("end char:", char)
		fmt.Println("end chunks:", chunks)
	}

	return nil, errors.New("!") // how'd we get here?
}

func IsOpen(c string) bool {
	return c == "(" || c == "[" || c == "{" || c == "<"
}

// func IsValid(p, n string) bool {
// 	switch p {
// 	case "START":
// 		return IsOpen(n)
// 	case "(":
// 		return IsOpen(n) || n == ")"
// 	case "[":
// 		return IsOpen(n) || n == "]"
// 	case "{":
// 		return IsOpen(n) || n == "}"
// 	case "<":
// 		return IsOpen(n) || n == ">"
// 	case ">":
// 	case "}":
// 	case "]":
// 	case ")":
// 	}
// }

func PartTwo(reader io.Reader) (int, error) {
	return -1, errors.New("not implemented")
}
