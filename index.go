package sprint

func LengthOf(myString string) int {
	return len([]rune(myString))
}

func Index(myString string, toFind string) int {


	if myString == toFind {
		return 0
	}

	sLen, toFindLen := LengthOf(myString), LengthOf(toFind)

	for i := 0; i <= sLen-toFindLen; i++ {
		if myString[i:i+toFindLen] == toFind {
			return i
		}
	}
	return -1
}
