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

	var lines []Instruction
	for _, ln := range strings.Split(string(bs), "\n") {
		if ln == "" {
			continue
		}
		split := strings.Split(ln, " ")
		arg, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		lines = append(lines, Instruction{Type: split[0], Arg: arg})
	}

	acc, _ := iterate(lines)
	fmt.Println("part 1:", acc)

	for i := range lines {
		orig := lines[i].Type
		switch orig {
		case "acc":
			continue
		case "jmp":
			lines[i].Type = "nop"
		case "nop":
			lines[i].Type = "jmp"
		}
		if acc, fr := iterate(lines); fr == -1 {
			fmt.Printf("part 2: had to change %d (%s %d), got acc %d\n", i, orig, lines[i].Arg, acc)
			return
		}
		lines[i].Type = orig
	}
	log.Fatal("part 2: we failed")
}

type Instruction struct {
	Type string
	Arg  int
}

func iterate(lines []Instruction) (acc, firstRepeat int) {
	stepRecord := make(map[int]int)
	for i := 0; i < len(lines); {
		if stepRecord[i] > 0 {
			return acc, i
		}
		stepRecord[i]++
		switch lines[i].Type {
		case "acc":
			acc += lines[i].Arg
		case "jmp":
			i += lines[i].Arg
			continue
		case "nop":
		}
		i++
	}
	return acc, -1
}
