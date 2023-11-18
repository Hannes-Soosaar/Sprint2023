package sprint

func StrRev(s string) string {
	resultString := []rune{}
	runes := []rune(s)
	for index := range runes {
		resultString = append(resultString, runes[len(runes)-1-index]) // len(runes) returns the count of runes  - 1 returns the index equivalent
	}
	return string(resultString)
}
