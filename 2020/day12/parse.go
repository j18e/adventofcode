package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var reLine = regexp.MustCompile(`^([NSEWLRF])(\d+)$`)

func parse() []Instruction {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var ix []Instruction
	for _, l := range strings.Split(string(bs), "\n") {
		if len(l) == 0 {
			continue
		}
		if !reLine.MatchString(l) {
			panic(string(l) + " does not match regexp")
		}
		matches := reLine.FindStringSubmatch(l)
		if len(matches) != 3 {
			panic(l)
		}
		num, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		var dir Direction
		switch rune(matches[1][0]) {
		case 'N':
			dir = N
		case 'S':
			dir = S
		case 'E':
			dir = E
		case 'W':
			dir = W
		case 'L':
			dir = L
		case 'R':
			dir = R
		case 'F':
			dir = F
		default:
			panic(l)
		}
		ix = append(ix, Instruction{dir, num})
	}
	return ix
}
