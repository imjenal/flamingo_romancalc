package pinkflamingo

func FizzBuzz(number int) (string, bool) {
	if isMultipleOf(number, 15) {
		return "FizzBuzz", true
	}

	if isMultipleOf(number, 3) {
		return "Fizz", false
	}
	if isMultipleOf(number, 5) {
		return "Buzz", false
	}

	return "", false
}

func isMultipleOf(number, divisor int) bool {
	return number%divisor == 0
}
