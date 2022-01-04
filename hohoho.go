package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/klandergren/advent-of-code-2021/day01"
	"github.com/klandergren/advent-of-code-2021/day02"
	"github.com/klandergren/advent-of-code-2021/day03"
	"github.com/klandergren/advent-of-code-2021/day04"
	"github.com/klandergren/advent-of-code-2021/day05"
	"github.com/klandergren/advent-of-code-2021/day06"
	"github.com/klandergren/advent-of-code-2021/day07"
	"github.com/klandergren/advent-of-code-2021/day08"
	"github.com/klandergren/advent-of-code-2021/day09"
	"github.com/klandergren/advent-of-code-2021/day10"
	"github.com/klandergren/advent-of-code-2021/day11"
	"github.com/klandergren/advent-of-code-2021/day12"
	"github.com/klandergren/advent-of-code-2021/day13"
	"github.com/klandergren/advent-of-code-2021/day14"
	"github.com/klandergren/advent-of-code-2021/day15"
	"github.com/klandergren/advent-of-code-2021/day16"
	"github.com/klandergren/advent-of-code-2021/day17"
	"github.com/klandergren/advent-of-code-2021/day18"
	"github.com/klandergren/advent-of-code-2021/day19"
	"github.com/klandergren/advent-of-code-2021/day20"
	"github.com/klandergren/advent-of-code-2021/day21"
	"github.com/klandergren/advent-of-code-2021/day22"
	"github.com/klandergren/advent-of-code-2021/day23"
	"github.com/klandergren/advent-of-code-2021/day24"
	"github.com/klandergren/advent-of-code-2021/day25"
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
		log.Fatalln("error:", err)
	}

	if err := validatePart(part); err != nil {
		log.Fatalln("error:", err)
	}

	if err := validateInputFilePath(inputFilePath); err != nil {
		log.Fatalln("error:", err)
	}

	if part == -1 {
		for _, p := range []int{1, 2} {
			err := loadFileAndRun(inputFilePath, day, p)
			if err != nil {
				log.Fatalln("error:", err)
			}
		}
	} else {
		err := loadFileAndRun(inputFilePath, day, part)
		if err != nil {
			log.Fatalln("error:", err)
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

	fmt.Printf("Answer day %d part %d (from input \"%s\"):\n%s\n", day, part, inputFilePath, result)
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

func run(day int, part int, reader io.Reader) (result string, err error) {
	switch day {
	case 1:
		switch part {
		case 1:
			r, err := day01.PartOne(reader)
			return strconv.Itoa(r), err
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
	case 9:
		switch part {
		case 1:
			return day09.PartOne(reader)
		case 2:
			return day09.PartTwo(reader)
		}
	case 10:
		switch part {
		case 1:
			return day10.PartOne(reader)
		case 2:
			return day10.PartTwo(reader)
		}
	case 11:
		switch part {
		case 1:
			return day11.PartOne(reader)
		case 2:
			return day11.PartTwo(reader)
		}
	case 12:
		switch part {
		case 1:
			return day12.PartOne(reader)
		case 2:
			return day12.PartTwo(reader)
		}
	case 13:
		switch part {
		case 1:
			return day13.PartOne(reader)
		case 2:
			return day13.PartTwo(reader)
		}
	case 14:
		switch part {
		case 1:
			return day14.PartOne(reader)
		case 2:
			return day14.PartTwo(reader)
		}
	case 15:
		switch part {
		case 1:
			return day15.PartOne(reader)
		case 2:
			return day15.PartTwo(reader)
		}
	case 16:
		switch part {
		case 1:
			return day16.PartOne(reader)
		case 2:
			return day16.PartTwo(reader)
		}
	case 17:
		switch part {
		case 1:
			return day17.PartOne(reader)
		case 2:
			return day17.PartTwo(reader)
		}
	case 18:
		switch part {
		case 1:
			return day18.PartOne(reader)
		case 2:
			return day18.PartTwo(reader)
		}
	case 19:
		switch part {
		case 1:
			return day19.PartOne(reader)
		case 2:
			return day19.PartTwo(reader)
		}
	case 20:
		switch part {
		case 1:
			return day20.PartOne(reader)
		case 2:
			return day20.PartTwo(reader)
		}
	case 21:
		switch part {
		case 1:
			return day21.PartOne(reader)
		case 2:
			return day21.PartTwo(reader)
		}
	case 22:
		switch part {
		case 1:
			return day22.PartOne(reader)
		case 2:
			return day22.PartTwo(reader)
		}
	case 23:
		switch part {
		case 1:
			return day23.PartOne(reader)
		case 2:
			return day23.PartTwo(reader)
		}
	case 24:
		switch part {
		case 1:
			return day24.PartOne(reader)
		case 2:
			return day24.PartTwo(reader)
		}
	case 25:
		switch part {
		case 1:
			return day25.PartOne(reader)
		case 2:
			return day25.PartTwo(reader)
		}
	}

	return "", errors.New("problem running")
}
