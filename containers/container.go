package containers

import (
	"github.com/sebastian-sz/GotwockAppServer/dbconnectors"
	"github.com/sebastian-sz/GotwockAppServer/distance"
	"github.com/sebastian-sz/GotwockAppServer/location"
)

//Distance initializers:
func initializeHaverstein() distance.Calculator {
	var haversine distance.Calculator = &distance.Haversine{}
	return haversine
}

// Database connector initializers
func initializeJSONDatabaseConnector() dbconnectors.DatabaseConnector {
	dataPath := "data/otwock_db.json"
	var jsonDBConnector dbconnectors.DatabaseConnector = &dbconnectors.JSONDataConnector{
		DataPath:   dataPath,
		CachedData: nil,
	}
	jsonDBConnector.Initialize()
	return jsonDBConnector
}

// Tourist Location Provider initializers
func initializeTouristLocationProvider(
	databaseConnector *dbconnectors.DatabaseConnector, distanceCalculator *distance.Calculator,
) location.TouristLocationProvider {
	return location.TouristLocationProvider{
		DatabaseConnector:  databaseConnector,
		DistanceCalculator: distanceCalculator,
	}
}

func CreateJSONHaversineTouristLocationProvider() location.TouristLocationProvider {
	haversine := initializeHaverstein()
	jsonConnector := initializeJSONDatabaseConnector()
	return initializeTouristLocationProvider(&jsonConnector, &haversine)
}
