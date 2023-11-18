package sprint

import "fmt"


// Rather hard to grasp this logic 

func Combinations() string {

	answer := ""
	for firstNumber := 0; firstNumber <= 7; firstNumber++ {
		for secondNumber := firstNumber + 1; secondNumber <= 8; secondNumber++ {
			for thirdNumber := secondNumber + 1; thirdNumber <= 9; thirdNumber++ {
				numbers := fmt.Sprintf("%d%d%d, ", firstNumber, secondNumber, thirdNumber)
				answer += numbers
			}
		}
	}
	return answer[:len(answer)-2]  // this removes the last two characters from the string
}
