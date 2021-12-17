package util

import (
	"bufio"
	"io"
)

func LoadLines(reader io.Reader) (lines []string, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
