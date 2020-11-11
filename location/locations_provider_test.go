package location

import (
	"github.com/sebastian-sz/GotwockAppServer/dbconnectors"
	"github.com/sebastian-sz/GotwockAppServer/distance"
	"github.com/sebastian-sz/GotwockAppServer/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Run parametrized test, to check whether function isDistanceOk correctly returns boolean value.
func Test_isDistanceOk(t *testing.T) {
	testCases := []struct {
		maxDistance    float32
		userDistance   float32
		expectedResult bool
	}{
		{ // Case max distance = 0, return all results.
			maxDistance:    0.0,
			userDistance:   1337.0,
			expectedResult: true,
		},
		{ // Case max distance = 0, see if userDistance doesn't matter.
			maxDistance:    0.0,
			userDistance:   0.0,
			expectedResult: true,
		},
		{ // Case maxDistance > userDistance
			maxDistance:    10.0,
			userDistance:   5.0,
			expectedResult: true,
		},
		{ // Case maxDistance < userDistance
			maxDistance:    10.0,
			userDistance:   15.0,
			expectedResult: false,
		},
		{ // Case maxDistance non zero and equal to userDistance
			maxDistance:    10.0,
			userDistance:   10.0,
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		outcome := isDistanceOk(testCase.maxDistance, testCase.userDistance)
		assert.Equal(t, testCase.expectedResult, outcome)

	}
}

// Run parametrized test to check, whether function sortResultsByDistance correctly sorts (in-place) given slice.
func Test_sortResultsByDistance(t *testing.T) {
	firstLocation := model.TouristLocation{
		Distance: 1.0,
	}
	secondLocation := model.TouristLocation{
		Distance: 8.0,
	}
	thirdLocation := model.TouristLocation{
		Distance: 16.0,
	}

	testCases := []struct {
		inputSlice    []model.TouristLocation
		expectedSlice []model.TouristLocation
	}{
		{ // Case mixed two elements
			inputSlice:    []model.TouristLocation{firstLocation, thirdLocation, secondLocation},
			expectedSlice: []model.TouristLocation{firstLocation, secondLocation, thirdLocation},
		},
		{ // Case the slice is reversed
			inputSlice:    []model.TouristLocation{thirdLocation, secondLocation, firstLocation},
			expectedSlice: []model.TouristLocation{firstLocation, secondLocation, thirdLocation},
		},
		{ // Case empty slices
			inputSlice:    []model.TouristLocation{},
			expectedSlice: []model.TouristLocation{},
		},
	}

	for _, testCase := range testCases {
		sortResultsByDistance(testCase.inputSlice)
		assert.Equal(t, testCase.inputSlice, testCase.expectedSlice)

	}
}

func initializeJsonHaversineTouristLocationProvider() TouristLocationProvider {
	var haversine distance.Calculator = &distance.Haversine{}
	var jsonDbConnector dbconnectors.DatabaseConnector = &dbconnectors.JSONDataConnector{
		DataPath:   "../data/otwock_db.json",
		CachedData: nil,
	}

	jsonDbConnector.Initialize()

	return TouristLocationProvider{
		DistanceCalculator: &haversine,
		DatabaseConnector:  &jsonDbConnector,
	}
}

// Test end-to-end flow for TouristLocationProvider
func TestTouristLocationProvider_GetAndParseTouristLocationData(t *testing.T) {
	tolerance := 0.002
	touristLocationProvider := initializeJsonHaversineTouristLocationProvider()
	userCoordinates := model.Coordinates{Latitude: 52.0989711, Longitude: 21.2715719} // City Park
	maxDistance := float32(5.0)
	expectedUserLocations := []model.TouristLocation{
		{
			ObjectId:    1,
			Name:        "City Hall",
			Description: "City Hall of the Otwock city.",
			Distance:    0.596,
			Position:    model.Coordinates{Latitude: 52.1039472, Longitude: 21.26832},
		},
		{
			ObjectId:    2,
			Name:        "Railway Station",
			Description: "Railway Station of the Otwock city.",
			Distance:    1.3157,
			Position:    model.Coordinates{Latitude: 52.1095869, Longitude: 21.2630788},
		},
		{
			ObjectId:    5,
			Name:        "Swiss village",
			Description: "Swiss village of the Otwock city.",
			Distance:    1.6929,
			Position:    model.Coordinates{Latitude: 52.1097756, Longitude: 21.2890232},
		},
		{
			ObjectId:    3,
			Name:        "Jozefow Railway Station",
			Description: "Railway Station of the Jozefow city.",
			Distance:    4.7878,
			Position:    model.Coordinates{Latitude: 52.1362265, Longitude: 21.2364616},
		},
	}

	actualUserLocations := touristLocationProvider.GetAndParseTouristLocationData(userCoordinates, maxDistance)

	// Due to floating point number (Distance) we must compare specifically:
	assert.Equal(t, len(expectedUserLocations), len(actualUserLocations))

	for resultIdx, expectedResult := range expectedUserLocations {
		actualResult := actualUserLocations[resultIdx]

		assert.Equal(t, expectedResult.Name, actualResult.Name)
		assert.Equal(t, expectedResult.Description, actualResult.Description)
		assert.Equal(t, expectedResult.ObjectId, actualResult.ObjectId)
		assert.Equal(t, expectedResult.Position, actualResult.Position)
		assert.InDelta(t, expectedResult.Distance, actualResult.Distance, tolerance)

	}
}
