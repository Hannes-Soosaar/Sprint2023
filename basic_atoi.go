package sprint

func BasicAtoi(inputString string) int {

	tempNumber := 0
	answer := 0
	for _, r := range inputString {

		
		singleRune := rune(r)
			
		if singleRune < '0' || singleRune > '9'{
				return 0
			}

		tempNumber = int(singleRune - '0')

		if tempNumber > 0 &&  tempNumber <9 {
			answer = answer*10 + tempNumber // it move the already inputed number one place up 
		}
	}
	return answer
}
