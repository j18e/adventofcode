package input

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func GetInput(name string) []int {
	bs, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	var input []int
	for _, ln := range strings.Split(string(bs), "\n") {
		if ln == "" {
			continue
		}
		n, err := strconv.Atoi(ln)
		if err != nil {
			panic(err)
		}
		input = append(input, n)
	}

	return input
}
