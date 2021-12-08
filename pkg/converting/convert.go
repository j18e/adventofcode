package converting

import "strconv"

func Atoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

func Btoi(s string) int {
	res, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(res)
}
