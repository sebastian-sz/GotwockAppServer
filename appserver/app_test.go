package appserver

import (
	"github.com/gorilla/mux"
	"github.com/sebastian-sz/GotwockAppServer/factory"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const requestHeader = "application/json"

func setUpTestRouterAndEndpoint() (*mux.Router, string) {
	locationsProvider := factory.CreateJSONHaversineTouristLocationProvider()
	app := App{
		EndpointPath:      "/",
		ServerAddr:        "127.0.0.1:9100",
		LocationsProvider: &locationsProvider,
	}
	appRouter := app.makeRouter()
	endpointPath := "http://" + app.ServerAddr + app.EndpointPath
	return appRouter, endpointPath
}

func TestAppHandlingBadRequests(t *testing.T) {
	appRouter, endpointPath := setUpTestRouterAndEndpoint()

	testCasesParameters := []struct {
		headerType      string
		requestData     string
		expectedCode    int
		expectedMessage string
	}{
		{ // Faulty header
			headerType:      "not-application/json",
			requestData:     "{\"Latitude\": 0.0, \"Longitude\": 0.0, \"maxDistance\": 1.0}",
			expectedCode:    http.StatusUnsupportedMediaType,
			expectedMessage: "Content-Type header is not application/json\n",
		},
		{ // Badly formatted json
			headerType:      requestHeader,
			requestData:     "{,\"Latitude\": 0.0, \"Longitude\": 0.0, \"maxDistance\": 1.0}", // comma at the front
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Request body contains badly-formed JSON (at position 2)\n",
		},
		{ // Typo in a keyword in request
			headerType:      requestHeader,
			requestData:     "{\"Latttitude\": 0.0, \"Longitude\": 0.0, \"maxDistance\": 1.0}",
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Request body contains unknown field \"Latttitude\"\n",
		},
		{ // Invalid data type in the json body
			headerType:      requestHeader,
			requestData:     "{\"latitude\": \"0.0\", \"longitude\": \"0.0\", \"maxDistance\": \"1.0\"}",
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Request body contains an invalid value for the \"latitude\" field (at position 18)\n",
		},
		{ // Extra unexpected field
			headerType:      requestHeader,
			requestData:     "{\"Latitude\": 0.0, \"Longitude\": 0.0, \"maxDistance\": 1.0, \"bonusValue\": 1.0}",
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Request body contains unknown field \"bonusValue\"\n",
		},
		{ // Empty request body
			headerType:      requestHeader,
			requestData:     "",
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Request body must not be empty\n",
		},
		{ // Request size exceeds 1 MB
			headerType:      requestHeader,
			requestData:     "{\"Latitude\": 10." + makeHeavyStringPayload(1048577) + "}",
			expectedCode:    http.StatusRequestEntityTooLarge,
			expectedMessage: "Request body must not be larger than 1MB\n",
		},
		{ // Missing required request field: Longitude
			headerType:      requestHeader,
			requestData:     "{\"latitude\": 1.0, \"maxDistance\": 1.0}",
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Missing required field(s) in request: Longitude\n",
		},
		{ // Passing empty request body with parenthesis
			headerType:      requestHeader,
			requestData:     "{}",
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Missing required field(s) in request: Longitude Latitude\n",
		},
		{ // Passing multiple json objects
			headerType:      requestHeader,
			requestData:     "{}, {}",
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Request body must only contain a single JSON object\n",
		},
	}

	for _, testCase := range testCasesParameters {
		request, err := http.NewRequest("POST", endpointPath, strings.NewReader(testCase.requestData))
		if err != nil {
			log.Fatal(err)
		}
		request.Header.Set("Content-Type", testCase.headerType)
		responseRecorder := httptest.NewRecorder()

		appRouter.ServeHTTP(responseRecorder, request)

		assert.Equal(t, testCase.expectedMessage, responseRecorder.Body.String())
		assert.Equal(t, testCase.expectedCode, responseRecorder.Code)
	}
}

// Utility function to create heavy strings (filled with zeros).
// Used to test whether the server will refuse too large requests.
func makeHeavyStringPayload(sizeInBytes int) string {
	var builder strings.Builder
	builder.Grow(sizeInBytes)
	for i := 0; i < sizeInBytes; i++ {
		builder.WriteString("0")
	}
	return builder.String()
}

func TestAppHandlingGoodRequests(t *testing.T) {
	const nothingFoundResponseBody = "{\"locations\":null}"
	const expectedResponseCode = http.StatusOK
	appRouter, endpointPath := setUpTestRouterAndEndpoint()

	testCaseParameters := []struct {
		requestData string
	}{
		{"{\"Latitude\":52.0989711, \"Longitude\": 21.2715719, \"maxDistance\": 5.1}"},
		{"{\"Latitude\":52.0989711, \"Longitude\": 21.2715719, \"maxDistance\": 0.0}"},
		{"{\"Latitude\":52.1101533, \"Longitude\": 21.2567803, \"maxDistance\": 3.0}"},
		{"{\"Latitude\":52.1031484, \"Longitude\": 21.2802653, \"maxDistance\": 7.0}"},
	}

	for _, testCase := range testCaseParameters {
		request, err := http.NewRequest("POST", endpointPath, strings.NewReader(testCase.requestData))
		if err != nil {
			log.Fatal(err)
		}
		request.Header.Set("Content-Type", requestHeader)
		responseRecorder := httptest.NewRecorder()

		appRouter.ServeHTTP(responseRecorder, request)

		assert.Equal(t, expectedResponseCode, responseRecorder.Code)
		assert.NotEqual(t, responseRecorder.Body.String(), nothingFoundResponseBody)
	}
}

func TestAppHandlingForbiddenRequests(t *testing.T) {
	const expectedResponseCode = http.StatusMethodNotAllowed
	const expectedMessageBody = ""
	appRouter, endpointPath := setUpTestRouterAndEndpoint()

	testCaseParameters := []struct {
		requestMethod string
	}{
		{"GET"},
		{"PUT"},
		{"DELETE"},
		{"UPDATE"},
	}

	for _, testCase := range testCaseParameters {
		request, err := http.NewRequest(testCase.requestMethod, endpointPath, nil)
		if err != nil {
			log.Fatal(err)
		}
		responseRecorder := httptest.NewRecorder()
		appRouter.ServeHTTP(responseRecorder, request)

		assert.Equal(t, expectedResponseCode, responseRecorder.Code)
		assert.Equal(t, expectedMessageBody, responseRecorder.Body.String())
	}
}
