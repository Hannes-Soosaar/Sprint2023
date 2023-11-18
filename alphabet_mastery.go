package sprint

func AlphabetMastery(n int) string {
	var outputString string = "a"

	if n == 0 {
		return ""
	 }

	for i := 1; i < n; i++ {
		outputString += string('a' + i)
	}

	return outputString
}