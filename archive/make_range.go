package sprint

func MakeRange(min, max int) []int {

	var result []int // creates a slice --> a dynamic array

	if min > max {
		return result
	}
	for i := min; i < max; i++ {
		result = append(result, i) // adds i to the slice result essentially (result =+ i) wehere the result is an array
	}
	return result
}
