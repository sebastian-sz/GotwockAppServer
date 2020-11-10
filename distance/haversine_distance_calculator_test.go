package distance

import (
	"github.com/sebastian-sz/GotwockAppServer/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

const tolerance = 0.02

func TestHaversine_CalculateDistance(t *testing.T) {
	haversineDistanceCalculator := Haversine{}
	parameters := []struct { // Runs parametrized test based on <input : expected output> pairs.
		firstCoordinates  model.Coordinates
		secondCoordinates model.Coordinates
		expectedDistance  float32
	}{
		{
			model.Coordinates{Latitude: 38.898556, Longitude: -77.037852},
			model.Coordinates{Latitude: 38.897147, Longitude: -77.043934},
			0.549,
		},
		{
			model.Coordinates{Latitude: 52.1039472, Longitude: 21.26832},
			model.Coordinates{Latitude: 52.1095869, Longitude: 21.2630788},
			0.707,
		},
		{
			model.Coordinates{Latitude: 52.1087527, Longitude: 21.2657239},
			model.Coordinates{Latitude: 52.1142018, Longitude: 21.2895059},
			1.74,
		},
		{
			model.Coordinates{Latitude: 52.1087183, Longitude: 21.2653876},
			model.Coordinates{Latitude: 52.0991107, Longitude: 21.27171360},
			1.17,
		},
		{
			model.Coordinates{Latitude: 52.1024033, Longitude: 21.2710734},
			model.Coordinates{Latitude: 52.0903794, Longitude: 21.2796765},
			1.48,
		},
		{
			model.Coordinates{Latitude: 52.2056663, Longitude: 21.1926853},
			model.Coordinates{Latitude: 52.1095869, Longitude: 21.2630788},
			11.71,
		},
	}

	for i := range parameters {
		actualDistance := parameters[i].expectedDistance
		calculatedDistance := haversineDistanceCalculator.CalculateDistance(
			parameters[i].firstCoordinates, parameters[i].secondCoordinates,
		)

		assert.InDelta(t, calculatedDistance, actualDistance, tolerance)

	}
}
