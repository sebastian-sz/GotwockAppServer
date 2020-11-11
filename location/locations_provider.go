package location

import (
	"github.com/sebastian-sz/GotwockAppServer/dbconnectors"
	"github.com/sebastian-sz/GotwockAppServer/distance"
)

type TouristLocationProvider struct {
	DistanceCalculator *distance.Calculator
	DatabaseConnector  *dbconnectors.DatabaseConnector
}
