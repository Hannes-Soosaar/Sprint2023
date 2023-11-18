package sprint

func LengthOf(myString string) int {
	return len([]rune(myString))
}

func Split(myString, separator string) []string {

	var result []string = []string{} // establishes start of the result
	start := 0                       // keep track of the start of the sting section to be checking

	for i := 0; i < LengthOf(myString); i++ { // creates a loop taht iterate over every character in the input string
		if i+LengthOf(separator) <= LengthOf(myString) { // makes sure that the separator is largher than the input string
			if myString[i:i+LengthOf(separator)] == separator { // mySting[i:i] check the string between the two indexex
				result = append(result, string([]rune(myString)[start:i]))
				start = i + LengthOf(separator) // moves the start to the new postions that is sep distance away from the star
				i = start - 1                   // if a match is found move the start back 1 place
			}
		}
	}
	if start < LengthOf(myString) { 
		result = append(result, string([]rune(myString)[start:]))
	}
	return result
}
