package distance

import (
	"github.com/sebastian-sz/GotwockAppServer/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_toRadians(t *testing.T) {
	tolerance := 1e-3
	testCases := []struct {
		degrees float32
		radians float64
	}{
		{
			degrees: 0,
			radians: 0,
		},
		{
			degrees: 180.0,
			radians: 3.142,
		},
		{
			degrees: 90,
			radians: 1.570,
		},
		{
			degrees: 1337,
			radians: 23.335,
		},
		{
			degrees: 360,
			radians: 6.283,
		},
	}

	for _, testCase := range testCases {
		calculatedRadians := toRadians(testCase.degrees)

		assert.InDelta(t, testCase.radians, calculatedRadians, tolerance)
	}
}

func TestHaversine_CalculateDistance(t *testing.T) {
	tolerance := 0.02
	haversineDistanceCalculator := Haversine{}
	testCases := []struct { // Runs parametrized test based on <input : expected output> pairs.
		firstCoordinates  model.Coordinates
		secondCoordinates model.Coordinates
		expectedDistance  float32
	}{
		{
			firstCoordinates:  model.Coordinates{Latitude: 38.898556, Longitude: -77.037852},
			secondCoordinates: model.Coordinates{Latitude: 38.897147, Longitude: -77.043934},
			expectedDistance:  0.549,
		},
		{
			firstCoordinates:  model.Coordinates{Latitude: 52.1039472, Longitude: 21.26832},
			secondCoordinates: model.Coordinates{Latitude: 52.1095869, Longitude: 21.2630788},
			expectedDistance:  0.707,
		},
		{
			firstCoordinates:  model.Coordinates{Latitude: 52.1087527, Longitude: 21.2657239},
			secondCoordinates: model.Coordinates{Latitude: 52.1142018, Longitude: 21.2895059},
			expectedDistance:  1.74,
		},
		{
			firstCoordinates:  model.Coordinates{Latitude: 52.1087183, Longitude: 21.2653876},
			secondCoordinates: model.Coordinates{Latitude: 52.0991107, Longitude: 21.27171360},
			expectedDistance:  1.17,
		},
		{
			firstCoordinates:  model.Coordinates{Latitude: 52.1024033, Longitude: 21.2710734},
			secondCoordinates: model.Coordinates{Latitude: 52.0903794, Longitude: 21.2796765},
			expectedDistance:  1.48,
		},
		{
			firstCoordinates:  model.Coordinates{Latitude: 52.2056663, Longitude: 21.1926853},
			secondCoordinates: model.Coordinates{Latitude: 52.1095869, Longitude: 21.2630788},
			expectedDistance:  11.71,
		},
	}

	for _, testCase := range testCases {
		actualDistance := testCase.expectedDistance
		calculatedDistance := haversineDistanceCalculator.CalculateDistance(
			testCase.firstCoordinates, testCase.secondCoordinates,
		)

		assert.InDelta(t, calculatedDistance, actualDistance, tolerance)

	}
}
