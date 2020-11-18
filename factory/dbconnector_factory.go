// This is a utility package for creating various combinations of location.LocationsProvider objects.
// In more detail: we can combine different distance estimation algorithms with different database connectors.
// To make it more clear I decided to keep all initialization code in this package.
package factory

import (
	"github.com/sebastian-sz/GotwockAppServer/dbconnectors"
	"github.com/sebastian-sz/GotwockAppServer/projectpath"
	"path/filepath"
)

// Database connector initializers
func initializeJSONDatabaseConnector() dbconnectors.DatabaseConnector {

	rootPath := projectpath.GetRootPath()
	dataPath := filepath.Join(rootPath, "data/otwock_db.json")
	var jsonDBConnector dbconnectors.DatabaseConnector = &dbconnectors.JSONDataConnector{
		DataPath:   dataPath,
		CachedData: nil,
	}
	jsonDBConnector.Initialize()
	return jsonDBConnector
}
