package sprint

func Accumulate(n int) int {
	
	var result int = 0;

	for  i := 1  ; i <= n ; i++ {

	result = result + i	
	}
	return result
}

