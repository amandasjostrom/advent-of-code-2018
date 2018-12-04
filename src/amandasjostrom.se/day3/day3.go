package day3

import (
	"amandasjostrom.se/common"
	"fmt"
	"strconv"
)

var overlappingClaims = map[string]bool{}

type Square struct {
	count int
	firstId string
}

func Run(input []string) common.Result {
	claims := claims(input)
	fabric := markClaimsOnFabric(claims)
	squareInchesUsedByMultipleClaims := squareInchesUsedByMultipleClaims(fabric)
	fmt.Println("DAY 3, PART 1: ", squareInchesUsedByMultipleClaims)
	partTwoId := idOfNonOverlappingClaim(claims)
	fmt.Println("DAY 3, PART 2: ", partTwoId)
	return common.Result{squareInchesUsedByMultipleClaims, partTwoId}
}

func markClaimsOnFabric(claims []Claim) map[string]Square {
	var fabric = make(map[string]Square)
	for _, claim := range claims {
		for x := 0; x < claim.width; x++ {
			for y := 0; y < claim.height; y++ {
				markSquareOnFabric(claim, x, y, fabric)
			}
		}
	}
	return fabric
}

func markSquareOnFabric(claim Claim, x int, y int, fabric map[string]Square) {
	currentX := claim.startX + x
	currentY := claim.startY + y
	key := key(currentX, currentY)
	if 	square := fabric[key]; square.firstId != "" {
		overlappingClaims[square.firstId] = true
		overlappingClaims[claim.id] = true
		square.count++
		fabric[key] = square
	} else {
		fabric[key] = Square{1, claim.id}
	}
}

func squareInchesUsedByMultipleClaims(fabric map[string]Square) (sum int) {
	for _, v := range fabric {
		if v.count > 1 {
			sum++
		}
	}
	return
}

func idOfNonOverlappingClaim(claims []Claim) string {
	for _, claim := range claims {
		if !overlappingClaims[claim.id] {
			return claim.id
		}
	}
	return "ERROR"
}

func key(x int, y int) string {
	return strconv.Itoa(x) + "+" + strconv.Itoa(y)
}

