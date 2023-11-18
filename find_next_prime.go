package sprint

func FindNextPrime(n int) int {
	isPrime := true
	if n <= 1 { // any number smaller than 1 ( included ) is not a prime number
		isPrime = false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}
	for !isPrime {
		return FindNextPrime(n + 1)
	}
	return n
}
