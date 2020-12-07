package appserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parsedRequestData_CheckIfAllDataProvided_BadRequest(t *testing.T) {
	testCaseParameters := []struct {
		requestData         parsedRequestData
		expectedRaisedError incompleteRequestError
	}{
		{ // Case Latitude has not been provided
			requestData: parsedRequestData{
				Longitude:   pointerOf(1.0),
				MaxDistance: pointerOf(1.0),
			},
			expectedRaisedError: incompleteRequestError{"Missing required field(s) in request: Latitude"},
		},
		{ // Case Longitude has not been provided
			requestData: parsedRequestData{
				Latitude:    pointerOf(1.0),
				MaxDistance: pointerOf(1.0),
			},
			expectedRaisedError: incompleteRequestError{"Missing required field(s) in request: Longitude"},
		},
		{ // Case both Longitude and Latitude has not been provided
			requestData: parsedRequestData{MaxDistance: pointerOf(1.0)},
			expectedRaisedError: incompleteRequestError{
				"Missing required field(s) in request: Longitude Latitude",
			},
		},
		{ // Case all the values were missing from the request body:
			requestData: parsedRequestData{},
			expectedRaisedError: incompleteRequestError{
				"Missing required field(s) in request: Longitude Latitude MaxDistance",
			},
		},
	}

	for _, testCase := range testCaseParameters {
		missingDataError := testCase.requestData.CheckIfAllDataProvided()
		assert.Equal(t, &testCase.expectedRaisedError, missingDataError)
	}
}

func Test_parsedRequestData_CheckIfAllDataProvided_GoodRequest(t *testing.T) {
	requestData := parsedRequestData{
		Latitude:    pointerOf(1.0),
		Longitude:   pointerOf(1.0),
		MaxDistance: pointerOf(1.0),
	}

	missingDataError := requestData.CheckIfAllDataProvided()

	assert.Equal(t, nil, missingDataError)

}

// Utility function to quickly fetch pointer of variables on the fly.
func pointerOf(x float32) *float32 {
	return &x
}
