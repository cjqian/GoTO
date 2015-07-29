//package for POSTing JSON files to the database
package sqlParser

import (
	"encoding/json"
	"errors"
	//	"fmt"
	"io/ioutil"
)

//takes in the fileName of a file that contains row information in JSON form
//and the name of the table to which the row should be added
func AddJsonCols(tableName string, fileName string) {
	//reads in the file
	fileStr, err := ioutil.ReadFile(fileName)
	if err != nil {
		outputError := errors.New("SP-AJC: File not found: " + fileName)
		panic(outputError)
	}

	//unmarshals the json into an interface
	var f []interface{}
	err2 := json.Unmarshal(fileStr, &f)
	if err2 != nil {
		outputError := errors.New("SP-AJC: Incorrect JSON formatting: " + fileName)
		panic(outputError)
	}

	//adds the interface row to table in database
	AddTablesToDatabase(f, tableName)
}
