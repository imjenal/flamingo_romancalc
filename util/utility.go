package util

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func HandleError(writer http.ResponseWriter, code int, message string) {
	writer.WriteHeader(code)
	writer.Write([]byte(message))
}

func HandleSuccess(writer http.ResponseWriter, message []byte) {
	writer.WriteHeader(http.StatusOK)
	writer.Write(message)
}

func IsEmpty(data string) bool {
	return len(data) == 0
}

func IsEquals(input, testInput string) bool {
	return input == testInput
}

func GetUserID(request *http.Request) string {
	for k, userId := range mux.Vars(request) {
		if strings.EqualFold(k, "userId") {
			return userId
		}
	}
	return ""
}


func PanicHandler(w http.ResponseWriter, r *http.Request) {
	if r := recover(); r != nil {
		fmt.Sprintf("Recovered in f %v", r)
		var err error
		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = errors.New("Unknown panic")
		}
		if err != nil {
			fmt.Printf("FAILED")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
