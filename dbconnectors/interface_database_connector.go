// Package for fetching data from various data sources.
package dbconnectors

import "github.com/sebastian-sz/GotwockAppServer/model"

// Interface that defines methods for various Data Base Connectors.
// Each database connector should have at lease a single method: ProvideData().
// This method returns a map of object id (int) to generic data field (SingleDataField).
// This map should be later passed to other components who will parse it,
// calculate distance and form into proper response (suitable for sending over REST or other services).
type DatabaseConnector interface {
	ProvideData() map[int]model.SingleDataField
}
