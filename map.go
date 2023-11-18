package sprint

func Map(f func(int) bool, a []int) []bool {
	resultSlice := []bool{}
	for _, i := range a {
		if f(i) {
			resultSlice = append(resultSlice, true)
		} else {
			resultSlice = append(resultSlice, false)
		}
	}
	return resultSlice
}
func IsNegative(numberForEvaluation int) bool {
	if numberForEvaluation < 0 {
		return true
	}
	return false
}
func IsPrime(n int) bool {
	if n <= 1 { // any number smaller than 1 ( included ) is not a prime number
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
