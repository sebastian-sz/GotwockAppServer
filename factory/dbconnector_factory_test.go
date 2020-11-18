package factory

import (
	"github.com/sebastian-sz/GotwockAppServer/dbconnectors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_initializeJSONDatabaseConnector(t *testing.T) {
	dataBaseConnector := initializeJSONDatabaseConnector()

	_, ok := dataBaseConnector.(dbconnectors.DatabaseConnector)
	dataBaseConnector.Initialize()

	assert.True(t, ok)
}
