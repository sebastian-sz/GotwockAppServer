// Package for creating mux server to connect the entire application with the outside world.
// In more detail: here, we provide the logic of parsing requests, running location.LocationsProvider, and formatting
// a proper response.
package appserver

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/sebastian-sz/GotwockAppServer/location"
	"github.com/sebastian-sz/GotwockAppServer/model"
	"log"
	"net/http"
	"time"
)

// Utility struct for creating json out of location.LocationsProvider GetAndParseLocationsData() output.
type ResponseData struct {
	Locations []model.Location `json:"locations"`
}

// Responsible for handling incoming requests, data parsing and response writing.
type App struct {
	EndpointPath      string
	ServerAddr        string
	LocationsProvider *location.LocationsProvider
}

// Create router for the application
func (app *App) makeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(app.EndpointPath, app.handleRequest()).Methods("POST")
	return r
}

// Handle incoming request.
// There are multiple checks performed before the request data is passed further to the location.LocationsProvider.
// For the detailed list of all checks please see functions decodeJSONBody and parsedRequestData.CheckIfAllDataProvided.
// If all the above checks have been passed the data is processed by te location.LocationsProvider and response is
// returned.
func (app *App) handleRequest() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {

		var requestData parsedRequestData
		decodeError := decodeJSONBody(responseWriter, request, &requestData)

		if decodeError != nil {
			var malformedRequest *malformedRequestError
			if errors.As(decodeError, &malformedRequest) {
				http.Error(responseWriter, malformedRequest.message, malformedRequest.status)
			} else {
				log.Println(decodeError.Error())
				http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		missingValuesError := requestData.CheckIfAllDataProvided()
		if missingValuesError != nil {
			http.Error(responseWriter, missingValuesError.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Received and parsed request: %v\n", requestData)

		userCoordinates := model.Coordinates{
			Latitude:  *requestData.Latitude,
			Longitude: *requestData.Longitude,
		}

		nearestLocations := app.LocationsProvider.GetAndParseLocationsData(
			userCoordinates, *requestData.MaxDistance,
		)

		responseLocationData := ResponseData{
			Locations: nearestLocations,
		}

		payload, marshallingError := json.Marshal(&responseLocationData)
		if marshallingError != nil {
			log.Println(marshallingError.Error())
			http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		log.Printf("Sending %v found location(s)", len(responseLocationData.Locations))

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		_, writeErr := responseWriter.Write(payload)
		if writeErr != nil {
			log.Println(writeErr.Error())
		}
	}
}

// Run method of the struct App. It creates mux http router and starts the server.
func (app *App) Run() {
	router := app.makeRouter()

	log.Printf("Starting server at: http://%v\n", app.ServerAddr)

	srv := &http.Server{
		Handler:      router,
		Addr:         app.ServerAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// Todo: Consider graceful shutdown
	// https://github.com/gorilla/mux#graceful-shutdown
	log.Fatal(srv.ListenAndServe())
}
