package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Score int

const (
	Rock     Score = 1
	Paper    Score = 2
	Scissors Score = 3

	Loss Score = 0
	Draw Score = 3
	Win  Score = 6
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	bs, err := os.ReadFile("input.txt")
	if err != nil {
		return err
	}
	var total1, total2 Score
	for _, ln := range strings.Split(string(bs), "\n") {
		if ln == "" {
			continue
		}
		total1 += play1(ln)
		total2 += play2(ln)
	}
	fmt.Println("part 1:", total1)
	fmt.Println("part 2:", total2)
	return nil
}

func convert(c byte) Score {
	switch c {
	case 'A', 'X':
		return Rock
	case 'B', 'Y':
		return Paper
	case 'C', 'Z':
		return Scissors
	}
	panic(fmt.Sprintf("%s is not one of [ABCXYZ]", string(c)))
}

func convertResult(c byte) Score {
	switch c {
	case 'X':
		return Loss
	case 'Y':
		return Draw
	case 'Z':
		return Win
	}
	panic(fmt.Sprintf("%s is not one of [XYZ]", string(c)))
}

func play1(s string) Score {
	opponent := convert(s[0])
	us := convert(s[2])
	switch us {
	case Rock:
		switch opponent {
		case Rock:
			return us + Draw
		case Paper:
			return us + Loss
		case Scissors:
			return us + Win
		}
		panic(s)
	case Paper:
		switch opponent {
		case Rock:
			return us + Win
		case Paper:
			return us + Draw
		case Scissors:
			return us + Loss
		}
		panic(s)
	case Scissors:
		switch opponent {
		case Rock:
			return us + Loss
		case Paper:
			return us + Win
		case Scissors:
			return us + Draw
		}
		panic(s)
	}
	panic(s)
}

func play2(s string) Score {
	opponent := convert(s[0])
	desiredResult := convertResult(s[2])
	switch desiredResult {
	case Loss:
		switch opponent {
		case Rock:
			return desiredResult + Scissors
		case Paper:
			return desiredResult + Rock
		case Scissors:
			return desiredResult + Paper
		}
		panic(s)
	case Draw:
		switch opponent {
		case Rock:
			return desiredResult + Rock
		case Paper:
			return desiredResult + Paper
		case Scissors:
			return desiredResult + Scissors
		}
		panic(s)
	case Win:
		switch opponent {
		case Rock:
			return desiredResult + Paper
		case Paper:
			return desiredResult + Scissors
		case Scissors:
			return desiredResult + Rock
		}
		panic(s)
	}
	panic(s)
}
