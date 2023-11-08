package sprint

func Join(strs []string, sep string) string {
	resultSrtings := []string{}
	var resultString string = ""
	for i := 0; i < len(strs)-1; i++ { // creates a slice of string adding the sep
		resultSrtings = append(resultSrtings, strs[i]) // make sures it adds the old to the new 
		resultSrtings = append(resultSrtings, sep) 		//	it stop before the very last member 	
	}
	resultSrtings = append(resultSrtings,strs[len(strs)-1]) // manually added the last member

	for i := 0; i < len(resultSrtings); i++ { // itterates through the result slice and concatinates it to a single string
		resultString += resultSrtings[i]
	}
	return resultString
	
}
