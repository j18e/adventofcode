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
	letters := make(map[rune]int)
	lines := strings.Split(string(bs), "\n")
	for _, ln := range lines {
		for _, r := range common(splitRucksack(ln)) {
			letters[r]++
		}
	}
	score := 0
	for r, cnt := range letters {
		score += letterScore(r) * cnt
	}
	fmt.Println("part 1:", score)

	score = 0
	for i := 0; i < len(lines)-1; i += 3 {
		groupLines := [3]string{lines[i], lines[i+1], lines[i+2]}
		group := NewGroup(groupLines)
		score += letterScore(group.Badge())
	}
	fmt.Println("part 2:", score)
	return nil
}

func NewGroup(lines [3]string) Group {
	var g Group
	g[0] = make(map[rune]bool)
	for _, r := range lines[0] {
		g[0][r] = true
	}
	g[1] = make(map[rune]bool)
	for _, r := range lines[1] {
		g[1][r] = true
	}
	g[2] = make(map[rune]bool)
	for _, r := range lines[2] {
		g[2][r] = true
	}
	return g
}

type Group [3]map[rune]bool

func (g Group) Badge() rune {
	for r := range g[0] {
		if g[0][r] && g[1][r] && g[2][r] {
			return r
		}
	}
	panic("could not find group's badge")
}

func letterScore(l rune) int {
	if l >= 'a' && l <= 'z' {
		return int(l - 'a' + 1)
	}
	if l >= 'A' && l <= 'Z' {
		return int(l - 'A' + 27)
	}
	panic(fmt.Sprintf(`"%s" is not letter!`, string(rune(l))))
}

func splitRucksack(sack string) (map[rune]bool, map[rune]bool) {
	if len(sack)%2 != 0 {
		panic(fmt.Sprintf("rucksack %s is not even length", sack))
	}
	half := len(sack) / 2
	string1, string2 := sack[:half], sack[half:]
	map1, map2 := make(map[rune]bool), make(map[rune]bool)
	for _, r := range string1 {
		map1[rune(r)] = true
	}
	for _, r := range string2 {
		map2[rune(r)] = true
	}
	return map1, map2
}

func common(map1, map2 map[rune]bool) []rune {
	var res []rune
	for r := range map1 {
		if map2[r] {
			res = append(res, r)
		}
	}
	return res
}
