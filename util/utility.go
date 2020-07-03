package util

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Response(writer http.ResponseWriter, code int, message string) {
	writer.WriteHeader(code)
	writer.Write([]byte(message))
}

func ResponseSuccess(writer http.ResponseWriter, message []byte) {
	writer.WriteHeader(http.StatusOK)
	writer.Write(message)
}

func IsEmpty(data string) bool {
	return len(data) == 0
}

func IsValidData(data string) bool {
	return len(data) > 0
}

func IsEquals(input, testInput interface{}) bool {
	return input == testInput
}

func IsSmaller(data1, data2 int) bool {
	return data1 < data2
}

func SplitBySeparator(data, separator string) []string{
	return strings.Split(data, separator)
}

func ToInteger(data string) (int, error) {
	value, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func ToString(data int) (string) {
	return strconv.Itoa(data)
}

func ToLowerCase(data string) string {
	return strings.ToLower(data)
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
