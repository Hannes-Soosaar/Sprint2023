package sprint

func IterativeFactorial(n int) int { 
	if n < 0{
		return 0
	}
	var factorial int = 1
	maxInt:= 1<<31 - 1 // in number terms it is 2147483647 however this is a clever way to check what the max number is we can do with binary number is base 10 
	for i:=1; i<=n ; i++ { // i must start at 1 as multply with 0 give always 0 
		if  factorial > maxInt/i {
			return 0
		} 
		factorial *=i
	}
return factorial
}