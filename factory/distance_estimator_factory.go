// This is a utility file for creating various distance.Estimator objects.
package factory

import "github.com/sebastian-sz/GotwockAppServer/distance"

func initializeHaversteinDistanceEstimator() distance.Estimator {
	var haversine distance.Estimator = &distance.Haversine{}
	return haversine
}
