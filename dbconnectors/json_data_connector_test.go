package dbconnectors

import (
	"github.com/sebastian-sz/GotwockAppServer/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

const mockDataPath = "./mock.json"

func makeExpectedLoadedContent() map[int]model.SingleLocationData {
	expectedLoadedContent := make(map[int]model.SingleLocationData)
	expectedLoadedContent[1] = model.SingleLocationData{
		Latitude:    52.1039472,
		Longitude:   21.26832,
		Name:        "City Hall",
		Description: "City Hall of the Otwock city.",
	}
	expectedLoadedContent[2] = model.SingleLocationData{
		Latitude:    52.1095869,
		Longitude:   21.2630788,
		Name:        "Railway Station",
		Description: "Railway Station of the Otwock city.",
	}
	return expectedLoadedContent
}

func TestJSONDataConnector_LoadContentToMemory(t *testing.T) {
	jsonDBConnector := JSONDataConnector{
		DataPath:   mockDataPath,
		CachedData: nil,
	}
	expectedLoadedContent := makeExpectedLoadedContent()

	jsonDBConnector.Initialize()

	assert.Equal(t, expectedLoadedContent, jsonDBConnector.CachedData)

}

func TestJSONDataConnector_ProvideData(t *testing.T) {
	jsonDBConnector := JSONDataConnector{
		DataPath:   mockDataPath,
		CachedData: nil,
	}
	expectedLoadedContent := makeExpectedLoadedContent()

	jsonDBConnector.Initialize()
	providedData := jsonDBConnector.ProvideData()

	assert.Equal(t, providedData, expectedLoadedContent)
}
