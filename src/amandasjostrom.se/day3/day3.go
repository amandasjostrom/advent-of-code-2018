package day3

import (
	"regexp"
	"amandasjostrom.se/common"
	"strings"
	"fmt"
	"strconv"
)

func Run() {
	claims := claims(common.GetInput(3, "\n"))
	fabric := markOnFabric(claims)
	squareInchesUsedByMultipleClaims := squareInchesUsedByMultipleClaims(fabric)
	fmt.Println("DAY 3, PART 1: ", squareInchesUsedByMultipleClaims)

}
func squareInchesUsedByMultipleClaims(fabric map[string]int) (sum int) {
	for _, v := range fabric {
		if v > 1 {
			sum++
		}
	}
	return
}
func markOnFabric(claims []Claim) map[string]int {
	var fabric = make(map[string]int)
	for _, claim := range claims {
		for x := 0; x < claim.width; x++ {
			for y := 0; y < claim.height; y++ {
				currentX := claim.startX + x
				currentY := claim.startY + y
				key := strconv.Itoa(currentX) + "+" + strconv.Itoa(currentY)
				fabric[key]++
			}
		}
	}
	fmt.Println(fabric)
	return fabric
}
func claims(input []string) []Claim {
	claims := make([]Claim, 0)
	for _, value := range input {
		remove := regexp.MustCompile("[ ]")
		cleaned := remove.ReplaceAllString(value, "")

		separator := regexp.MustCompile("[@:x]")
		result := separator.ReplaceAllString(cleaned, ",")
		claimInput := strings.Split(result, ",")
		claim := Claim{id: claimInput[0],
			startX: common.ToInt(claimInput[1]),
			startY: common.ToInt(claimInput[2]),
			width: common.ToInt(claimInput[3]),
			height: common.ToInt(claimInput[4]),
		}
		claims = append(claims, claim)
	}
	return claims
}

type Claim struct {
	id     string
	startX int
	startY int
	width  int
	height int
}
