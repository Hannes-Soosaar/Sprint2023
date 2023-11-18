package sprint

func BalancedParentheses(input string) bool {
	buffer := make([]rune, 0)
	parenthesesMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range input {
		switch char {
		case '(', '{', '[':
			buffer = append(buffer, char)
		case ')', '}', ']':
			if len(buffer) == 0 || buffer[len(buffer)-1] != parenthesesMap[char] {
				return false
			}
			buffer = buffer[:len(buffer)-1]
		}
	}
	return len(buffer) == 0
}
