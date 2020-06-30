package handler

import (
	"deeptrace/util"
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
			start, _ := strconv.Atoi(from)
			end, _ := strconv.Atoi(to)
			if start < end {
				util.HandleSuccess(writer, "Range is")
			} else {
				util.HandleError(writer, http.StatusBadRequest, `{"Please specify the range properly"}`)
			}

		} else {
			util.HandleError(writer, http.StatusBadRequest, `{"Please specify a range"}`)
		}

	}
}
