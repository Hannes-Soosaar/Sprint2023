package sprint

import "fmt"

func CombN(n int) []string {
	var result []string								// declares a nil result  string Slice  
	var backtrack func(path string, start int)		//  creates a function variable named backtrac
	backtrack = func(path string, start int) {		//  adds a anonymous fucntion to the function variable
		if len(path) == n {							// ends the loop if the  lenght of the paht is the same as the input variabl
			result = append(result, path)			// adds the path string to the nill valued 	
			return
		}
		for i := start; i <= 9; i++ {
			backtrack(path+fmt.Sprintf("%d", i), i+1) // runs the anonymous funcstion passing in the value of i+1 to be saved 
		}
	}
	backtrack("", 0)								// runs the backtrack function with 
	return result

}
