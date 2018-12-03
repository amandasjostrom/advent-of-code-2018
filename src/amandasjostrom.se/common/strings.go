package common

import "strconv"

func NumberOfOccurences(find string, list []string) (sum int) {
	for _, value := range list {
		if value == find {
			sum = sum + 1
		}
	}
	return
}

func ToInt(input string) int {
	val, _ := strconv.ParseInt(input, 10, 64)
	return int(val)
}