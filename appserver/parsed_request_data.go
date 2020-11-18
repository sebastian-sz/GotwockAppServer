package appserver

import (
	"strings"
)

// Struct describing data format expected to receive in the request.
// I am extending this struct with one method: checkMissingValues(). Go will auto set all missing float32 to 0, so the
// only way to make sure they are present (and required) in the request body is to check whether they are equal to zero.
// This is not the nicest way, but it is not expected in the real world scenario to provide exactly 0 longitude or
// exactly zero latitude.
type parsedRequestData struct {
	Latitude            float32 `json:"latitude"`
	Longitude           float32 `json:"longitude"`
	MaxDistanceFromUser float32 `json:"maxDistance"`
}

// Check if required fields (Latitude or Longitude) have been set with default zero values.
// If they were that would indicate that Go has autofilled them with zeros, which means they were missing in the
// request body.
// This method prevents weird behaviour where you could receive the entire database after passing '{}' in the request
// body.
// Returns string message indicating which fields where missing and boolean for quick check if anything was missing.
func (requestData *parsedRequestData) checkMissingValues() (string, bool) {
	initialMessage := "Missing required field(s) in request:"

	var missingFieldsMessage strings.Builder
	missingFieldsMessage.WriteString(initialMessage)

	if requestData.Longitude == 0 {
		missingFieldsMessage.WriteString(" Longitude")
	}
	if requestData.Latitude == 0 {
		missingFieldsMessage.WriteString(" Latitude")
	}

	if missingFieldsMessage.String() != initialMessage {
		return missingFieldsMessage.String(), true
	} else {
		return "", false
	}

}
