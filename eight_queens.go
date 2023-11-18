package sprint

import "strconv"
import "strings"

func SolveNQueens(n int) string {
	var result strings.Builder  // establishes that the result will be a strings.Builder data type
	quens := make([]int, n)     // creates an empty slice
	Solve(quens, 0, n, &result) // passes in the address to result
	return result.String()      // convets the result to a string
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func QueensToString(queens []int) string {
	var sb strings.Builder
	for _, queen := range queens {
		sb.WriteString(strconv.Itoa(queen + 1))
	}
	return sb.String()
}
func isValid(queens []int, row int) bool {
	for i := 0; i < row; i++ {
		if queens[i] == queens[row] || Abs(queens[i]-queens[row]) == row-i {
			return false
		}
	}
	return true
}

func Solve(queens []int, row, n int, result *strings.Builder) {
	if row == n {
		result.WriteString(QueensToString(queens) + "\n") // creates a string from all the queens
		return
	}
	for i := 0; i < n; i++ {
		queens[row] = i
		if isValid(queens, row) {
			Solve(queens, row+1, n, result)
		}
	}
}
func EightQueens() string {
	n := 8
	s := SolveNQueens(n)
	return s[:len(s)-1]
}
