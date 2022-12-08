package common

import (
	"os"
	"strings"
)

func ReadInput(file string) []string {
	bs, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(bs)), "\n")
}
