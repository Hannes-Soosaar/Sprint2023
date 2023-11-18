package sprint

func Overlap(arr1, arr2 []int) []int {

	frequencyMap := make(map[int]int)
	for _, num := range arr1 {
		frequencyMap[num]++
	}
	var result []int
	for _, num := range arr2 {
		if frequencyMap[num] > 0 {
			result = append(result, num)
			frequencyMap[num]--
		}
	}
	if len(result) == 0 {
		return []int{}
	}

	return SortIntegerTable(result)
}

func SortIntegerTable(table []int) []int {
	for i := 0; i < len(table); i++ {
		min := i
		for j := i + 1; j < len(table); j++ {
			if table[j] < table[min] {
				min = j
			}
		}
		if min != i {
			table[i], table[min] = table[min], table[i]
		}
	}
	return table
}
