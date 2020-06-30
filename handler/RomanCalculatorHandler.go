package handler

import (
	"deeptrace/internals/romancalculator"
	"deeptrace/util"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Expression struct {
	Expr string `json:"expr"`
}

func RomanCalculatorHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer util.PanicHandler(writer, request)
		writer.Header().Set("Content-Type", "application/json")

		var val Expression
		decoder := json.NewDecoder(request.Body)
		invalidRequest := decoder.Decode(&val)

		if invalidRequest != nil {
			util.HandleError(writer, http.StatusBadRequest, "Invalid request body")
		} else {
			if len(val.Expr) > 0 {
				romancnv := romancalculator.NewRoman()
				var expr string
				res1 := strings.Split(val.Expr, " ")
				for _, arr := range res1 {
					if !isOperator(arr) {
						intResult := romancnv.ToNumber(arr)
						expr += strconv.Itoa(intResult)
					} else {
						expr += arr
					}
				}
				result, err := romancalculator.EvaluateinBodmas(expr)
				if err != nil {
					util.HandleError(writer, http.StatusBadRequest, `{"Evaluation error"}`)
				}

				integerResult := int(result.(float64))
				if integerResult > 0 {
					romanResult := romancnv.ToRoman(integerResult)
					util.HandleSuccess(writer, []byte(romanResult))
				} else {
					util.HandleSuccess(writer, []byte("Unable to perform calculation"))
				}

			} else {
				util.HandleError(writer, http.StatusBadRequest, `{"Please specify a value"}`)
			}
		}
		defer request.Body.Close()
	}
}

func isOperator(op string) bool {
	switch op {
	case "(", ")", "*", "/", "+", "-":
		return true
	default:
		return false
	}
}
