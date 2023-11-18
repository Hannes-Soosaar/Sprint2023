package sprint

import "fmt"

func isLastPair(firstNumber,secondNumber int) bool {
 if firstNumber == 98 && secondNumber == 99{
	return true
 }
 return false
}

func Pairs() string {

	var firstNumber int = 0
	var secondNumber int = 0
	var tempString string = ""

	for firstNumber < 100 {
		secondNumber = firstNumber +1; // there might be an issue here where the numebr will not start.
		for (secondNumber) < 100 {
			  	tempString += fmt.Sprintf("%02d %02d", firstNumber, secondNumber)
				if len(tempString) != 0 && !isLastPair(firstNumber,secondNumber) {
					tempString +=", "
				}
				secondNumber++
		}
		firstNumber++
	}
	return tempString
}