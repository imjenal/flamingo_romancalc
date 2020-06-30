package handler

import (
	"deeptrace/internals/pinkflamingo"
	"deeptrace/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func PinkFlamingoHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer util.PanicHandler(writer, request)
		defer request.Body.Close()
		writer.Header().Set("Content-Type", "application/json")

		to := request.FormValue(strings.ToLower("To"))
		from := request.FormValue(strings.ToLower("From"))
		if len(to) > 0 && len(from) > 0 {
			start, err := strconv.Atoi(from)
			end, err := strconv.Atoi(to)
			if err != nil {
				util.HandleError(writer, http.StatusBadRequest, `{"Unable to Convert String to Integer"}`)
				return
			}
			if start < end {
				result := make([]interface{}, 0)
				for i := start; i <= end; i++ {
					//6765 pink flamingo
					fizzBuzz, _ := pinkflamingo.FizzBuzz(i)
					if pinkflamingo.IsPinkFlmaingo(i) {
						result = append(result, "Pink Flamingo")
					} else if pinkflamingo.IsFibonacci(i) {
						result = append(result, "Flamingo")
					} else if fizzBuzz != "" {
						result = append(result, fizzBuzz)
					} else {
						fmt.Println(i)
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
