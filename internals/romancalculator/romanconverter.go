package romancalculator

func FromInt(number int) string {
	conversions := []struct {
		value   int
		numeral string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman string
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.numeral
			number -= conversion.value
		}
	}
	return roman
}
