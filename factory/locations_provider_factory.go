// This is a utility package for creating various combinations of location.LocationsProvider objects.
// In more detail: we can combine different distance estimation algorithms with different database connectors.
// To make it more clear I decided to keep all initialization code in this package.
package factory

import (
	"github.com/sebastian-sz/GotwockAppServer/dbconnectors"
	"github.com/sebastian-sz/GotwockAppServer/distance"
	"github.com/sebastian-sz/GotwockAppServer/location"
)

func initializeTouristLocationProvider(
	databaseConnector *dbconnectors.DatabaseConnector, distanceCalculator *distance.Estimator,
) location.LocationsProvider {
	return location.LocationsProvider{
		DatabaseConnector: databaseConnector,
		DistanceEstimator: distanceCalculator,
	}
}

func CreateJSONHaversineTouristLocationProvider() location.LocationsProvider {
	haversine := initializeHaversteinDistanceEstimator()
	jsonConnector := initializeJSONDatabaseConnector()
	return initializeTouristLocationProvider(&jsonConnector, &haversine)
}
