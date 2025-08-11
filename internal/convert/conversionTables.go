package convert

var lengthConvertionTable = map[string]float64{
	"centimeter": 100,
	"meter":      1000,
	"kilometer":  1000000,
	"inch":       2.54,
	"yard":       91.44,
	"mile":       160934.4,
}

var weightConvertionTable = map[string]float64{
	"gram":     1000,
	"kilogram": 1000000,
	"ounce":    28349.5231,
	"pound":    453592.37,
}

var temperatureConvertionTable = map[string]float64{
	"fahrenheit": 33.8,
	"kelvin":     274.15,
}
