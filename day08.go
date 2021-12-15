package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	part1()
	part2()

	fmt.Println("done")
}

type Reading struct {
	SignalPatterns []string
	OutputValues   []string
}

func (r *Reading) NumUnique() int {
	count := 0

	for _, ov := range r.OutputValues {

		l := len(ov)
		if (l == 2) || (l == 3) || (l == 4) || (l == 7) {
			count++
		}
	}

	return count
}

func (r *Reading) String() string {
	return fmt.Sprintf("%s | %s\n", r.SignalPatterns, r.OutputValues)
}

func part1() {
	readings := load()

	totalUnique := 0

	for _, r := range readings {
		totalUnique += r.NumUnique()
	}

	fmt.Println("part 1: ", totalUnique)
}

func part2() {
	fmt.Println("part 2: ", 0)
}

func load() (readings []*Reading) {
	//file, err := os.Open("./test-data/day08.txt")
	file, err := os.Open("./input-data/day08.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		data := strings.Split(string(line), "|")
		signalPatterns := strings.Fields(data[0])
		outputValues := strings.Fields(data[1])

		readings = append(readings, &Reading{
			SignalPatterns: signalPatterns,
			OutputValues:   outputValues,
		})

	}

	return readings
}
