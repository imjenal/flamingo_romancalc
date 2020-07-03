package http

import (
	"deeptrace/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func AddRoutes(router *mux.Router) error {
	router.Handle("/pinkflamingo", handler.PinkFlamingoHandler()).Methods(http.MethodGet)
	router.Handle("/romancalc", handler.RomanCalculatorHandler()).Methods(http.MethodPost)
	fileServer := http.FileServer(http.Dir("./swaggerui/"))
	router.PathPrefix("/api/").Handler(http.StripPrefix("/api/", fileServer))

	return nil
}
