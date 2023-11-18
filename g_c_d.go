package sprint

func IsCommonDivisor(a, b, i int) bool {
	return a%i == 0 && b%i == 0
}
func GCD(a, b int) int {

	if a == 0 {
		return b
	} else if b == 0 {
		return a
	}

	largestCommonInteger := 1
	if a < b {
		largestCommonInteger = a
	} else {
		largestCommonInteger = b
	}
	largestDivider := 1

	for i := 1; i <= largestCommonInteger; i++ { // starts with 1 as you can not divide by zero
		if IsCommonDivisor(a, b, i) {
			largestDivider = i
		}
	}
	return largestDivider
}
