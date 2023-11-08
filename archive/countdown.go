package sprint

func Countdown(n int) string {
	result := ""
	for i := n; i > 0; i -= 2 {
		numberAsALetter := rune('0' + i)
		result += string(numberAsALetter) + ", "
	}
	result += "0!"
	return result
}
