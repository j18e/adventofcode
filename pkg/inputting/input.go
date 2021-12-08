package inputting

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func GetInput(name string) string {
	bs, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

func GetInputStrings(name string) []string {
	var res []string
	for _, ln := range strings.Split(GetInput(name), "\n") {
		if ln == "" {
			continue
		}
		res = append(res, ln)
	}

	return res
}

func GetInputInts(name string) []int {
	var res []int
	for _, ln := range GetInputStrings(name) {
		n, err := strconv.Atoi(ln)
		if err != nil {
			panic(err)
		}
		res = append(res, n)
	}

	return res
}
