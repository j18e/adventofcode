package main

import (
	"fmt"

	"github.com/j18e/adventofcode/pkg/inputting"
)

func main() {
	input := inputting.GetInput("input.txt")
	nums := parseNumbers(input)
	boards := parseBoards(input)
	fmt.Println(part1(nums, boards))
	fmt.Println(part2(nums, boards))
}

func part1(nums []int, boards []Board) int {
	lowestRound := len(nums)
	lowestRoundScore := 0
	for _, board := range boards {
		round, score := board.Play(nums)
		if round < lowestRound {
			lowestRound = round
			lowestRoundScore = score
		}
	}
	return lowestRoundScore
}

func part2(nums []int, boards []Board) int {
	highestRound := 0
	highestRoundScore := 0
	for _, board := range boards {
		round, score := board.Play(nums)
		if round > highestRound {
			highestRound = round
			highestRoundScore = score
		}
	}
	return highestRoundScore
}

type Board [][]int

func (b Board) All() []int {
	var res []int
	for _, row := range b {
		for _, n := range row {
			res = append(res, n)
		}
	}
	return res
}

func (b Board) Has(num int) bool {
	for _, row := range b {
		for _, n := range row {
			if n == num {
				return true
			}
		}
	}
	return false
}

func (b Board) Play(nums []int) (round, score int) {
	var drawn []int
	for i := 0; i < len(nums); i++ {
		drawn = append(drawn, nums[i])
		if won(drawn, b) {
			return i, totalScore(drawn, b)
		}
	}
	return -1, -1
}

func totalScore(drawn []int, board Board) int {
	var sum int
	final := drawn[len(drawn)-1]
	for _, num := range board.All() {
		if !intInSlice(num, drawn) {
			sum += num
		}
	}
	return sum * final
}

func won(drawn []int, board Board) bool {
	// check rows
	for _, row := range board {
		if allIntsInSlice(row, drawn) {
			return true
		}
	}

	// check columns
	cols := len(board[0])
	rows := len(board)
	for i := 0; i < cols; i++ {
		matches := 0
		for _, row := range board {
			if intInSlice(row[i], drawn) {
				matches++
			}
		}
		if matches == rows {
			return true
		}
	}
	return false
}

func allIntsInSlice(ints, ix []int) bool {
	for _, i := range ints {
		if !intInSlice(i, ix) {
			return false
		}
	}
	return true
}

func intInSlice(i int, ix []int) bool {
	for _, ii := range ix {
		if i == ii {
			return true
		}
	}
	return false
}
