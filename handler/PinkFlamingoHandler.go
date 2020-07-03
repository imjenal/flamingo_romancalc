package handler

import (
	"deeptrace/internals/pinkflamingo"
	"deeptrace/util"
	"net/http"
)

func PinkFlamingoHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer util.PanicHandler(writer, request)
		defer request.Body.Close()
		writer.Header().Set("Content-Type", "application/json")
		to, from := extractQueryParams(request)
		if isValidQueryParams(to, from) {
			pinkflamingo.FizzBuzzPinkFlamingo{}.Execute(writer, to, from)
		} else {
			util.Response(writer, http.StatusBadRequest, `{"Please specify a range"}`)
		}
	}
}

func extractQueryParams(request *http.Request) (string, string) {
	to := request.FormValue(util.ToLowerCase("To"))
	from := request.FormValue(util.ToLowerCase("From"))
	return to, from
}

func isValidQueryParams(to, from string) bool {
	return util.IsValidData(to) && util.IsValidData(from)
}
