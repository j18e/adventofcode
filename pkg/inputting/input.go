package inputting

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func GetInput(name string) []string {
	bs, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	var input []string
	for _, ln := range strings.Split(string(bs), "\n") {
		if ln == "" {
			continue
		}
		input = append(input, ln)
	}

	return input
}

func GetInputInts(name string) []int {
	var input []int
	for _, ln := range GetInput(name) {
		n, err := strconv.Atoi(ln)
		if err != nil {
			panic(err)
		}
		input = append(input, n)
	}

	return input
}
