package sprint

import (
	"fmt"
)

func StrCompress(input string) string {
	if input == "" {
		return input
	}
	compress := 0
	compresseString := ""
	canCompress := false
	for i := 0; i < len(input)-1; i++ {
		if rune(input[i]) == rune(input[i+1]) {
			compress++
			canCompress = true
		} else {
			if canCompress {
				compresseString += fmt.Sprintf("%d%s", compress+1, string(input[i]))
				compress = 0
				canCompress = false
			} else {
				compresseString += fmt.Sprintf(string(input[i]))
			}
		}
	}
	if canCompress {
		compresseString += fmt.Sprintf("%d%s", compress+1, string(rune(input[len(input)-1])))
	} else {
		compresseString += fmt.Sprintf("%s", string(rune(input[len(input)-1])))
	}
	return compresseString
}
