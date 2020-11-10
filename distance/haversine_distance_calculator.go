package distance

import (
	"github.com/sebastian-sz/GotwockAppServer/model"
	"math"
)

const earthRadiusInKM = 6373

// Struct implementing Haversine formula for distance approximation.
// is suitable for calculating short distances and non-critical applications (like not boat sailing or navigating).
// which apply here as we just want the distance estimate.
// Source:
// 		Blog: https://andrew.hedges.name/experiments/haversine/
//		Wikipedia: https://en.wikipedia.org/wiki/Haversine_formula
type Haversine struct{}

// Returns the approximated distance based on Haversine formula.
func (h *Haversine) CalculateDistance(firstCoordinates, secondCoordinates model.Coordinates) float32 {
	lon1 := toRadians(firstCoordinates.Longitude)
	lon2 := toRadians(secondCoordinates.Longitude)
	lat1 := toRadians(firstCoordinates.Latitude)
	lat2 := toRadians(secondCoordinates.Latitude)

	dLon := lon2 - lon1
	dLat := lat2 - lat1

	haversine := haversineFunc(dLat) + math.Cos(lat1)*math.Cos(lat2)*haversineFunc(dLon)
	return float32(2 * earthRadiusInKM * math.Asin(math.Sqrt(haversine)))
}

func toRadians(value float32) float64 {
	return float64(value * math.Pi / 180)
}

func haversineFunc(value float64) float64 {
	return math.Pow(math.Sin(value/2), 2)
}
