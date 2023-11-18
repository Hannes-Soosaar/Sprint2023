package sprint

import (
	"strings"
)

func AreAnagrams(str1, str2 string) bool {

	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))
	if len(str1) != len(str2) {
		return false
	}
	return getRuneSum(str1) == getRuneSum(str2)
}
func getRuneSum(s string) int {
	sum := 0
	for _, r := range s {
		sum += int(r)
	}
	return sum
}
