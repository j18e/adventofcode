package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var valid1, valid2 int
	for _, l := range strings.Split(string(bs), "\n") {
		if l == "" {
			continue
		}
		entry := parseEntry(l)
		if entry.Valid1() {
			valid1++
		}
		if entry.Valid2() {
			valid2++
		}
	}
	fmt.Println("valid1:", valid1)
	fmt.Println("valid2:", valid2)
}

func parseEntry(line string) Entry {
	minMax := strings.Split(strings.Split(line, " ")[0], "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])
	policy := Policy{
		Letter: rune(strings.Split(line, " ")[1][0]),
		Min:    min,
		Max:    max,
	}
	pw := strings.Split(line, ": ")[1]
	letters := make(map[rune]int)
	for _, b := range pw {
		letters[b]++
	}
	return Entry{policy, pw, letters}
}

type Entry struct {
	Policy
	Password     string
	LetterCounts map[rune]int
}

type Policy struct {
	Letter   rune
	Min, Max int
}

func (e Entry) Valid1() bool {
	if e.Min > e.LetterCounts[e.Letter] {
		return false
	}
	if e.Max < e.LetterCounts[e.Letter] {
		return false
	}
	return true
}

func (e Entry) Valid2() bool {
	l1 := rune(e.Password[e.Min-1])
	l2 := rune(e.Password[e.Max-1])
	return (l1 == e.Letter && l2 != e.Letter) || (l1 != e.Letter && l2 == e.Letter)
}
