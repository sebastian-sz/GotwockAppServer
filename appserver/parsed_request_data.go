package appserver

import (
	"strings"
)

// Custom error that will be raised if the request doesn't contain mandatory fields.
type incompleteRequestError struct {
	message string
}

func (error *incompleteRequestError) Error() string {
	return error.message
}

// Struct describing data format expected to receive in the request.
// I am using pointers so that values missing in request are filled with easy-to-check <nil> values, rather than
// quite ambiguous, default zero values (there is no way to know whether the user has passed zero or did the Go language
// autofilled the value).
type parsedRequestData struct {
	Latitude    *float32 `json:"latitude"`
	Longitude   *float32 `json:"longitude"`
	MaxDistance *float32 `json:"maxDistance"`
}

// Check if required fields (Latitude, Longitude and MaxDistance) have been provided in the request.
// Returns incompleteRequestError if one or more required fields have been missing from the request body. If all fields
// are present, returns nil.
func (requestData *parsedRequestData) CheckIfAllDataProvided() error {
	initialMessage := "Missing required field(s) in request:"

	var missingFieldsMessage strings.Builder
	missingFieldsMessage.WriteString(initialMessage)

	if requestData.Longitude == nil {
		missingFieldsMessage.WriteString(" Longitude")
	}
	if requestData.Latitude == nil {
		missingFieldsMessage.WriteString(" Latitude")
	}
	if requestData.MaxDistance == nil {
		missingFieldsMessage.WriteString(" MaxDistance")
	}

	if missingFieldsMessage.String() != initialMessage {
		return &incompleteRequestError{message: missingFieldsMessage.String()}
	} else {
		return nil
	}

}
