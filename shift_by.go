package sprint

func ShiftBy(lowerCaseLetter rune, step int) rune {

	newLowerCaseLetter := lowerCaseLetter + rune(step%26)
	if newLowerCaseLetter > 'z' {
		newLowerCaseLetter = 'a' + (newLowerCaseLetter - 'z' - 1)
	} else if newLowerCaseLetter < 'a' {
		newLowerCaseLetter = 'z' - ('a' - newLowerCaseLetter - 1)
	}
	return newLowerCaseLetter
}
