package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/klandergren/advent-of-code-2021/day01"
	"github.com/klandergren/advent-of-code-2021/day02"
	"github.com/klandergren/advent-of-code-2021/day03"
	"github.com/klandergren/advent-of-code-2021/day04"
	"github.com/klandergren/advent-of-code-2021/day05"
	"github.com/klandergren/advent-of-code-2021/day06"
	"github.com/klandergren/advent-of-code-2021/day07"
	"github.com/klandergren/advent-of-code-2021/day08"
)

var dayFlag = flag.Int("day", 1, "[1-25] the day number to run, without leading 0")
var partFlag = flag.Int("part", -1, "[1-2] the part number to run")
var inputFilePathFlag = flag.String("inputFilePath", "", "the input file path")

func main() {
	flag.Parse()

	day := *dayFlag
	part := *partFlag
	inputFilePath := *inputFilePathFlag

	if err := validateDay(day); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	if err := validatePart(part); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	if err := validateInputFilePath(inputFilePath); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	if part == -1 {
		for _, p := range []int{1, 2} {
			err := loadFileAndRun(inputFilePath, day, p)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
		}
	} else {
		err := loadFileAndRun(inputFilePath, day, part)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}

}

func loadFileAndRun(inputFilePath string, day int, part int) error {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	result, err := run(day, part, reader)
	if err != nil {
		return err
	}

	fmt.Printf("Answer day %d part %d (from input \"%s\"):\n%d\n", day, part, inputFilePath, result)
	return nil
}

func validateDay(day int) error {
	if day < 1 || 25 < day {
		return fmt.Errorf("%d is outside the accepted day range of 1 to 25 (inclusive)", day)
	}
	return nil
}

func validatePart(part int) error {
	// special case for running both
	if part == -1 {
		return nil
	}

	if part < 1 || 2 < part {
		return fmt.Errorf("%d is outside the accepted part range of 1 to 2 (inclusive)", part)
	}
	return nil
}

func validateInputFilePath(inputFilePath string) error {
	if len(inputFilePath) == 0 {
		return errors.New("input file path must be specified")
	}
	return nil
}

func run(day int, part int, reader *bufio.Reader) (result int, err error) {
	switch day {
	case 1:
		switch part {
		case 1:
			return day01.PartOne(reader)
		case 2:
			return day01.PartTwo(reader)
		}
	case 2:
		switch part {
		case 1:
			return day02.PartOne(reader)
		case 2:
			return day02.PartTwo(reader)
		}
	case 3:
		switch part {
		case 1:
			return day03.PartOne(reader)
		case 2:
			return day03.PartTwo(reader)
		}
	case 4:
		switch part {
		case 1:
			return day04.PartOne(reader)
		case 2:
			return day04.PartTwo(reader)
		}
	case 5:
		switch part {
		case 1:
			return day05.PartOne(reader)
		case 2:
			return day05.PartTwo(reader)
		}
	case 6:
		switch part {
		case 1:
			return day06.PartOne(reader)
		case 2:
			return day06.PartTwo(reader)
		}
	case 7:
		switch part {
		case 1:
			return day07.PartOne(reader)
		case 2:
			return day07.PartTwo(reader)
		}
	case 8:
		switch part {
		case 1:
			return day08.PartOne(reader)
		case 2:
			return day08.PartTwo(reader)
		}
	}

	return -1, errors.New("problem running")
}
