package handler

import (
	"deeptrace/internals/romancalculator"
	"deeptrace/util"
	"encoding/json"
	"net/http"
)

type Expression struct {
	Expr string `json:"expr"`
}

func RomanCalculatorHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer util.PanicHandler(writer, request)
		writer.Header().Set("Content-Type", "application/json")
		requestValue, invalidRequestBody := decodeRequestBody(request)
		if invalidRequestBody != nil {
			util.Response(writer, http.StatusBadRequest, "Invalid request body")
		} else {
			romancalculator.RomanCalculator{}.Execute(writer, requestValue.Expr, "")
		}
		defer request.Body.Close()
	}
}

func decodeRequestBody(request *http.Request) (Expression, error) {
	var requestValue Expression
	decoder := json.NewDecoder(request.Body)
	invalidRequestBody := decoder.Decode(&requestValue)
	return requestValue, invalidRequestBody

}
