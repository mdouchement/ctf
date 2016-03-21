package extendedmath

import "math"

// Round rounds a float64 :D
// Credits https://gist.github.com/DavidVaini/10308388
func Round(a float64) float64 {
	if a < 0 {
		return math.Ceil(a - 0.5)
	}
	return math.Floor(a + 0.5)
}

// RoundPlus rounds a float64 :D
//   RoundPlus(123.555555, 3)
//   => 123.556
//
// Credits https://gist.github.com/DavidVaini/10308388
func RoundPlus(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return Round(f*shift) / shift
}
