package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/j18e/adventofcode/2022/common"
)

type Monkey struct {
	Items          []int
	Operation      func(i int) int
	DivisibleBy    int
	IfTrueThrowTo  int
	IfFalseThrowTo int

	Inspections int
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := common.ReadInputString("input.txt")
	monkeys := parseInput(input)
	monkeyBusiness(monkeys, 20, true)
	return nil
}

func monkeyBusiness(mx []Monkey, rounds int, divideBy3 bool) {
	for i := 0; i < rounds; i++ {
		runRound(mx, divideBy3)
	}
	var res []int
	for _, m := range mx {
		res = append(res, m.Inspections)
	}
	sort.Ints(res)
	first := res[len(res)-1]
	second := res[len(res)-2]
	fmt.Println(first * second)
}

func runRound(mx []Monkey, divideBy3 bool) {
	for mi, m := range mx {
		for _, item := range m.Items {
			mx[mi].Inspections++
			item = m.Operation(item)
			if divideBy3 {
				item /= 3
			}
			if item%m.DivisibleBy == 0 {
				mx[m.IfTrueThrowTo].Items = append(mx[m.IfTrueThrowTo].Items, item)
			} else {
				mx[m.IfFalseThrowTo].Items = append(mx[m.IfFalseThrowTo].Items, item)
			}
		}
		mx[mi].Items = []int{}
	}
}

func parseInput(input string) []Monkey {
	var monkeys []Monkey
	for _, block := range strings.Split(input, "\n\n") {
		lines := strings.Split(block, "\n")
		monkey := Monkey{}
		for i, ln := range lines {
			ln = strings.TrimSpace(ln)
			switch i {
			case 0:
			case 1:
				split := strings.Split(strings.Split(ln, ":")[1], ",")
				for _, str := range split {
					monkey.Items = append(monkey.Items, strToInt(str))
				}
			case 2:
				split := strings.Split(ln, " ")
				switch split[4] {
				case "+":
					if split[5] == "old" {
						monkey.Operation = func(i int) int {
							return i + i
						}
					} else {
						monkey.Operation = func(i int) int {
							return i + strToInt(split[5])
						}
					}
				case "*":
					if split[5] == "old" {
						monkey.Operation = func(i int) int {
							return i * i
						}
					} else {
						monkey.Operation = func(i int) int {
							return i * strToInt(split[5])
						}
					}
				default:
					panic(ln)
				}
			case 3:
				monkey.DivisibleBy = strToInt(strings.Split(ln, " ")[3])
			case 4:
				split := strings.Split(ln, " ")
				if split[1] != "true:" {
					panic("line[1] != true: " + split[1])
				}
				monkey.IfTrueThrowTo = strToInt(split[5])
			case 5:
				split := strings.Split(ln, " ")
				if split[1] != "false:" {
					panic(ln)
				}
				monkey.IfFalseThrowTo = strToInt(split[5])
			default:
				panic(fmt.Sprintf("unrecognized line %d: %s", i, ln))
			}
		}
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func strToInt(s string) int {
	num, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return num
}

var input = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`
