package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	reRoot     = regexp.MustCompile(`^(\w+ \w+) bags contain (.*)\.$`)
	reEmpty    = regexp.MustCompile(`^no other bags$`)
	reContents = regexp.MustCompile(`((\d) (\w+ \w+) bags?(, )?)`)
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	bags := make(BagList)
	for _, ln := range strings.Split(string(bs), "\n") {
		if ln == "" {
			continue
		}
		bags.AddBag(ln)
	}

	cnt := 0
	for c := range bags {
		if bags.containsShinyGold(c) {
			cnt++
		}
	}
	fmt.Println(cnt - 1)

	fmt.Println(bags.CountBags("shiny gold") - 1)
}

type BagList map[string]map[string]int

func (bags BagList) CountBags(colour string) int {
	total := 1
	if bags[colour] == nil {
		return total
	}
	for c, cnt := range bags[colour] {
		total += bags.CountBags(c) * cnt
	}
	return total
}

func (bags BagList) containsShinyGold(colour string) bool {
	if colour == "shiny gold" {
		return true
	}
	for c := range bags[colour] {
		if bags.containsShinyGold(c) {
			return true
		}
	}
	return false
}

func (bags BagList) AddBag(s string) {
	matches := reRoot.FindStringSubmatch(s)
	rootBag := matches[1]
	bagContents := matches[2]
	if reEmpty.MatchString(bagContents) {
		return
	}
	bags[rootBag] = make(map[string]int)
	for _, m := range reContents.FindAllStringSubmatch(bagContents, -1) {
		i, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		bags[rootBag][m[3]] = i
	}
}
