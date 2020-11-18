// This is a utility package for creating various combinations of location.LocationsProvider objects.
// In more detail: we can combine different distance estimation algorithms with different database connectors.
// To make it more clear I decided to keep all initialization code in this package.
package factory

import "github.com/sebastian-sz/GotwockAppServer/distance"

//Distance initializers:
func initializeHaversteinDistanceEstimator() distance.Estimator {
	var haversine distance.Estimator = &distance.Haversine{}
	return haversine
}
