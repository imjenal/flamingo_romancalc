package handler

import (
	"deeptrace/util"
	"net/http"
)

func PinkFlamingoHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer util.PanicHandler(writer, request)
		writer.Header().Set("Content-Type", "application/json")
		util.HandleError(writer, http.StatusBadRequest, "Invalid Hit")
		defer request.Body.Close()
	}
}


