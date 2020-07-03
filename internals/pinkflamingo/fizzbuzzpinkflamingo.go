package pinkflamingo

import (
	"deeptrace/util"
	"encoding/json"
	"fmt"
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
			result = computeForEachNumber(i, result)
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

func computeForEachNumber(i int, result []interface{}) []interface{}{
	fizzBuzz, _ := FizzBuzz(i)
	if IsPinkFlmaingo(i) {
		result = append(result, PINKFLAMINGO)
		fmt.Println(PINKFLAMINGO)
	} else if IsFibonacci(i) {
		result = append(result, FLAMINGO)
		fmt.Println(FLAMINGO)
	} else if fizzBuzz != "" {
		result = append(result, fizzBuzz)
		fmt.Println(fizzBuzz)
	} else {
		result = append(result, i)
		fmt.Println(i)
	}
	fmt.Println(i)
	return result
}
