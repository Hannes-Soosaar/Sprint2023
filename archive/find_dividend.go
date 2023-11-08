package sprint

// func divideByZero(divisor int) bool {

// 	if divisor == 0 {
// 		return true
// 	}
// 	return false
// }


func FindDividend(from, to, divisor int) int {

// if divideByZero(divisor){
// return 0
// }

for i := from ; i < to ; i++ {
	if i%divisor == 0 {
		return i
	}

}
return -1
}