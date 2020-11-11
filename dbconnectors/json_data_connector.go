// Contains package code for reading data from a JSON file.
package dbconnectors

import (
	"encoding/json"
	"github.com/sebastian-sz/GotwockAppServer/model"
	"io/ioutil"
	"log"
)

// Database connector responsible for reading data from a .json file.
// In order not to re-read the file for each ProvideData() call the data is loaded into the memory on the first
// json read via LoadContentToMemory() method. The data is stored in memory throughout the life of the program.
// JSONDataConnector should only be used for either debugging / prototyping purposes or when the database size is small
// enough that the data can be safely held in memory.
// It is advised to move onto proper database when the data size becomes too large.
// Fields:
//     DataPath: string. Defines the path to the .json file that contains data.
//     CachedData: map. Should be initialised with nil value as this will be overwritten by LoadContentToMemory()
//     method.
type JSONDataConnector struct {
	DataPath   string
	CachedData map[int]model.SingleDataField
}

// JSONDataConnector method for reading content of a .json file.
// Returns file content.
func (j *JSONDataConnector) readJson() []byte {
	file, err := ioutil.ReadFile(j.DataPath)
	if err != nil {
		log.Fatal("Error loading file: ", err)
	}
	return file
}

// JSONDataConnector method for reading the .json file and loading it's content to memory.
// The loaded content will be held under CachedData field of the JSONDataConnector struct.
// Note that this method should be called right after creating the JSONDataConnector struct.
func (j *JSONDataConnector) loadContentToMemory() {
	fileContent := j.readJson()
	err := json.Unmarshal(fileContent, &j.CachedData)
	if err != nil {
		log.Fatal("Error unmarshalling json: ", err)
	}
}

// Initialize the connection with a database.
// In case of JSONDataConnector that means loading the content in memory.
func (j *JSONDataConnector) Initialize() {
	j.loadContentToMemory()
}

// Overloaded interface method. It returns the map of object Id (int) and generic data (SingleDataField)
// In practice, for JSONDataConnector this method only returns cached content that is loaded after calling
// LoadContentToMemory().
func (j *JSONDataConnector) ProvideData() map[int]model.SingleDataField {
	return j.CachedData
}
