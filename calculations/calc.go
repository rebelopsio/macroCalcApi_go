package calculations

import "strings"

func BMR(Sex string, Age int, HeightInches int, Weight int) float32 {
	var BMRVal float32
	BMRVal = 0.0
	switch strings.ToUpper(Sex) {
	case "MALE":
		BMRVal = (10 * (Weight / 2.2)) + (6.25 * (HeightInches * 2.54)) - ((5 * Age) + 5) //  Mifflin-St. Jeor formula
	case "FEMALE":
		BMRVal = (10 * (Weight / 2.2)) + (6.25 * (HeightInches * 2.54)) - ((5 * Age) + 5) - 161 //  Mifflin-St. Jeor formula
	}
	return BMRVal
}

func Cals(BMRVal float32, Type string) float32 {
	var calories float32
	var multiplier float32
	switch strings.ToUpper(Type) {
	case "REST":
		multiplier = 1.2
	case "LIGHT":
		multiplier = 1.375
	case "MODERATE":
		multiplier = 1.55
	case "HARD":
		multiplier = 1.725
	}
	calories = BMRVal * multiplier
	return calories
}