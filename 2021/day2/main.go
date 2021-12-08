package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/j18e/adventofcode/pkg/inputting"
)

type Direction int

const (
	Forward Direction = iota
	Up
	Down
)

type Command struct {
	Direction Direction
	Units     int
}

func main() {
	input := inputting.GetInputStrings("input.txt")
	part1(input)
	part2(input)
}

func part1(input []string) {
	pos, depth := 0, 0
	for _, cmd := range getCommands(input) {
		switch cmd.Direction {
		case Forward:
			pos += cmd.Units
		case Up:
			depth -= cmd.Units
		case Down:
			depth += cmd.Units
		}
	}
	fmt.Println(pos * depth)
}

func part2(input []string) {
	var pos, depth, aim int
	for _, cmd := range getCommands(input) {
		switch cmd.Direction {
		case Forward:
			pos += cmd.Units
			depth += aim * cmd.Units
		case Up:
			aim -= cmd.Units
		case Down:
			aim += cmd.Units
		}
	}
	fmt.Println(pos * depth)
}

func getCommands(input []string) []Command {
	var commands []Command
	for _, val := range input {
		split := strings.Split(val, " ")
		if len(split) != 2 {
			panic(fmt.Sprintf("%s: expected 2 words, got %d", val, len(split)))
		}
		var dir Direction
		switch split[0] {
		case "forward":
			dir = Forward
		case "up":
			dir = Up
		case "down":
			dir = Down
		default:
			panic(fmt.Sprintf("value %s: expected one of [forward up down]", split[0]))
		}
		units, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		commands = append(commands, Command{dir, units})
	}
	return commands
}
