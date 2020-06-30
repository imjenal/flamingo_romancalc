package internals

func FizzBuzz(n int) (string,bool) {
	if isMultipleOf(n, 15) {
		return "FizzBuzz", true
	}

	if isMultipleOf(n, 3) {
		return "Fizz", false
	}
	if isMultipleOf(n, 5) {
		return "Buzz", false
	}

	return "", false
}

func isMultipleOf(n, divisor int) bool {
	return n % divisor == 0
}

