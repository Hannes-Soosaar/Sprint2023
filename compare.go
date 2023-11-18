package sprint

const (
	EQUAL           = 0
	FIRST_IS_FIRST  = -1
	SECOND_IS_FIRST = 1
)

func Compare(firstS, secondS string) int {

	if firstS == secondS {
		return EQUAL
	} else if firstS < secondS {
		return FIRST_IS_FIRST
	} else {
		return SECOND_IS_FIRST
	}
}
