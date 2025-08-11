package convert

import (
	"errors"
	"fmt"
	"strconv"
)

func toMilimeter(value float64, valueName string) (float64, error) {
	if valueName == "millimeter" {
		return value, nil
	}

	convertKey, exists := lengthConvertionTable[valueName]
	if !exists {
		errorTxt := fmt.Sprintf("%v convertion value isn't supported in ToMilimeter", valueName)
		return 0, errors.New(errorTxt)
	}

	return value * convertKey, nil
}

func milimeterTo(value float64, valueName string) (float64, error) {
	if valueName == "millimeter" {
		return value, nil
	}

	convertKey, exists := lengthConvertionTable[valueName]
	if !exists {
		errorTxt := fmt.Sprintf("%v convertion value isn't supported in MilimeterTo", valueName)
		return 0, errors.New(errorTxt)
	}

	return value / convertKey, nil
}

func ConvertLength(rawValue string, from, to string) (float64, error) {
	length, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return 0, err
	}

	millimeter, err := toMilimeter(length, from)
	if err != nil {
		return 0, err
	}

	result, err := milimeterTo(millimeter, to)
	if err != nil {
		return 0, err
	}

	return result, nil
}
