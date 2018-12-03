package day3

import (
	"regexp"
	"strings"
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
