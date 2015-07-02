/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at
  http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

//structGenerator.go
//generates 'structs' package
package main

import (
	"./../sqlParser"
	"io/ioutil"
	"os"
	"strings"
)

var (
	username    = os.Args[1]
	password    = os.Args[2]
	environment = os.Args[3]
	db          = sqlParser.ConnectToDatabase(username, password, environment)
)

//error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//writes struct, interface, and map files to structs package
func main() {
	MakeStructs()
	MakeStructInterface()
	MakeStructMap()
	MakeStructValidMap()
}

//writes the struct file (structs.go), which has an object for each
//database table, ith each table field as a member variable
func MakeStructs() {
	structStr := "package structs\n"
	tableList := sqlParser.GetTableNames(db)

	//add a struct for each table
	for _, table := range tableList {
		structStr += "type " + strings.Title(table) + " struct {\n"
		columnList := sqlParser.GetColumnNames(db, table)
		for _, column := range columnList {
			//they are all string types, hope that's cool
			structStr += "\t" + strings.Title(column) + "\t\t" + "string\n"
		}

		structStr += "}\n"
	}

	//writes in relation to home directory
	WriteFile(structStr, "./../structs/structs.go")
}

//writes structInterface.go, which has a function that takes in *Rows,
//makes them into an array of []TableName structs, and encodes this
//array into JSON format
func MakeStructInterface() {
	//header, imports
	structInterface := "package structs\n"
	structInterface += "import (\n"
	structInterface += "\t\"github.com/jmoiron/sqlx\"\n"
	structInterface += "\t\"encoding/json\"\n"
	structInterface += "\t\"net/http\"\n"
	structInterface += ")\n"

	//makes a function for each object
	tableList := sqlParser.GetTableNames(db)
	for _, table := range tableList {
		//function declaration
		structInterface += "func EncodeStruct" + strings.Title(table) + "(rows *sqlx.Rows, w http.ResponseWriter) {\n" //make new array
		structInterface += "\tsa := make([]" + strings.Title(table) + ", 0)\n"
		//make new instance
		structInterface += "\tt := " + strings.Title(table) + "{}\n\n"
		//loops through all columns and translates to JSON
		structInterface += "\tfor rows.Next() {\n"
		structInterface += "\t\t rows.StructScan(&t)\n"
		structInterface += "\t\t sa = append(sa, t)\n"
		structInterface += "\t}\n\n"
		structInterface += "\tenc := json.NewEncoder(w)\n"
		structInterface += "\tenc.Encode(sa)\n"
		structInterface += "}\n"
	}

	//writes in relation to home directory
	WriteFile(structInterface, "./../structs/structInterface.go")
}

//maps each table in the database to the boolean "true,"
//used to confirm validity of URL
func MakeStructValidMap() {
	structValid := "package structs\n"

	structValid += "var ValidStruct = map[string]bool {\n"

	tableList := sqlParser.GetTableNames(db)
	for _, table := range tableList {
		structValid += "\t\"" + table + "\" : true,\n"
	}

	structValid += "}\n"

	WriteFile(structValid, "./../structs/structValidMap.go")
}

//writes structMap.go, which has a function that maps each tableName string to
//its respective function in structInterface.go
func MakeStructMap() {
	//declaration, imports
	structMap := "package structs\n"
	structMap += "import (\n\t\"github.com/jmoiron/sqlx\"\n"
	structMap += "\t\"net/http\"\n)\n"
	structMap += "func MapTableToJson(tableName string, rows *sqlx.Rows, w http.ResponseWriter) {\n"

	//each table has a case mapping name with structInterface function
	tableList := sqlParser.GetTableNames(db)
	for _, table := range tableList {
		structMap += "\tif tableName == \"" + table + "\"{\n"
		structMap += "\t\tEncodeStruct" + strings.Title(table) + "(rows, w)\n"
		structMap += "\t}\n"
	}

	//if invalid table, returns error
	structMap += "}\n"

	//writes in relation to home directory
	WriteFile(structMap, "./../structs/structMap.go")
}

//writes string str to fileName
func WriteFile(str string, fileName string) {
	strByte := []byte(str)
	err := ioutil.WriteFile(fileName, strByte, 0644)
	check(err)
}
