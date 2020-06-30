package romancalculator

import (
	"github.com/Knetic/govaluate"
)

func EvaluateinBodmas(expr string) (interface{}, error) {
	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return nil, err
	}
	result, err := expression.Evaluate(nil);
	if err != nil {
		return result,err
	}
	return result, nil
}