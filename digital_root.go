package sprint

func DigitalRoot(n int) int {
	sum := 0

	if n < 10 {
		return n
	}
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return DigitalRoot(sum)
}
