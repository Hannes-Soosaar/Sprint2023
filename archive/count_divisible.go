package sprint

func divideByZero(divisor int) bool {

	if divisor == 0 {
		return true
	}
	return false
}

func validStep(step int) bool {

	if step <= 0 {
		return false
	}
	return true
}

func DivisibleCount(from, to, step, divisor int) int{
 var result int = 0

	for i := from ; i < to; i += step {
	
		if i % divisor == 0 {
			result ++
		}
	}
return result
}

func CountDivisible(from, to, step, divisor int) int {

	if validStep(step) && !divideByZero(divisor) {
		return DivisibleCount(from,to,step,divisor)
	} else {
		return 0
	}

}
