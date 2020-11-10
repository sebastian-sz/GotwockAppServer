// Package containing various dataclasses used inside this project.

package model

// Coordinates, describing the location of an object.
type Coordinates struct {
	Latitude  float32
	Longitude float32
}

// Single location data as present in the database.
// This is passed (in form of a map) via DatabaseConnector.ProvideData() method.
type SingleDataField struct {
	Latitude    float32
	Longitude   float32
	Name        string
	Description string
}

// Todo: add documentation.
type TouristLocation struct {
	ObjectId    int32
	Name        string
	Description string
	Distance    string
	Position    Coordinates
}
