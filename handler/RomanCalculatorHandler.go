package handler

import (
	"deeptrace/internals/romancalculator"
	"deeptrace/util"
	"net/http"
	"strconv"
	"strings"
)

func RomanCalculatorHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer util.PanicHandler(writer, request)
		defer request.Body.Close()
		writer.Header().Set("Content-Type", "application/json")

		val := request.FormValue(strings.ToLower("value"))
		if len(val) > 0  {
			value, _ := strconv.Atoi(val)
			result := romancalculator.FromInt(value)
				util.HandleSuccess(writer, []byte(result))
			} else {
			util.HandleError(writer, http.StatusBadRequest, `{"Please specify a value"}`)
		}
	}
}
