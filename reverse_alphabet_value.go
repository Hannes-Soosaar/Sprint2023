package sprint

func ReverseAlphabetValue(ch rune) rune {

	var placeInTheAlphabetFromTheStart = ('z' - ch)

	return ('a' + placeInTheAlphabetFromTheStart)
}
