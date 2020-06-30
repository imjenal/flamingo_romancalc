package main

import (
	"deeptrace/handler"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	allowedHeaders := handlers.AllowedHeaders([]string{"Origin", "Accept", "Content-Type", "Authorization", "DateUsed", "X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions})

	router.Handle("/pinkflamingo", handler.PinkFlamingoHandler()).Methods(http.MethodGet)
	router.Handle("/romancalc", handler.RomanCalculatorHandler()).Methods(http.MethodGet)

	fileServer := http.FileServer(http.Dir("./swaggerui/"))
	router.PathPrefix("/api/").Handler(http.StripPrefix("/api/", fileServer))

	fmt.Println("Application loaded successfully ")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}
