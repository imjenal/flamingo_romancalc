package pinkflamingo

func IsPinkFlmaingo(n int) bool {
	_, isMultipleOf15 := FizzBuzz(n)
	return IsFibonacci(n) && isMultipleOf15
}
