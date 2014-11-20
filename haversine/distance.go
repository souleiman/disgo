package haversine

import "math"

// ToRadians returns a radians from the converted degree value given.
func ToRadians(degree float64) float64 {
	return (degree * math.Pi) / 180
}

// Distance takes in a (latitude, longitude) of source and destination and computes the distance between the two points
// and with the provided radius of the given spherical atmosphere.
// Uses Haversine formula: http://en.wikipedia.org/wiki/Haversine_formula
func Distance(src_lat, src_long, dest_lat, dest_long, radius float64) float64 {
	lat := Haversine(ToRadians(dest_lat - src_lat))
	long := Haversine(ToRadians(dest_long - src_long))

	return 2 * radius * math.Asin(math.Sqrt(lat+math.Cos(ToRadians(src_lat))*math.Cos(ToRadians(dest_lat))*long))
}
