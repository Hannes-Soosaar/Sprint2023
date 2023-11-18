package sprint

func LengthOf(myString string) int {
	return len([]rune(myString))
}

func containsForbidenRunes(myString string) bool {
	for i, r := range myString {
		if r == '+' || r == '-' {
			return true
		}
		for j := i + 1; j < LengthOf(myString); j++ { // checks to see if  there are any duplicate runes

			if myString[j] == myString[i] {
				return true
			}
		}
	}
	return false
}

func ContainsTwoOrMoreUnique(myString string) bool {
	for i, r := range myString {
		if rune(myString[0]) != r && i > 0 {
			return true
		}
	}

	return false
}

func NbrBase(n int, base string) string {
	if containsForbidenRunes(base) || LengthOf(base) < 2 || !ContainsTwoOrMoreUnique(base) {
		return "NV"
	}
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	baseLen := LengthOf(base)
	result := ""

	for n > 0 {
		result = string(base[n%baseLen]) + result
		n /= baseLen
	}

	if result == "" {
		return "0"
	}

	if negative {
		result = "-" + result
	} 
	return result
}
