package handler

import (
	"deeptrace/internals/pinkflamingo"
	"deeptrace/util"
	"encoding/json"
	"net/http"
)

func PinkFlamingoHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer util.PanicHandler(writer, request)
		defer request.Body.Close()
		writer.Header().Set("Content-Type", "application/json")

		to := request.FormValue(util.ToLowerCase("To"))
		from := request.FormValue(util.ToLowerCase("From"))
		if util.IsValidNumber(to)  && util.IsValidNumber(from) {
			initialValue, err := util.ToInteger(from)
			finalValue, err := util.ToInteger(to)
			if err != nil {
				util.HandleError(writer, http.StatusBadRequest, `{"Unable to Convert String to Integer"}`)
				return
			}
			if initialValue < finalValue {
				result := make([]interface{}, 0)
				for i := initialValue; i <= finalValue; i++ {
					//6765 pink flamingo
					fizzBuzz, _ := pinkflamingo.FizzBuzz(i)
					if pinkflamingo.IsPinkFlmaingo(i) {
						result = append(result, "Pink Flamingo")
					} else if pinkflamingo.IsFibonacci(i) {
						result = append(result, "Flamingo")
					} else if fizzBuzz != "" {
						result = append(result, fizzBuzz)
					} else {
						result = append(result, i)
					}
				}
				b, err := json.Marshal(result)
				if err != nil {
					util.HandleError(writer, http.StatusInternalServerError, `{"Unable to Marshal Result`)
				}
				util.HandleSuccess(writer, b)
			} else {
				util.HandleError(writer, http.StatusBadRequest, `{"Please specify the range properly"}`)
			}

		} else {
			util.HandleError(writer, http.StatusBadRequest, `{"Please specify a range"}`)
		}

	}
}
