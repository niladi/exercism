package scrabble

import (
	"regexp"
	"strings"
)

func Score(input string) int {
	scrabble := map[int]*regexp.Regexp{
		1:  regexp.MustCompile("[AEIOULNRST]"),
		2:  regexp.MustCompile("[DG]"),
		3:  regexp.MustCompile("[BCMP]"),
		4:  regexp.MustCompile("[FHVWY]"),
		5:  regexp.MustCompile("[K]"),
		8:  regexp.MustCompile("[JX]"),
		10: regexp.MustCompile("[QZ]"),
	}
	input = strings.ToUpper(input)
	sum := 0
	for k, v := range scrabble {
		sum += len(v.FindAll([]byte(input), -1)) * k
	}
	return sum
}
