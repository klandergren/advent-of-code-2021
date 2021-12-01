package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	parta()
	partb()

	fmt.Println("done")
}

func parta() {
	//	file, err := os.Open("./test-data/day01.txt")
	file, err := os.Open("./input-data/day01.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	prev_depth := math.MaxInt64

	c := 0

	for {
		l, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		depth, _ := strconv.Atoi(string(l))

		if prev_depth == math.MaxInt64 {
			fmt.Println(depth, "(N/A - no previous measurement)")
		} else if prev_depth < depth {
			fmt.Println(depth, "(increased)")
			c++
		} else {
			fmt.Println(depth, "(decreased)")
		}
		prev_depth = depth

	}

	fmt.Println("part a: ", c)
}

func partb() {
	//	file, err := os.Open("./test-data/day01.txt")
	file, err := os.Open("./input-data/day01.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	prev_window := math.MaxInt64
	depths := make([]int, 0)

	c := 0

	for {
		l, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		depth, _ := strconv.Atoi(string(l))

		depths = append(depths, depth)

		if len(depths) < 3 {
			// not enough measurements yet
			continue
		} else if len(depths) == 3 {
			// we have first measurement
		} else {
			// pop
			_, depths = depths[0], depths[1:]
		}

		window := 0
		for _, d := range depths {
			window += d
		}

		if prev_window == math.MaxInt64 {
			fmt.Println(window, "(N/A - no previous measurement)")
		} else if prev_window < window {
			fmt.Println(window, "(increased)")
			c++
		} else if prev_window == window {
			fmt.Println(window, "(no change)")
		} else {
			fmt.Println(window, "(decreased)")
		}
		prev_window = window

	}

	fmt.Println("part b: ", c)
}
