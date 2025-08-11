package convert

import (
	"errors"
	"fmt"
)

func ToMilimeter(value float64, valueName string) (float64, error) {
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

func MilimeterTo(value float64, valueName string) (float64, error) {
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
