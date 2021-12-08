package main

import (
	"fmt"
	"strings"

	"github.com/j18e/adventofcode/pkg/converting"
)

func parseBoards(input string) []Board {
	lines := strings.Split(input, "\n")
	i := 0

	// get past lines with numbers
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
	}

	// get boards
	var boards []Board
	for i < len(lines) {
		if lines[i] == "" {
			i++
			continue
		}
		boards = append(boards, parseBoard(lines[i:i+5]))
		i += 5
	}

	return boards
}

func parseNumbers(input string) []int {
	lines := strings.Split(input, "\n")
	i := 0

	// get numbers
	var res []int
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
		for _, num := range strings.Split(lines[i], ",") {
			if num == "" {
				continue
			}
			res = append(res, converting.Atoi(num))
		}
	}
	return res
}

func parseBoard(lines []string) Board {
	if len(lines) != 5 {
		panic(fmt.Sprintf("need 5 lines in parseBoard, got %d", len(lines)))
	}
	var board Board
	for _, ln := range lines {
		board = append(board, parseBoardRow(ln))
	}
	return board
}

func parseBoardRow(input string) []int {
	var res []int
	for _, val := range strings.Split(input, " ") {
		if val == "" {
			continue
		}
		res = append(res, converting.Atoi(val))
	}
	return res
}
