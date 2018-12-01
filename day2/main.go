package main

import (
	"strings"
	"fmt"
	"../common"
)

func main() {
	boxIds := common.GetInput(2, "\n")
	fmt.Println("Checksum: ", part1(boxIds))
	fmt.Print("FOUND IT: ", part2(boxIds))
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
	sumExactlyTwo := int(0)
	sumExactlyThree := int(0)
	for _, id := range boxIds {
		letters := strings.Split(id, "")
		foundTwo := false
		foundThree := false
		for _, letter := range letters {
			number := numberOfInstances(letter, letters)
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

func numberOfInstances(letterToFind string, letters []string) (sum int) {
	for _, letter := range letters {
		if letter == letterToFind {
			sum = sum + 1
		}
	}
	return
}

