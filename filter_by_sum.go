package sprint

func FilterBySum(arr [][]int, limit int) [][]int { // Arr [i][j]int

	var result [][]int // this will return a null value 

	for _, iArr := range arr { // iterates over the first array with the index I that cointains smaller Arrays
		sum := 0 // initializes and sets to zero the sum of the jArr
			for _, jArr := range iArr { // iterates over the small array with index J that contains the values
				sum += jArr // J
			}
		if sum >= limit {
			result = append(result, iArr) // Adds the IArr array to the result as a subarray
		}
	}

	if len(result) == 0 {
		return [][]int{}  // this will return an empty 2D Slice 
	}

	return result
}
