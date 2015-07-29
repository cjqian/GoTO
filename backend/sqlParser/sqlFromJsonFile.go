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
func AddJsonRow(tableName string, fileName string) {
	//reads in the file
	fileStr, err := ioutil.ReadFile(fileName)
	if err != nil {
		outputError := errors.New("SP-AJR: File not found: " + fileName)
		panic(outputError)
	}

	//unmarshals the json into an interface
	var f []interface{}
	err2 := json.Unmarshal(fileStr, &f)
	if err2 != nil {
		outputError := errors.New("SP-AJR: Incorrect JSON formatting: " + fileName)
		panic(outputError)
	}

	//adds the interface row to table in database
	AddRowsToTable(f, tableName)
}

func PutJsonRow(tableName string, parameters []string, fileName string) {
	//reads in the file
	fileStr, err := ioutil.ReadFile(fileName)
	if err != nil {
		outputError := errors.New("SP-PJR: File not found: " + fileName)
		panic(outputError)
	}

	//unmarshals the json into an interface
	var f interface{}
	err2 := json.Unmarshal(fileStr, &f)
	if err2 != nil {
		outputError := errors.New("SP-PJR: Incorrect JSON formatting: " + fileName)
		panic(outputError)
	}

	//adds the interface row to table in database
	UpdateRow(f, tableName, parameters)
}
