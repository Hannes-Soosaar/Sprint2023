package sprint

func Atoi(inputString string) int {
    answer := 0
    isNegative := false

	//checks to see if the  first and second positions do not contain the same operator char

    if inputString[0] == '+' && inputString[0] == inputString[1] || inputString[0] == '-' && inputString[0] == inputString[1] {
        return 0
    }

    for _, digit := range inputString {
        switch {
        case digit == '+': // if the code has a + it will contine to the start 
            continue
        case digit == '-': // if the code has a - it will continue to the start 
            isNegative = true	// changes the local variable to true
            continue
        case digit == ' ': // if there is  a space retruns the reuslt as
            return 0
        case digit >= '0' && digit <= '9': // if there is a number it will get added to the aswer as an integer
             answer = answer*10 + int(digit-'0')
        }
    }

    if isNegative {  // if there was one negative value it will return the code as negative
        answer = -answer
    }

    return answer 
}