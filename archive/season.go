package sprint

func Season(month string) string {

	var seasonCase int = 0;
	var invalidChoice string = ""

	if (month == "jan" || month == "feb" || month == "dec"){
		seasonCase = 1
	}else if (month == "mar" || month == "apr" || month == "may"){
		seasonCase =2
	}else if (month == "jun" || month == "jul" || month == "aug"){
		seasonCase = 3
	}else if (month == "sep" || month == "oct" || month == "nov"){
		seasonCase = 4
		}else{
			seasonCase =5
			invalidChoice = "invalid input: " + month
		}
	

	switch seasonCase {
	case 1: 
		return "winter" 
	case 2:
		return "spring"
	case 3:
		return "summer"
	case 4:
		return "autumn"
	case 5:
	return invalidChoice	
	default:
		return "notdefined"
	}	


}