package factory

import (
	"github.com/sebastian-sz/GotwockAppServer/dbconnectors"
	"github.com/sebastian-sz/GotwockAppServer/distance"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateJSONHaversineTouristLocationProvider(t *testing.T) {
	locationsProvider := CreateJSONHaversineTouristLocationProvider()

	_, distanceOk := (*locationsProvider.DistanceEstimator).(distance.Estimator)
	_, databaseOk := (*locationsProvider.DatabaseConnector).(dbconnectors.DatabaseConnector)

	assert.True(t, distanceOk)
	assert.True(t, databaseOk)
}
