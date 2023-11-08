package sprint

func isDivisibleByFour(year int) bool {
	if year%4 == 0 {
		return true
	}
	return false
}

func isDivisibleByOnehundred(year int) bool {
	if year%100 == 0 {
		return true
	}
	return false
}

func isDivisibleByFourhundred(year int) bool {
	if year%400 == 0 {
		return true
	}
	return false
}

func IsLeapYear(year int) bool {

	if !isDivisibleByFour(year) {
		return false
	} else if isDivisibleByFour(year) && isDivisibleByOnehundred(year) && !isDivisibleByFourhundred(year) {
		return false
	} else {
		return true
	}

}
