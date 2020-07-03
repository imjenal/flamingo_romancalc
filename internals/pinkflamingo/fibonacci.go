package pinkflamingo

import (
	"deeptrace/util"
	"math"
)

func IsFibonacci(n int) bool {
	return isPerfectSquare(5*n*n+4) || isPerfectSquare(5*n*n-4)
}

func isPerfectSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return util.IsEquals(sqrt * sqrt, n)
}
