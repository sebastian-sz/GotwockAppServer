// This is a utility file for creating various dbconnectors.DatabaseConnector objects.
package factory

import (
	"github.com/sebastian-sz/GotwockAppServer/dbconnectors"
	"github.com/sebastian-sz/GotwockAppServer/projectpath"
	"path/filepath"
)

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
