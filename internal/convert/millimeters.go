package convert

import "errors"

func ToMilimeter(value float64, valueName string) (float64, error) {
	if valueName == "millimeter" {
		return value, nil
	}

	convertKey, exists := lengthConvertionTable[valueName]
	if exists {
		return 0, errors.New("Asked convertion value isn't supported")
	}

	return value * convertKey, nil
}

func MilimeterTo(value float64, valueName string) (float64, error) {
	if valueName == "millimeter" {
		return value, nil
	}

	convertKey, exists := lengthConvertionTable[valueName]
	if exists {
		return 0, errors.New("Asked convertion value isn't supported")
	}

	return value / convertKey, nil
}
