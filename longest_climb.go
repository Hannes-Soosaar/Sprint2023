package sprint

func LongestClimb(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	start, maxStart, maxLength := 0, 0, 1

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			if i-start+1 > maxLength {
				maxLength = i - start + 1
				maxStart = start
			}
		} else {
			start = i
		}
	}

	return arr[maxStart : maxStart+maxLength]
}

// func LongestClimb(arr []int) []int {

// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	firstTempSlice := []int{}
// 	secondTempSlice := []int{}

// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] > arr[i-1] {
// 			firstTempSlice = append(firstTempSlice, arr[i])
// 		} else {
// 			if len(firstTempSlice) > len(secondTempSlice) {
// 				secondTempSlice = make([]int, len(firstTempSlice))
// 				copy(secondTempSlice, firstTempSlice)
// 			}
// 			firstTempSlice = []int{arr[i]}
// 		}
// 	}
// 	if len(firstTempSlice) > len(secondTempSlice) {
// 		secondTempSlice = make([]int, len(firstTempSlice))
// 		copy(secondTempSlice, firstTempSlice)
// 	}
// 	return secondTempSlice
// }
