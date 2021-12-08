package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, gs := range strings.Split(string(bs), "\n\n") {
		if gs == "" {
			continue
		}
		grp := NewGroup(gs)
		fmt.Println(grp)
		sum += grp.SumAllYes()
	}
	fmt.Println(sum)
}

func NewGroup(gs string) *Group {
	group := Group{Letters: make(map[rune]int)}
	for _, ln := range strings.Split(gs, "\n") {
		group.NewMember(ln)
	}
	return &group
}

type Group struct {
	Members int
	Letters map[rune]int
}

func (g *Group) NewMember(ms string) {
	if ms == "" {
		return
	}
	g.Members++
	for _, r := range ms {
		if r < 'a' || r > 'z' {
			continue
		}
		g.Letters[r]++
	}
}

func (g *Group) AllYes(r rune) bool {
	return g.Letters[r] == g.Members
}

func (g *Group) SumAllYes() int {
	sum := 0
	for r := range g.Letters {
		if g.AllYes(r) {
			sum++
		}
	}
	return sum
}

func part1(groups []string) int {
	var sum int
	for _, l := range groups {
		letters := make(map[rune]bool)
		for _, r := range l {
			if r < 'a' || r > 'z' {
				continue
			}
			letters[r] = true
		}
		sum += len(letters)
	}
	return sum
}
