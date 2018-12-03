package common

import "strconv"

func ToInt(input string) int {
	val, _ := strconv.ParseInt(input, 10, 64)
	return int(val)
}