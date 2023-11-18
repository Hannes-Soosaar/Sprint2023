package sprint

const (
	EQUAL           = 0
	FIRST_IS_FIRST  = -1
	SECOND_IS_FIRST = 1
)

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {

	numberOfElementsInSlice := len(a)

	swapedAnElement := true

	if numberOfElementsInSlice == 1 {
		return a
	} else if numberOfElementsInSlice == 0 {

		return []string{}
	}

	for swapedAnElement {
		swapedAnElement = false
		for i := 0; i < numberOfElementsInSlice-1; i++ { // iterates over the whole lenght of the table Slice members
			for j := 0; j < numberOfElementsInSlice-1-i; j++ { // iterates over the whole lenght of the table if the last loop was completed than the next loop will be that much shorter
				if f(a[j], a[j+1]) == 1 { // the value at index J is larger than the value at J + 1
					a[j], a[j+1] = a[j+1], a[j] // if we wanted to do this with two line we would need to use a temp variable
					swapedAnElement = true
				}
			}
		}
	}
	return a
}

func Compare(firstS, secondS string) int {

	if firstS == secondS {
		return EQUAL
	} else if firstS < secondS {
		return FIRST_IS_FIRST
	} else {
		return SECOND_IS_FIRST
	}
}
