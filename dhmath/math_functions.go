package dhmath

import (
	"math"
)

func DegToRad(deg float32) float32 {
	return math.Pi / 180.0 * deg
}
func RadToDeg(rad float32) float32 {
	return 180.0 * rad / math.Pi
}

func Clamp_int(value int, min int, max int) int {
	var ret int

	if value < min {
		ret = min
	} else if value > max {
		ret = max
	} else {
		ret = value
	}

	return ret
}
func Clamp_float32(value float32, min float32, max float32) float32 {
	var ret float32

	if value < min {
		ret = min
	} else if value > max {
		ret = max
	} else {
		ret = value
	}

	return ret
}
