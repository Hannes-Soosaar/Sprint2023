package sprint

var romanNumeralMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func FromRoman(s string) int {

	result := 0
	previousValu := 0

	for i := len(s) - 1; i >= 0; i-- {
		value := romanNumeralMap[s[i]] //
		if value < previousValu {
			result -= value
		} else {
			result += value
		}
		previousValu = value
	}
	return result
}
