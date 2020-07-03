package pinkflamingo

import (
	"deeptrace/util"
	"encoding/json"
	"net/http"
)

const (
	PINKFLAMINGO = "Pink Flamingo"
	FLAMINGO     = "Flamingo"
)

type FizzBuzzPinkFlamingo struct {
}

func (f FizzBuzzPinkFlamingo) Execute(writer http.ResponseWriter, from, to string) {
	initialValue, finalValue := convertToIntegers(writer, from, to)
	if util.IsSmaller(initialValue, finalValue) {
		result := make([]interface{}, 0)
		for i := initialValue; i <= finalValue; i++ {
			computedValue := computeForEachNumber(i)
			result = append(result, computedValue)
		}
		marshalledResult, err := json.Marshal(result)
		if err != nil {
			util.Response(writer, http.StatusInternalServerError, `{"Unable to Marshal Result}`)
		}
		util.ResponseSuccess(writer, marshalledResult)
	} else {
		util.Response(writer, http.StatusBadRequest, `{"Please specify the range properly"}`)

	}
}

func convertToIntegers(writer http.ResponseWriter, to, from string) (int, int) {
	initialValue, err := util.ToInteger(from)
	finalValue, err := util.ToInteger(to)
	if err != nil {
		util.Response(writer, http.StatusBadRequest, `{"Unable to Convert String to Integer"}`)
	}
	return initialValue, finalValue
}

func computeForEachNumber(number int) interface{} {
	fizzBuzz, _ := FizzBuzz(number)
	if IsPinkFlamingo(number) {
		return PINKFLAMINGO
	} else if IsFibonacci(number) {
		return FLAMINGO
	} else if fizzBuzz != "" {
		return  fizzBuzz
	} else {
		return number
	}
}
