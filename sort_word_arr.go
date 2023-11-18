package sprint

func SortWordArr(a []string) []string {

	tableLenght := len(a)

	swapedAnElement := true

	if tableLenght == 1 {
		return a
	} else if tableLenght == 0 {
		return a
	}
	for swapedAnElement {
		swapedAnElement = false
		for i := 0; i < tableLenght-1; i++ { // iterates over the whole lenght of the table Slice members
			for j := 0; j < tableLenght-1-i; j++ { // iterates over the whole lenght of the table if the last loop was completed than the next loop will be that much shorter
				if a[j] > a[j+1] { // the value at index J is larger than the value at J + 1
					a[j], a[j+1] = a[j+1], a[j] // if we wanted to do this with two line we would need to use a temp variable
					swapedAnElement = true
				}
			}
		}
	}
	return a

}
