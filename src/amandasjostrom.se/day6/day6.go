package day6

import (
	"amandasjostrom.se/common"
	strings2 "strings"
	"fmt"
	"strconv"
)

type Coordinate struct {
	x  int
	y  int
	id string
}

func Run(input []string) common.Result {
	originals, largestX, largestY := parse(input)
	var sizePerId = map[string]int{}
	var infiniteIds = map[string]bool{}
	var grid = map[string]string{}
	sizeOfRegion := 0
	for x := 0; x <= largestX; x++ {
		for y := 0; y <= largestY; y++ {
			sum := 0
			nearest := Coordinate{0, 0, "0"}
			minDistance := 9999
			onlyOneFound := true
			for _, value := range originals {
				distance := getManhattan(value, x, y)
				sum+=distance
				if distance < minDistance {
					onlyOneFound = true
					minDistance = distance
					nearest = value
				} else if distance == minDistance {
					onlyOneFound = false
				}
			}
			key := toKey(x, y)
			if onlyOneFound {
				sizePerId[nearest.id]++
				grid[key] = nearest.id
				if x == 0 || x == largestX+1 || y == 0 || y == largestY+1 {
					//remove infinite
					infiniteIds[nearest.id] = true
				}
			} else {
				grid[key] = "."
			}
			if sum < 10000 {
				sizeOfRegion++
			}
		}
	}
	fmt.Println("Dessa var på gränsen: ", infiniteIds)
	fmt.Println("size per id: ", sizePerId)
	printGrid(grid, largestX, largestY)

	largestSize := -1
	for k, v := range sizePerId {
		if infiniteIds[k] {
			continue
		}
		if v > largestSize {
			largestSize = v
		}

	}

	fmt.Println("Ehmm jag hittade detta på part 1 men bli inte arg om det är fel, okej...? Bra! Här: ", largestSize)
	return common.Result{largestSize, sizeOfRegion}
}
func toKey(x int, y int) string {
	return strconv.Itoa(x) + "+" + strconv.Itoa(y)
}
func printGrid(grid map[string]string, largestX int, largestY int) {
	for y := 0; y <= largestY; y++ {
		for x := 0; x <= largestX; x++ {
			fmt.Print(grid[toKey(x, y)])
		}
		fmt.Println()
	}
}
func getManhattan(coordinate Coordinate, x int, y int) int {
	xDiff := abs(coordinate.x - x)
	yDiff := abs(coordinate.y - y)
	return yDiff + xDiff
}
func abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

func parse(strings []string) (originals map[string]Coordinate, largestX int, largestY int) {
	alphabet := "abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ"
	originals = map[string]Coordinate{}
	largestX = -1
	largestY = -1
	for index, v := range strings {
		split := strings2.Split(v, ", ")
		x := common.ToInt(split[0])
		y := common.ToInt(split[1])
		id := string(alphabet[index])
		originals[id] = Coordinate{x, y, id}

		if x > largestX {
			largestX = x
		}
		if y > largestY {
			largestY = y
		}
	}
	return originals, largestX, largestY
}
