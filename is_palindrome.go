package sprint

func IsPalindrome(s string) bool {
	if s == "" {
		return true
	}
	sLenghts := len([]rune(s)) - 1
	compareString := ""
	for i := 0; i <= sLenghts; i++ {
		compareString += string(s[sLenghts-i])
	}
	if compareString == s {
		return true
	}
	return false
}
