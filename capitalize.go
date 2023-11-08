package sprint

func Capitalize(s string) string {
	resultString := []rune{}
	capitalizeNext := false // logic to see if the next letter should be capital or not
	for index, r := range s {
		if index == 0 && r >= 'a' && r <= 'z' { // if the letter is the first in the string check if small if so make big
			r -= 'a' - 'A'
		}
		if capitalizeNext && r >= 'a' && r <= 'z' { // if the letter need to be capitalized makes it a capital
			r -= 'a' - 'A'
			capitalizeNext = false
		} else if !capitalizeNext && r >= 'A' && r <= 'Z' && index != 0 { //   if the letter is a capital letter and should not be makes it lowercase
			r += 'a' - 'A'
		}
		resultString = append(resultString, r) // adds the rune to the slice of runes

		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' { // checks if the letter is a non letter
			capitalizeNext = false
		} else {
			capitalizeNext = true
		}
	}
	return string(resultString)
}