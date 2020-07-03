package pinkflamingo

func IsPinkFlamingo(n int) bool {
	_, isMultipleOf15 := FizzBuzz(n)
	return IsFibonacci(n) && isMultipleOf15
}
