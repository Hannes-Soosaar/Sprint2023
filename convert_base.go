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
func containsForbidenRunesAndDuplicates(myString string) bool {
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
	if containsForbidenRunesAndDuplicates(base) || LengthOf(base) < 2 || !ContainsTwoOrMoreUnique(base) {
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
func AtoiBase(s string, myString string) int {

	if containsForbidenRunes(myString) || !ContainsTwoOrMoreUnique(myString) {
	return 0
	}
	
	if s[0]  == '-' {
		return 0
	}
	
	baseMap := make(map[rune]int)
	for i, r := range myString {
		baseMap[r] = i
	}
	baseLen := LengthOf(myString)
	result := 0
	for _, r := range s {
		val, ok := baseMap[r]
		if !ok {
			return 0
		}
		result = result*baseLen + val
	}
	
	return result 
	}

	func ConvertBase(nbr, baseFrom, baseTo string) string {
		return NbrBase(AtoiBase(nbr,baseFrom),baseTo)
	}