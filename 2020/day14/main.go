package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

var (
	reMask = regexp.MustCompile(`^mask = ([01X]{36})$`)
	reMem  = regexp.MustCompile(`mem\[(\d+)\] = (\d+)$`)
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type InstructionSet struct {
	Mask        Mask
	Assignments []Assignment
}

type Assignment struct {
	Address, Value int
}

type Program map[int]int

func (p Program) Sum() int {
	var sum int
	for _, val := range p {
		sum += val
	}
	return sum
}

type Mask string

func NewMask(s string) (Mask, error) {
	if len(s) != 36 {
		return "", errors.New("mask is not 36 chars long")
	}
	for _, r := range s {
		if !(r == '0' || r == '1' || r == 'X') {
			return "", errors.New("mask should only contain [01X]")
		}
	}
	return Mask(s), nil
}

func getBinary(num int) (string, error) {
	bin := strconv.FormatInt(int64(num), 2)
	if len(bin) > 36 {
		return "", fmt.Errorf("binary conversion is more than 36 bits long: %s", bin)
	}
	for len(bin) < 36 {
		bin = "0" + bin
	}
	return bin, nil
}

func (m Mask) CountX() int {
	cnt := 0
	for _, r := range m {
		if r == 'X' {
			cnt++
		}
	}
	return cnt
}

func (m Mask) ApplyToAddress(addr int) ([]int, error) {
	bin, err := getBinary(addr)
	if err != nil {
		return nil, err
	}
	var applied string
	for i := range bin {
		switch m[i] {
		case '1', 'X':
			applied += string(m[i])
		default:
			applied += string(bin[i])
		}
	}
	addrMap := make(map[string]bool)
	listPossibles(applied, addrMap)
	var results []int
	for s := range addrMap {
		i, err := strconv.ParseInt(s, 2, 64)
		if err != nil {
			return nil, err
		}
		results = append(results, int(i))
	}
	return results, nil
}

func listPossibles(s string, results map[string]bool) {
	if results[s] {
		return
	}
	x := strings.Index(s, "X")
	if x == -1 {
		results[s] = true
		return
	}
	listPossibles(fmt.Sprintf("%s0%s", s[:x], s[x+1:]), results)
	listPossibles(fmt.Sprintf("%s1%s", s[:x], s[x+1:]), results)
}

func (m Mask) ApplyToValue(n int) (int, error) {
	bin, err := getBinary(n)
	if err != nil {
		return 0, err
	}
	var resStr string
	for i := range bin {
		switch m[i] {
		case '0':
			resStr += "0"
		case '1':
			resStr += "1"
		default:
			resStr += string(bin[i])
		}
	}
	res, err := strconv.ParseInt(resStr, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(res), nil
}

func run() error {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sets []InstructionSet
	for _, l := range strings.Split(string(bs), "\n") {
		if l == "" {
			continue
		}
		if match := reMask.FindStringSubmatch(l); len(match) > 1 {
			mask, err := NewMask(match[1])
			if err != nil {
				return err
			}
			sets = append(sets, InstructionSet{Mask: mask})
			continue
		}
		if !reMem.MatchString(l) {
			return fmt.Errorf("line matches neither mask nor mem: %s", l)
		}

		matches := reMem.FindStringSubmatch(l)
		val, err := strconv.Atoi(matches[2])
		if err != nil {
			return fmt.Errorf("converting %s: %w", matches[2], err)
		}
		addr, err := strconv.Atoi(matches[1])
		if err != nil {
			return fmt.Errorf("converting %s: %w", matches[1], err)
		}
		sets[len(sets)-1].Assignments = append(
			sets[len(sets)-1].Assignments,
			Assignment{Address: addr, Value: val},
		)
	}

	if err := part1(sets); err != nil {
		return err
	}

	return part2(sets)
}

func part1(sets []InstructionSet) error {
	program := make(Program)
	for _, set := range sets {
		mask := set.Mask
		for _, a := range set.Assignments {
			res, err := mask.ApplyToValue(a.Value)
			if err != nil {
				return fmt.Errorf("applying %d: %w", a.Value, err)
			}
			program[a.Address] = res
		}
	}

	fmt.Println(program.Sum())
	return nil
}

func part2(sets []InstructionSet) error {
	program := make(Program)
	var wg sync.WaitGroup
	for _, set := range sets {
		for _, ass := range set.Assignments {
			ax, err := set.Mask.ApplyToAddress(ass.Address)
			if err != nil {
				return fmt.Errorf("applying %d: %v", ass.Address, err)
			}
			for _, add := range ax {
				program[add] = ass.Value
			}
		}
	}

	wg.Wait()
	fmt.Println(program.Sum())
	return nil
}
