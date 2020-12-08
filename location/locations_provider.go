// Package, containing the location logic of this application. The main goal is to fetch and parse the data from the
// database and return an easily jsonable structure that can be exposed via Web Server.
// We make use of two other packages available in this repo, mainly distance (for calculating distance) and dbconnectors
// (for fetching data).
package location

import (
	"github.com/sebastian-sz/GotwockAppServer/dbconnectors"
	"github.com/sebastian-sz/GotwockAppServer/distance"
	"github.com/sebastian-sz/GotwockAppServer/model"
	"sort"
)

// The main struct of this package. It should accept initialised Database Connector and Distance Estimator. Keep in
// mind that pointers are being used in order not to replicate heavy objects (eg. JSONDataConnector).
type LocationsProvider struct {
	DistanceEstimator *distance.Estimator
	DatabaseConnector *dbconnectors.DatabaseConnector
}

// This method fetches data from DatabaseConnector. This data is parsed in a way that:
// 		1. Distance from the user is calculated (via DistanceEstimator) for each location.
//		2. If the distance from user is smaller than max distance, the Location object
//		is created and appended to the final results array (slice).
//		3. The results array (slice) is sorted based on the distance from the user.
func (t *LocationsProvider) GetAndParseLocationsData(
	userCoordinates model.Coordinates,
	maxDistanceFromUser float32,
) []model.Location {

	var results []model.Location
	singleDataFieldToIntMap := (*t.DatabaseConnector).ProvideData()

	for objectId, dataField := range singleDataFieldToIntMap {
		locationCoordinates := model.Coordinates{
			Latitude:  dataField.Latitude,
			Longitude: dataField.Longitude,
		}

		distanceToLocation := (*t.DistanceEstimator).EstimateDistance(userCoordinates, locationCoordinates)

		if distanceToLocation <= maxDistanceFromUser {
			touristLocation := model.Location{
				ObjectId:    int32(objectId),
				Name:        dataField.Name,
				Description: dataField.Description,
				Distance:    distanceToLocation,
				Position:    locationCoordinates,
			}

			results = append(results, touristLocation)
		}
	}
	sortResultsByDistance(results)

	return results
}

// Sorts the Location slice by distance from the user, ascending.
func sortResultsByDistance(resultSlice []model.Location) {
	sort.SliceStable(resultSlice, func(i, j int) bool {
		return (resultSlice[i]).Distance < ((resultSlice)[j]).Distance
	})
}
