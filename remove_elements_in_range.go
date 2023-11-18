package sprint

func sortAscending(firstNum, secondNum int) (int, int) {	
	if firstNum > secondNum {
		return firstNum, secondNum
	}
	return secondNum, firstNum
} // sorts the numbers to be in the right order

func clampToBounds( array []float64, num int) int {
	bigestIndexInArray  := len(array)
  if num <0 {
	num = 0
  } else if num > bigestIndexInArray{ 
	return bigestIndexInArray
  }
return num
}
func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	biggerNum, smallerNum := sortAscending(from, to)

		biggerNum = clampToBounds(arr,biggerNum)
		smallerNum = clampToBounds(arr,smallerNum)

	return append(arr[:smallerNum], arr[biggerNum:]...)
}
