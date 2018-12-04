package day3

import (
	"amandasjostrom.se/common"
	"fmt"
	"strconv"
)

func Run(input []string) common.Result {
	claims := claims(input)
	fabric := markClaimsOnFabric(claims)
	squareInchesUsedByMultipleClaims := squareInchesUsedByMultipleClaims(fabric)
	fmt.Println("DAY 3, PART 1: ", squareInchesUsedByMultipleClaims)
	partTwoId := idOfNonOverlappingClaim(claims, fabric)
	fmt.Println("DAY 3, PART 2: ", partTwoId)
	return common.Result{squareInchesUsedByMultipleClaims, partTwoId}
}

func markClaimsOnFabric(claims []Claim) map[string]int {
	var fabric = make(map[string]int)
	for _, claim := range claims {
		for x := 0; x < claim.width; x++ {
			for y := 0; y < claim.height; y++ {
				markSquareOnFabric(claim, x, y, fabric)
			}
		}
	}
	return fabric
}

func markSquareOnFabric(claim Claim, x int, y int, fabric map[string]int) {
	currentX := claim.startX + x
	currentY := claim.startY + y
	fabric[key(currentX, currentY)]++
}

func squareInchesUsedByMultipleClaims(fabric map[string]int) (sum int) {
	for _, v := range fabric {
		if v > 1 {
			sum++
		}
	}
	return
}

func idOfNonOverlappingClaim(claims []Claim, fabric map[string]int) string {
	for _, claim := range claims {
		if !isOverLapping(claim, fabric) {
			return claim.id
		}
	}
	return "ERROR"
}

func isOverLapping(claim Claim, fabric map[string]int) bool {
	for x := 0; x < claim.width; x++ {
		for y := 0; y < claim.height; y++ {
			currentX := claim.startX + x
			currentY := claim.startY + y
			if fabric[key(currentX, currentY)] > 1 {
				return true
			}
		}
	}
	return false
}

func key(x int, y int) string {
	return strconv.Itoa(x) + "+" + strconv.Itoa(y)
}

