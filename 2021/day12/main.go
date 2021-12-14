package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/j18e/adventofcode/pkg/inputting"
)

const (
	Start = "start"
	End   = "end"
)

var reBigCave = regexp.MustCompile(`^[A-Z]+$`)

func main() {
	input := inputting.GetInputStrings("input.txt")
	fmt.Println(part1(input))
	// fmt.Println(part2(input))
}

func parseInput(input []string) Network {
	caves := make(map[string]*Cave)
	for _, ln := range input {
		split := strings.Split(ln, "-")
		c0, c1 := split[0], split[1]
		if caves[c0] == nil {
			caves[c0] = NewCave(c0)
		}
		if caves[c1] == nil {
			caves[c1] = NewCave(c1)
		}
		caves[c0].Neighbours[caves[c1]] = true
		if c1 != End {
			caves[c1].Neighbours[caves[c0]] = true
		}
	}
	var res Network
	for _, c := range caves {
		res = append(res, c)
	}
	return res
}

func part1(input []string) int {
	network := parseInput(input)
	paths := network.Paths(Path{network.Start()})
	sortPaths(paths)
	for _, p := range paths {
		fmt.Println(p)
	}
	return len(paths)
}

func part2(input string) int {
	return 0
}

type Network []*Cave

func (n Network) Start() *Cave {
	for _, c := range n {
		if c.Name == Start {
			return c
		}
	}
	panic("start cave not found")
}

type Path []*Cave

func (p Path) String() string {
	var res string
	for _, c := range p {
		res += " " + c.Name
	}
	return strings.TrimSpace(res)
}

func (p Path) Names() []string {
	var res []string
	for _, c := range p {
		res = append(res, c.Name)
	}
	return res
}

func NewCave(name string) *Cave {
	return &Cave{
		Name:       name,
		Neighbours: make(map[*Cave]bool),
	}
}

type Cave struct {
	Name       string
	Neighbours map[*Cave]bool
}

func (c *Cave) IsEnd() bool {
	return c.Name == End
}

func (c *Cave) IsBig() bool {
	return reBigCave.MatchString(c.Name)
}

func (n Network) Paths(path Path) []Path {
	start := path[len(path)-1]
	var res []Path
	if start.IsEnd() {
		// return append(res, path)
		return nil
	}
	for c := range start.Neighbours {
		if visited(path, c) && !c.IsBig() {
			continue
		}
		np := append(path, c)
		if c.Name == End {
			res = append(res, np)
		} else {
			newPaths := n.Paths(np)
			res = append(res, newPaths...)
		}
	}
	return res
}

func visited(path Path, cave *Cave) bool {
	for _, c := range path {
		if c == cave {
			return true
		}
	}
	return false
}

func sortPaths(paths []Path) {
	sort.Slice(paths, func(i, j int) bool {
		p1, p2 := paths[i], paths[j]
		for i := range p1 {
			if len(p2)-1 < i {
				continue
			}
			if p1[i] == p2[i] {
				if len(p1)-1 == i {
					return true
				}
				continue
			}
			return p1[i].Name < p2[i].Name
		}
		return false
	})
}
