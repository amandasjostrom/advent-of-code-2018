package day2

import (
	"strings"
	"fmt"
	"amandasjostrom.se/common"
)

func Run() {
	boxIds := common.GetInput(2, "\n")
	fmt.Println("DAY 2 Checksum: ", part1(boxIds))
	fmt.Println("DAY 2 FOUND IT: ", part2(boxIds))
}

func part2(boxIds []string) string {
	for i := range boxIds {
		//for each box
		for j := i + 1; j < len(boxIds); j++ {
			//compare to the other boxes
			moreThanOneDiff, charactersInCommon := hasMoreThanOneDiff(boxIds[i], boxIds[j])
			if moreThanOneDiff {
				continue
			} else {
				return charactersInCommon
			}
		}
	}
	return ""
}

func hasMoreThanOneDiff(box1 string, box2 string) (moreThanOneDiff bool, charactersInCommon string) {
	foundFirstDiff := false
	for index := range box1 {
		if box1[index] == box2[index] {
			charactersInCommon = charactersInCommon + string(box1[index])
		} else {
			if foundFirstDiff {
				return true, ""
			} else {
				foundFirstDiff = true
			}
		}
	}
	return false, charactersInCommon
}

func part1(boxIds []string) int {
	sumExactlyTwo, sumExactlyThree := 0, 0
	for _, id := range boxIds {
		letters := strings.Split(id, "")
		foundTwo, foundThree := false, false
		for _, letter := range letters {
			number := strings.Count(id, string(letter))
			if number == 2 {
				foundTwo = true
			} else if number == 3 {
				foundThree = true
			}
		}
		if foundTwo {
			sumExactlyTwo = sumExactlyTwo + 1
		}
		if foundThree {
			sumExactlyThree = sumExactlyThree + 1
		}
	}
	fmt.Println("Sum duplicates: ", sumExactlyTwo, "; Sum triplets: ", sumExactlyThree)
	return sumExactlyTwo * sumExactlyThree
}

