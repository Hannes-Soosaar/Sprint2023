package sprint

const (
	EQUAL           = 0
	ACENDING_ORDER  = -1
	DECENDING_ORDER = 1
)

//You're tasked with creating a function called IsSorted that returns true if the slice of string is sorted according to a provided comparison function, and false otherwise

func IsSorted(f func(a, b string) int, arr []string) bool {
	isAscending := true
	isDecending := true

	for i := 0; i < len(arr)-1; i++ { //stops the count one before the end of the slice so it does not request an outofbounds member
		compareValue := f(arr[i], arr[i+1])
		if compareValue != ACENDING_ORDER && compareValue != EQUAL {
			isAscending = false
		} else if compareValue != DECENDING_ORDER && compareValue != EQUAL {
			isDecending = false
		}
	}
	return isAscending || isDecending
}
func Compare(firstS, secondS string) int {

	if firstS == secondS {
		return EQUAL
	} else if firstS < secondS {
		return ACENDING_ORDER
	} else {
		return DECENDING_ORDER
	}
}
