package haversine

import "math"

// Haversine function returns the result based on the given radians.
func Haversine(radian float64) float64 {
	return (1 - math.Cos(radian)) / 2
}
