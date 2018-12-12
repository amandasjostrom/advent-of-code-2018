package day11

import (
	"amandasjostrom.se/common"
	"strconv"
	"fmt"
)

type Square struct {
	topLeftX   int
	topLeftY   int
	totalPower int
	size       int
}

const maxSize = 300

func Run(input []string) common.Result {
	gridSerialNumber := common.ToInt(input[0])
	largestTotalPowerPart1 := part1(gridSerialNumber)
	largestTotalPowerPart2 := part2(gridSerialNumber)
	part1 := strconv.Itoa(largestTotalPowerPart1.topLeftX) + "," + strconv.Itoa(largestTotalPowerPart1.topLeftY)
	part2 := strconv.Itoa(largestTotalPowerPart2.topLeftX) + "," + strconv.Itoa(largestTotalPowerPart2.topLeftY) + "," + strconv.Itoa(largestTotalPowerPart2.size)
	fmt.Println("Part1 largest was: ", part1)
	fmt.Println("Part2 largest was: ", part2)
	return common.Result{Part1: part1, Part2: part2}
}
func part2(gridSerialNumber int) Square {
	largestTotalPower := Square{}
	for x := 1; x <= maxSize; x++ {
		for y := 1; y <= maxSize; y++ {
			totalPower, size := squarePowerAnySize(x, y, gridSerialNumber)
			if totalPower > largestTotalPower.totalPower {
				largestTotalPower = Square{x, y, totalPower, size}
			}
		}
	}
	return largestTotalPower
}
func part1(gridSerialNumber int) Square {
	largestTotalPower := Square{}
	for x := 1; x <= maxSize; x++ {
		for y := 1; y <= maxSize; y++ {
			totalPower := squarePower3By3(x, y, gridSerialNumber)
			if totalPower > largestTotalPower.totalPower {
				largestTotalPower = Square{x, y, totalPower, 3}
			}
		}
	}
	return largestTotalPower
}

func squarePower3By3(topLeftX int, topLeftY int, gridSerialNumber int) int {
	return squarePower(topLeftX, topLeftY, gridSerialNumber, 3)
}

func squarePower(topLeftX int, topLeftY int, gridSerialNumber int, size int) int {
	sum := 0
	for x := topLeftX; x < topLeftX+size && x <= maxSize; x++ {
		for y := topLeftY; y < topLeftY+size && y <= maxSize; y++ {
			power := cellPower(x, y, gridSerialNumber)
			sum += power
		}
	}
	return sum
}

func squarePowerAnySize(topLeftX int, topLeftY int, gridSerialNumber int) (power int, size int) {
	largestPower := 0
	largestPowerSize := 0
	currentSize := 1
	for topLeftX+currentSize-1 <= maxSize && topLeftY+currentSize-1 <= maxSize {
		power := squarePower(topLeftX, topLeftY, gridSerialNumber, currentSize)
		if power > largestPower {
			largestPower = power
			largestPowerSize = currentSize
		}
		currentSize++
	}
	return largestPower, largestPowerSize
}

func cellPower(x int, y int, gridSerialNumber int) int {
	rackId := x + 10
	startingFuelLevel := rackId * y
	powerLevel := (startingFuelLevel + gridSerialNumber) * rackId
	onlyHundredsDigit := (powerLevel / 100) % 10
	power := onlyHundredsDigit - 5
	return power
}
