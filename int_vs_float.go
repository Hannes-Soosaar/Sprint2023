package sprint

func IntVsFloat(integerNumber int, floatNumber float32) string {

	floatAsInteger := int(floatNumber)
	var floatFraction float32 = floatNumber - float32(floatAsInteger)
	
	if (integerNumber == floatAsInteger && floatFraction == 0){
		return "Same"
	} else if (integerNumber < floatAsInteger || floatFraction >= 0 && integerNumber == floatAsInteger) {
		return "Float"
	} else {
		return "Integer"
	}

}