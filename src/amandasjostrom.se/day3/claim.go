package day3

import (
	"regexp"
	"amandasjostrom.se/common"
)

type Claim struct {
	id     string
	startX int
	startY int
	width  int
	height int
}

func claims(input []string) []Claim {
	claims := make([]Claim, 0)
	for _, value := range input {
		claimInput := regexp.MustCompile("\\d+").FindAllString(value, -1)
		claim := Claim{
			id: "#" + claimInput[0],
			startX: common.ToInt(claimInput[1]),
			startY: common.ToInt(claimInput[2]),
			width: common.ToInt(claimInput[3]),
			height: common.ToInt(claimInput[4]),
		}
		claims = append(claims, claim)
	}
	return claims
}
