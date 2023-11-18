package sprint

func RecursiveFactorial(n int) int { // starts from the value n multiplies and add untill reaching 0
	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	}	
	maxInt:= 1<<31 - 1 // in number terms it is 2147483647 however this is a clever way to check what the max number is we can do with binary number is base 10 
		if n * RecursiveFactorial(n-1)> maxInt{
			return 0
		}
	return  n * RecursiveFactorial(n-1)
}
