package main

import (
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := input.GetInput("input.txt")
	part1(input)
	part2(input)
	return nil
}

func part1(input []int) {
	res := 0
	fmt.Println(res)
}

func part2(input []int) {
	res := 0
	fmt.Println(res)
}
