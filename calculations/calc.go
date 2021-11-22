package calculations

import "strings"

func BMR(Sex string, Age int, HeightInches int, Weight int) float32 {
	var BMRVal float32
	BMRVal = 0.0
	switch strings.ToUpper(Sex) {
	case "MALE":
		BMRVal = (10 * (Weight / 2.2)) + (6.25 * (HeightInches * 2.54)) - ((5 * Age) + 5) //  Mifflin-St. Jeor formula
		break;
	case "FEMALE":
		BMRVal = (10 * (Weight / 2.2)) + (6.25 * (HeightInches * 2.54)) - ((5 * Age) + 5) - 161 //  Mifflin-St. Jeor formula
	}
	return BMRVal
}