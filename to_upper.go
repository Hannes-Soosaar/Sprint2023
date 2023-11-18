package sprint

func ToUpper(s string) string {
	resultSrting := []rune{} // string consisting of
	for _, r := range s {
		if r >= 'a' &&  r <= 'z' {
			r = r - 32 // could also use the a-A to get the distance
		}
		resultSrting = append(resultSrting, r)
	}
	return string(resultSrting)
}
