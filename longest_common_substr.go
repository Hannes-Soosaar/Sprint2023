package sprint

func LongestCommonSubstr(str1, str2 string) string {

	if len(str1) == 0 || len(str2) == 0 {
		return "" // Handle empty strings
	}

	len1, len2 := len(str1), len(str2) //get the lenght of the strings, positions
	matrix := make([][]int, len1+1)    // makes a 2d slice on indexes

	for i := range matrix {
		matrix[i] = make([]int, len2+1) // makes the sub matrix holding one more poiton than len2
	}
	var maxLength, endIndex int // marks the index on str1
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if str1[i-1] == str2[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1

				if matrix[i][j] > maxLength {
					maxLength = matrix[i][j]
					endIndex = i - 1
				}
			}
		}
	}
	if maxLength == 0 {
		return "" // No common substring found
	}

	startIndex := endIndex - maxLength + 1
	return str1[startIndex : endIndex+1]
}
