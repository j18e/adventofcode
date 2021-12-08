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

	lines := strings.Split(string(bs), "\n")
	earliest, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal(err)
	}

	steps := make(StepSet)
	for i, time := range strings.Split(lines[1], ",") {
		if time == "x" {
			continue
		}
		bus, err := strconv.Atoi(time)
		if err != nil {
			log.Fatal(err)
		}
		steps[i] = bus
	}

	fmt.Println(steps.EarliestDeparture(earliest))
	fmt.Println(steps.Multiplier())
}

type Departure struct {
	Bus, Time int
}

type StepSet map[int]int

func (ss StepSet) EarliestDeparture(earliest int) int {
	dep := Departure{-1, -1}
	for _, bus := range ss {
		time := earliest / bus * bus
		if earliest%bus != 0 {
			time += bus
		}
		if dep.Time == -1 || time < dep.Time {
			dep = Departure{bus, time}
		}
	}
	return dep.Bus * (dep.Time - earliest)
}

func (ss StepSet) Multiplier() int {
	timestamp, step := 0, 1
	for offset, bus := range ss {
		for (timestamp+offset)%bus != 0 {
			timestamp += step
		}
		step *= bus
	}
	return timestamp
}
