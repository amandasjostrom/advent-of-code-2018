package main

import (
	"strings"
	"strconv"
	"log"
	"fmt"
	"io/ioutil"
	"../common"
)
func main() {
	values := toIntSlice(common.GetInput(1, "\n"))
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
		index = getNextIndex(index, len(values))

		if seenFrequency(seenFrequencies, sum){
			fmt.Print("Seen this before! Answer part two: ", sum)
			break
		}
		seenFrequencies = append(seenFrequencies, sum)
	}
}

func toIntSlice(values []string) (numbers []int) {
	for _, v := range values {
		number, _ := strconv.ParseInt(v, 10, 64)
		numbers = append(numbers, int(number))
	}
	return
}
func getNextIndex(currentIndex int, maxIndex int) int {
	return (currentIndex +1 ) % maxIndex
}

func seenFrequency(frequencies []int, frequency int) bool {
	for _, value := range frequencies {
		if value == frequency {
			return true
		}
	}
	return false
}

func fromFile(filename string, separator string) (values []string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(":(")
		log.Fatal(err)
	}
	values = strings.Split(string(file), separator)
	return
}