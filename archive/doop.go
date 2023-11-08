package sprint

func DetermineOperation(opeator string) int {

	if opeator == "+" {
		return 1
	} else if opeator == "-" {
		return 2
	} else if opeator == "/" {
		return 3
	} else if opeator == "*" {
		return 4
	} else if opeator == "%" {
		return 5
	} else {
		return 6
	}

}

func Add(firstNumber, secondNumber int) int {
	return firstNumber + secondNumber
}

func Subtract(firstNumber, secondNumber int) int {
	return firstNumber - secondNumber
}

func Multiply(firstNumber, secondNumer int) int {
	return firstNumber * secondNumer
}

func Divide(divided, divisior int) int {
	if divisior == 0 {
		return 0
	}
	return divided / divisior
}

func Modulo(divided, divisior int) int {
	return divided % divisior
}

func Doop(a int, op string, b int) int {

	switch DetermineOperation(op) {

	case 1:
		return Add(a, b)
	case 2:
		return Subtract(a, b)
	case 3:
		return Divide(a, b)
	case 4:
		return Multiply(a, b)
	case 5:
		return Modulo(a, b)
	case 6:
		return 0
	}
return 0

}
