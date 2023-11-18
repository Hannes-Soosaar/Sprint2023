package sprint

func TransposeMatrix(matrix [][]int) [][]int {

	// create empty matrix where the rows equla col and cols equal rows
	rows := len(matrix)               // get the primary slice lenght
	cols := len(matrix[0])            // get the secondary matrix lenght at the primary matrix pos 0
	transposed := make([][]int, cols) // create an empty primary slice where the number of cols now is the numebr of rows
	for i := range transposed {       // create an empty seondary slice where the number of rowes now is the numebr of cols
		transposed[i] = make([]int, rows) //
	}
	// copy one matrix to the other with a for loop
	for i := 0; i < rows; i++ { // moves through each original matix row
		for j := 0; j < cols; j++ { // moves through each original matrix column
			transposed[j][i] = matrix[i][j] // swaps the values from row to column
		}
	}
	return transposed
}
