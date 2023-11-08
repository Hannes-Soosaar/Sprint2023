package sprint

func ReverseAlphabet(step int) string {

	var outputString string = ""
	if step <= 0 {
		step = 1
	}
	for i := 0; i < 26; i += step {
		outputString += string('z' - i)
	}

	return outputString
}