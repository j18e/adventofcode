package converting

import (
	"strconv"
	"strings"
)

func Atoi(s string) int {
	s = strings.TrimSpace(s)
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

func Btoi(s string) int {
	s = strings.TrimSpace(s)
	res, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(res)
}
