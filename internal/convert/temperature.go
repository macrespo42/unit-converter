package convert

import (
	"errors"
	"fmt"
	"strconv"
)

func toCelsius(value float64, valueName string) (float64, error) {
	if valueName == "celsius" {
		return value, nil
	}

	convertKey, exists := temperatureConvertionTable[valueName]
	if !exists {
		errorTxt := fmt.Sprintf("%v convertion value isn't supported", valueName)
		return 0, errors.New(errorTxt)
	}

	return value * convertKey, nil
}

func celsiusTo(value float64, valueName string) (float64, error) {
	if valueName == "celsius" {
		return value, nil
	}

	convertKey, exists := temperatureConvertionTable[valueName]
	if !exists {
		errorTxt := fmt.Sprintf("%v convertion value isn't supported", valueName)
		return 0, errors.New(errorTxt)
	}

	return value / convertKey, nil
}

func ConvertTemperature(rawValue string, from, to string) (float64, error) {
	length, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return 0, err
	}

	celsius, err := toCelsius(length, from)
	if err != nil {
		return 0, err
	}

	result, err := celsiusTo(celsius, to)
	if err != nil {
		return 0, err
	}

	return result, nil
}
