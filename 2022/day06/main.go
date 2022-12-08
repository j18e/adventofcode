package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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
	input := strings.TrimSpace(string(bs))
	fmt.Println("part 1:", findMarker(input, 4))
	fmt.Println("part 2:", findMarker(input, 14))
	return nil
}

func findMarker(input string, markerLen int) int {
	for i := range input[markerLen-1:] {
		if i < markerLen-1 {
			continue
		}
		if allUnique(input[i-markerLen+1 : i+1]) {
			return i + 1
		}
	}
	return -1
}

func allUnique(chars string) bool {
	seen := make(map[rune]bool)
	for _, c := range chars {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}
