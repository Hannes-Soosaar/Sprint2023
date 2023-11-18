package sprint

func CountIf(f func(string) bool, tab []string) int {
	var count int

	for _, s := range tab {
		if f(s) {
			count++
		}
	}
	return count
}

func IsUpper(s string) bool {
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}
func IsLower(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
func IsAlphanumeric(s string) bool {
	for _, r := range s {
		if !(r >= '0' && r <= '9' || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return true

}
