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
				Longitude:   pointerTo(1.0),
				MaxDistance: pointerTo(1.0),
			},
			expectedRaisedError: incompleteRequestError{"Missing required field(s) in request: Latitude"},
		},
		{ // Case Longitude has not been provided
			requestData: parsedRequestData{
				Latitude:    pointerTo(1.0),
				MaxDistance: pointerTo(1.0),
			},
			expectedRaisedError: incompleteRequestError{"Missing required field(s) in request: Longitude"},
		},
		{ // Case both Longitude and Latitude has not been provided
			requestData: parsedRequestData{MaxDistance: pointerTo(1.0)},
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
		Latitude:    pointerTo(1.0),
		Longitude:   pointerTo(1.0),
		MaxDistance: pointerTo(1.0),
	}

	missingDataError := requestData.CheckIfAllDataProvided()

	assert.Equal(t, nil, missingDataError)

}

// Utility function to quickly fetch pointer of the variable.
func pointerTo(x float32) *float32 {
	return &x
}
