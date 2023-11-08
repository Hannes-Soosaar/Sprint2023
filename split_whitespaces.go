package sprint

func SplitWhitespaces(s string) []string {
	tempString := ""
	resultStrinSlice := []string{}
	for i, r := range s {

		if r != '\t' && r != '\n' && r !=' '{
			tempString += string(r)
		}
		if r == ' ' || r == '\t' || r == '\n' {
			resultStrinSlice = append(resultStrinSlice, tempString)
			tempString = ""
		} else if i == len([]rune(s))-1 {
			resultStrinSlice = append(resultStrinSlice, tempString)
			tempString = ""
		}
	}
	return resultStrinSlice
}
