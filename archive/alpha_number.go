package sprint

func AlphaNumber(n int) string {
	var anwser string = ""
	var tempAnswer string = ""
	var addMinus bool = false

	// looks to see if n is negative, preserves the minus and makes n positive

	if n < 0 {
		addMinus = true
		n = n * -1
	} else if n == 0{
		return "a"
	}

	for n > 0 {
		temp := n % 10                        // extracts last digit
		tempAnswer = string(rune('a' + temp)) // saves the digit to the string
		anwser = tempAnswer + anwser
		n = n / 10 // "pop" the last digit
	}

	if addMinus {
		anwser = "-" + anwser
	}
	return anwser
}
