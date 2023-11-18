package sprint
// import (
// )

func LengthOf(myString string) int {
	return len([]rune(myString))
}
func ContainsTwoOrMoreUnique(myString string) bool {
	for i, r := range myString {
		if rune(myString[0]) != r && i > 0 {
			return true
		}
	}
	return false
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
func AtoiBase(s string, myString string) int {

if containsForbidenRunes(myString) || !ContainsTwoOrMoreUnique(myString) {
return 0
}

if s[0]  == '-' {
	return 0
}

baseMap := make(map[rune]int)  // creates an empty map 
for i, r := range myString {  // save the runes in map that can be accesed with the index 1
	baseMap[r] = i  // creat
}

baseLen := LengthOf(myString)
result := 0
for _, r := range s {
	val, ok := baseMap[r]  // checks to see if "r" rune is prent OK is a bool variable and val sotores the "r" values
	if !ok {
		return 0
	}
	result = result*baseLen + val
}

return result 
}
