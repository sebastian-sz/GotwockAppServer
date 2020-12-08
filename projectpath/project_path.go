// This is a utility package for managing relative / absolute file paths used in this repository.
// For example dbconnectors.JSONDataConnector uses data stored in "data/otwock_db.json". If we pass this path
// directly, we can get IO Errors depending on the directory from which we start the GotwockAppServer (assuming that
// dbconnectors.JSONDataConnector is used).
// In order to fix this issue one can invoke GetRootPath and join the result with relative path from the directory.
// That way components like dbconnectors.JSONDataConnector are going to receive full, absolute path and it will be
// possible to run the server from different directories.
package projectpath

import (
	"path/filepath"
	"runtime"
)

func GetRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../")
}
