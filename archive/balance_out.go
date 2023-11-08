package sprint


var equalTrueAndFalse bool 

func evaluateBalance(trueCount,falseCount int) bool {

	if trueCount == falseCount {
	   return true
	}
	return false
}

func BalanceOut(arr []bool) []bool {
	var trueCount int
	var falseCount int

	for _, value := range arr {
		if value {
			trueCount++
		} else {
			falseCount++
		}
	}
	
for !evaluateBalance(trueCount,falseCount){

	trueCount = 0
	falseCount = 0
	
	for _, value := range arr {
		if value {
			trueCount++
		} else {
			falseCount++
		}
	}

	if trueCount == falseCount {
		// fmt.Printf("\n\n GOAL REACHED : %v \n\n",arr)
	return arr
	}else if trueCount < falseCount{
		arr = append(arr, true)
		// fmt.Printf("added a  TRUE %v \n", arr)
	} else {
		arr = append(arr, false)
		// fmt.Printf("added a  FALSE %v \n", arr)
	}

}

	// fmt.Printf("\n End of loop array value : %v \n",arr)
		return arr
}
