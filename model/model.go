// Package containing various dataclasses used inside this project.
package model

// Coordinates, describing the location of an object.
type Coordinates struct {
	Latitude  float32
	Longitude float32
}

// Single data field as present in the database.
// This is passed (in form of a map) via DatabaseConnector.ProvideData() method.
type SingleLocationData struct {
	Latitude    float32
	Longitude   float32
	Name        string
	Description string
}

// Struct describing single location data as returned by the location.LocationsProvider. This is the data that
// (in form of a slice) will be passed by the server to the client.
type Location struct {
	ObjectId    int32
	Name        string
	Description string
	Distance    float32
	Position    Coordinates
}
