package convert

import (
	"errors"
	"fmt"
	"strconv"
)

func toMilligram(value float64, valueName string) (float64, error) {
	if valueName == "milligram" {
		return value, nil
	}

	convertKey, exists := weightConvertionTable[valueName]
	if !exists {
		errorTxt := fmt.Sprintf("%v convertion value isn't supported", valueName)
		return 0, errors.New(errorTxt)
	}

	return value * convertKey, nil
}

func milligramTo(value float64, valueName string) (float64, error) {
	if valueName == "milligram" {
		return value, nil
	}

	convertKey, exists := weightConvertionTable[valueName]
	if !exists {
		errorTxt := fmt.Sprintf("%v convertion value isn't supported", valueName)
		return 0, errors.New(errorTxt)
	}

	return value / convertKey, nil
}

func ConvertWeight(rawValue string, from, to string) (float64, error) {
	length, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return 0, err
	}

	milligram, err := toMilligram(length, from)
	if err != nil {
		return 0, err
	}

	result, err := milligramTo(milligram, to)
	if err != nil {
		return 0, err
	}

	return result, nil
}
