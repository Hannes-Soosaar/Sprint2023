package sprint

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
