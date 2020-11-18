// Main package responsible for running the server.
package main

import (
	"flag"
	"github.com/sebastian-sz/GotwockAppServer/appserver"
	"github.com/sebastian-sz/GotwockAppServer/factory"
	"strconv"
	"strings"
)

const endpointPath = "/"
const localhost = "127.0.0.1"

// Create localhost server address given pointer to integer, specifying port.
func createAddress(port *int) string {
	var finalAddress strings.Builder

	finalAddress.WriteString(localhost)
	finalAddress.WriteString(":")
	finalAddress.WriteString(strconv.Itoa(*port))

	return finalAddress.String()
}

// Parses user provided flag(s) and starts the server on the specified port.
func main() {
	port := flag.Int("port", 9100, "Port on which to start the server.")
	flag.Parse()

	serverAddress := createAddress(port)

	touristLocationProvider := factory.CreateJSONHaversineTouristLocationProvider()
	app := appserver.App{
		EndpointPath: endpointPath, ServerAddr: serverAddress, LocationsProvider: &touristLocationProvider,
	}

	app.Run()
}
