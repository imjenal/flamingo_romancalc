package main

import (
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

	if err := AddRoutes(router); err != nil {
		log.Println("Failed adding routes", err.Error())
		return
	}

	fmt.Println("Application loaded successfully ")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}
