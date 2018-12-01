package main

import (
	"strings"
	"strconv"
	//"github.com/amandasjostrom/advent-of-code-2017"
	"log"
	"fmt"
	"os"
	"io/ioutil"
	//"reflect"
)
func main() {
	part2()
}

func part1() {
	absolutePath, _ := os.Getwd()
	values := fromFile(absolutePath + "/src/github.com/amandasjostrom/advent-of-code-2017/day7/input.txt", "\n")
	sum := int64(0)
	seenFrequencies := make([]int64, 0)
	for _, value := range values {
		//fmt.Println(value)
		number, _ := strconv.ParseInt(value, 10, 64)
		//fmt.Println("index: ", i, " sum: ", sum, " value: ", number)
		sum+=number

		//fmt.Println("Have seen these: ", seenFrequencies)
		//if seenFrequency(seenFrequencies, sum){
		//	fmt.Print("Seen this before: ", sum)
		//	break
		//}

		seenFrequencies = append(seenFrequencies, sum)
	}
	fmt.Println("RESULTAT: ", sum)
}

func part2() {
	absolutePath, _ := os.Getwd()
	values := toIntSplice(fromFile(absolutePath + "/src/github.com/amandasjostrom/advent-of-code-2017/day7/input.txt", "\n"))
	sum := int64(0)
	seenFrequencies := make([]int64, 0)
	index := 0
	for ; ; {
		//fmt.Println(value)
		//fmt.Println("index: ", i, " sum: ", sum, " value: ", number)
		sum+= values[index]
		index = getNextIndex(index, len(values)-1)

		//fmt.Println("Have seen these: ", seenFrequencies)
		if seenFrequency(seenFrequencies, sum){
			fmt.Print("Seen this before: ", sum)
			break
		}

		seenFrequencies = append(seenFrequencies, sum)
	}
}
func toIntSplice(values []string) (numbers []int64) {
	for _, v := range values {
		number, _ := strconv.ParseInt(v, 10, 64)
		numbers = append(numbers, number)
	}
	return
}
func getNextIndex(currentIndex int, maxIndex int) int {
	if currentIndex+1 > maxIndex {
		return 0
	} else {
		return currentIndex + 1
	}
}

func seenFrequency(frequencies []int64, frequency int64) bool {
	for _, value := range frequencies {
		if(value == frequency){

		//}
		//if reflect.DeepEqual(value, frequency) {
			return true
		}
	}
	return false
}

func fromFile(filename string, separator string) []string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(":(")
		log.Fatal(err)
	}
	//fmt.Print(":)")
	values := strings.Split(string(file), separator)
	//fmt.Print(values)
	return values
}