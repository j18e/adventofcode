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
	Items     []int
	Operation func(i int) int
	Test      func(i int) int

	Inspections int
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := common.ReadInputString("input.txt")
	monkeys, lcm := parseInput(input)
	monkeyBusiness(monkeys, 20, func(i int) int { return i / 3 })
	monkeys, lcm = parseInput(input)
	monkeyBusiness(monkeys, 10000, func(i int) int { return i % lcm })
	return nil
}

func monkeyBusiness(mx []Monkey, rounds int, op func(int) int) {
	for i := 0; i < rounds; i++ {
		runRound(mx, op)
	}
	var res []int
	for _, m := range mx {
		res = append(res, m.Inspections)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	fmt.Println(res[0] * res[1])
}

func runRound(mx []Monkey, op func(int) int) {
	for mi, m := range mx {
		for _, item := range m.Items {
			mx[mi].Inspections++
			item = op(m.Operation(item))
			throwTo := m.Test(item)
			mx[throwTo].Items = append(mx[throwTo].Items, item)
		}
		mx[mi].Items = nil
	}
}

func parseInput(input string) ([]Monkey, int) {
	monkeys := []Monkey{}
	lcm := 1
	for _, block := range strings.Split(input, "\n\n") {
		lines := strings.Split(block, "\n")
		monkey := Monkey{}
		var divisibleBy, ifTrue, ifFalse int
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
				divisibleBy = strToInt(strings.Split(ln, " ")[3])
			case 4:
				split := strings.Split(ln, " ")
				ifTrue = strToInt(split[5])
			case 5:
				split := strings.Split(ln, " ")
				ifFalse = strToInt(split[5])
			default:
				panic(fmt.Sprintf("unrecognized line %d: %s", i, ln))
			}
		}
		monkey.Test = func(i int) int {
			if i%divisibleBy == 0 {
				return ifTrue
			}
			return ifFalse
		}
		monkeys = append(monkeys, monkey)
		lcm *= divisibleBy
	}
	return monkeys, lcm
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
