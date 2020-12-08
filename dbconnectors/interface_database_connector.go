// Package for fetching data from various data sources.
package dbconnectors

import "github.com/sebastian-sz/GotwockAppServer/model"

// Interface that defines methods for various Data Base Connectors.
// Each database connector should override two methods: Initialize() and ProvideData().
// Initialize() handles the logic of initializing the connection to the database.
// ProvideData() is the core method of this struct. This method returns a map of object id (int) to generic data
// field (SingleLocationData).
// This map should be later passed to other components who will parse it,
// calculate distance and form into proper response (suitable for sending over REST or other services).
type DatabaseConnector interface {
	Initialize()
	ProvideData() map[int]model.SingleLocationData
}
