package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var data []string
	for _, l := range strings.Split(string(bs), "\n") {
		if l == "" {
			continue
		}
		data = append(data, l)
	}
}
