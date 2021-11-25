package calculations

import "strings"

type MacrosStruct struct {
	protein int
	carbs int
	fats int
	TDEE float32
}

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

func Cals(BMRVal float32, intensity string) float32 {
	var calories float32
	var multiplier float32
	switch strings.ToUpper(intensity) {
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

func Macros(calories int, weight int, intensity string) MacrosStruct {
	var user_macros MacrosStruct
	user_macros.protein = weight
	if (intensity == "rest") {
	user_macros.carbs = int(weight * 0.5)
	}

	if (intensity == "light") {
	user_macros.carbs = int(weight * 1)
	}

	if (intensity == "moderate") {
	user_macros.carbs = int(weight * 1.5)
	}

	if (intensity == "hard") {
	user_macros.carbs = int(weight * 2)
	}
	user_macros.fats = int((calories - ((carbs + protein) * 4)) / 9)

	return user_macros
	}
