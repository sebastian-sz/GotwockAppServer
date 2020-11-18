package appserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parsedRequestData_checkMissingValues(t *testing.T) {
	testCaseParameters := []struct {
		requestData     parsedRequestData
		expectedBool    bool
		expectedMessage string
	}{
		{ // Case both values are non-zero
			requestData: parsedRequestData{
				Latitude:  1.0,
				Longitude: 1.0,
			},
			expectedBool:    false,
			expectedMessage: "",
		},
		{ // Case Latitude has been set with default value
			requestData: parsedRequestData{
				Longitude: 1.0,
			},
			expectedBool:    true,
			expectedMessage: "Missing required field(s) in request: Latitude",
		},
		{ // Case Longitude has been set with default value
			requestData: parsedRequestData{
				Latitude: 1.0,
			},
			expectedBool:    true,
			expectedMessage: "Missing required field(s) in request: Longitude",
		},
		{ // Case both Longitude and Latitude are set with default values
			requestData:     parsedRequestData{},
			expectedBool:    true,
			expectedMessage: "Missing required field(s) in request: Longitude Latitude",
		},
	}

	for _, testCase := range testCaseParameters {
		message, boolCheck := testCase.requestData.checkMissingValues()
		assert.Equal(t, testCase.expectedBool, boolCheck)
		assert.Equal(t, testCase.expectedMessage, message)
	}
}
