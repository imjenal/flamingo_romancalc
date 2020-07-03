package romancalculator

import (
	"deeptrace/util"
	"net/http"
)

type RomanCalculator struct {
}

func (r RomanCalculator) Execute(writer http.ResponseWriter, expr, temp string) {
	if util.IsValidData(expr) {
		romanNumber := NewRoman()
		numberEquivalentExpression := numberEquivalentExpression(expr, romanNumber)
		evaluatedValue, err := EvaluateinBodmas(numberEquivalentExpression)
		if err != nil {
			util.Response(writer, http.StatusBadRequest, `{"Evaluation error"}`)
		}
		evaluateRomanValue(writer, evaluatedValue, romanNumber)
	} else {
		util.Response(writer, http.StatusBadRequest, `{"Please specify a value"}`)
	}
}

func numberEquivalentExpression(expr string, romanNumber *Roman) string {
	var numberEquivalentExpression string
	expressionComponents := util.SplitBySeparator(expr, " ")
	for _, expressionComponent := range expressionComponents {
		if !IsOperator(expressionComponent) {
			numberEquivalent := romanNumber.ToNumber(expressionComponent)
			numberEquivalentExpression += util.ToString(numberEquivalent)
		} else {
			numberEquivalentExpression += expressionComponent
		}
	}
	return numberEquivalentExpression
}

func evaluateRomanValue(writer http.ResponseWriter, evaluatedValue interface{}, romanNumber *Roman) {
	evaluatedIntegerValue := int(evaluatedValue.(float64))
	if evaluatedIntegerValue > 0 {
		romanEvaluatedValue := romanNumber.ToRoman(evaluatedIntegerValue)
		util.ResponseSuccess(writer, []byte(romanEvaluatedValue))
	} else {
		util.ResponseSuccess(writer, []byte(`{"Unable to perform calculation""}`))
	}
}
