package day1

import (
	"fmt"
	"amandasjostrom.se/common"
)
func Run() {
	values := common.ToIntSlice(common.GetInput(1, "\n"))
	sum := int(0)
	seenFrequencies := make([]int, 0)
	foundFirstPart := false
	index := 0
	for ; ; {
		sum+= values[index]
		if !foundFirstPart && index == len (values)-1{
			fmt.Println("Answer part one: ", sum)
			foundFirstPart = true
		}
		index = common.GetNextIndex(index, len(values))

		if seenFrequency(seenFrequencies, sum){
			fmt.Print("Seen this before! Answer part two: ", sum)
			break
		}
		seenFrequencies = append(seenFrequencies, sum)
	}
}

func seenFrequency(frequencies []int, frequency int) bool {
	for _, value := range frequencies {
		if value == frequency {
			return true
		}
	}
	return false
}