package common

import "strconv"

func ToIntSlice(values []string) (numbers []int) {
	for _, v := range values {
		number, _ := strconv.ParseInt(v, 10, 64)
		numbers = append(numbers, int(number))
	}
	return
}

func GetNextIndex(currentIndex int, listLength int) int {
	return (currentIndex +1 ) % listLength
}
