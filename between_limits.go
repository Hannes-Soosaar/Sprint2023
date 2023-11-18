package sprint

func DetermineFirst(first, second rune) rune {
	if first < second {
		return first + 1
	} else {
		return second + 1
	}
}

func DetermineSecond(first, second rune) rune {
	if first < second {
		return second
	} else {
		return first
	}
}

func BetweenLimits(first, second rune) string {

	var output string = ""
	for i := DetermineFirst(first, second); i < DetermineSecond(first, second); i++ {
		output = output + string(i)
	}
return output
}
