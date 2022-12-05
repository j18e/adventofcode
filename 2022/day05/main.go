package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

var reMoveLine = regexp.MustCompile(`^move (\d+) from (\d) to (\d)`)

func run() error {
	bs, err := os.ReadFile("input.txt")
	if err != nil {
		return err
	}
	var crateLines, moveLines []string
	for _, ln := range strings.Split(string(bs), "\n") {
		if len(ln) == 0 {
			continue
		}
		switch ln[0] {
		case '[':
			crateLines = append(crateLines, ln)
		case 'm':
			moveLines = append(moveLines, ln)
		}
	}
	numStacks := 0
	for _, char := range crateLines[len(crateLines)-1] {
		if char >= 'A' && char <= 'Z' {
			numStacks++
		}
	}

	stacks1, stacks2 := make(Stacks, numStacks), make(Stacks, numStacks)
	for i := 0; i < numStacks; i++ {
		stacks1[i] = NewStack(crateLines, i)
		stacks2[i] = NewStack(crateLines, i)
	}

	for _, ln := range moveLines {
		move := parseMove(ln)
		stacks1.Move9000(move)
		stacks2.Move9001(move)
	}
	fmt.Println(stacks1.Tops())
	fmt.Println(stacks2.Tops())
	return nil
}

func (sx Stacks) Move9000(move Move) {
	from, to := move.From, move.To
	for i := 0; i < move.Cnt; i++ {
		r := sx[from][len(sx[from])-1]
		sx[from] = sx[from][:len(sx[from])-1]
		sx[to] = append(sx[to], r)
	}
}

func (sx Stacks) Move9001(move Move) {
	cnt, from, to := move.Cnt, move.From, move.To
	fromLen := len(sx[from])
	runes := sx[from][fromLen-cnt:]
	sx[from] = sx[from][:fromLen-cnt]
	sx[to] = append(sx[to], runes...)
}

func (sx Stacks) Tops() string {
	var res string
	for _, stack := range sx {
		if len(stack) == 0 {
			continue
		}
		res += string(stack[len(stack)-1])
	}
	return res
}

type Stacks []Stack

func (sx Stacks) String() string {
	highest := 0
	for _, stack := range sx {
		if l := len(stack); l > highest {
			highest = l
		}
	}

	var lines []string
	for i := highest - 1; i >= 0; i-- {
		line := ""
		for _, stack := range sx {
			if len(stack) < i+1 {
				line += "   "
			} else {
				line += "  " + string(stack[i])
			}
		}
		lines = append(lines, strings.TrimSpace(line))
	}

	var divider string
	var numLine string
	for i := 0; i < len(sx); i++ {
		numLine += fmt.Sprintf("  %d", i+1)
		divider += "-  "
	}
	numLine = strings.TrimSpace(numLine)

	return strings.Join(lines, "\n") + "\n" + divider + "\n" + numLine
}

type Stack []rune

func (c Stack) String() string {
	var res string
	for _, r := range c {
		res += string(r) + " "
	}
	return strings.TrimSpace(res)
}

func NewStack(lines []string, stack int) Stack {
	var res []rune
	for i := len(lines) - 1; i >= 0; i-- {
		ln := lines[i]
		r := ln[stack*4+1]
		if r == ' ' {
			break
		}
		if r < 'A' || r > 'Z' {
			panic(fmt.Sprintf("rune %s should be upperalpha or space", string(r)))
		}
		res = append(res, rune(r))
	}
	return res
}

type Move struct {
	Text          string
	Cnt, From, To int
}

func (m Move) String() string {
	return m.Text
}

func parseMove(ln string) Move {
	matches := reMoveLine.FindStringSubmatch(ln)
	if len(matches) != 4 {
		panic(ln)
	}
	cnt, err1 := strconv.Atoi(matches[1])
	from, err2 := strconv.Atoi(matches[2])
	to, err3 := strconv.Atoi(matches[3])
	check(fmt.Sprintf("parsing %s: ", ln), err1, err2, err3)

	// subtract 1 from to and from since stack number starts from 0
	return Move{strings.TrimSpace(ln), cnt, from - 1, to - 1}
}

func check(msg string, ex ...error) {
	for _, err := range ex {
		if err != nil {
			panic(msg + err.Error())
		}
	}
}
