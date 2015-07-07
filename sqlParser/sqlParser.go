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

//sqlParser.go
//interacts with database
package sqlParser

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	globalEnvironment = ""
	globalDB          sqlx.DB
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//connects to and returns a pointer to the database
func ConnectToDatabase(username string, password string, environment string) sqlx.DB {
	db, err := sqlx.Connect("mysql", username+":"+password+"@tcp(localhost:3306)/"+environment)
	check(err)

	//set globalEnvironment
	globalEnvironment = environment
	globalDB = *db

	return *db
}

//returns array of table name strings from queried database
func GetTableNames() []string {
	var tableNames []string

	tableRawBytes := make([]byte, 1)
	tableInterface := make([]interface{}, 1)

	tableInterface[0] = &tableRawBytes

	rows, err := globalDB.Query("SELECT TABLE_NAME FROM information_schema.tables WHERE TABLE_TYPE='BASE TABLE' and TABLE_SCHEMA='" + globalEnvironment + "'")
	check(err)

	for rows.Next() {
		err := rows.Scan(tableInterface...)
		check(err)

		tableNames = append(tableNames, string(tableRawBytes))
	}

	return tableNames
}

//returns *Rows from given table (name) from queried database
func GetRows(tableName string, request string) *sqlx.Rows {
	if request == "" {
		rows, err := globalDB.Queryx("SELECT * from " + tableName)
		check(err)
		return rows
	} else {
		rows, err := globalDB.Queryx("SELECT " + request + " from " + tableName)
		check(err)
		return rows
	}

	//this should never happen
	return nil
}

//returns array of column names from table in database
func GetColumnNames(tableName string) []string {
	var colNames []string

	colRawBytes := make([]byte, 1)
	colInterface := make([]interface{}, 1)

	colInterface[0] = &colRawBytes

	rows, err := globalDB.Query("SELECT COLUMN_NAME FROM information_schema.columns WHERE TABLE_NAME='" + tableName + "' and TABLE_SCHEMA='" + globalEnvironment + "'")
	check(err)

	for rows.Next() {
		err := rows.Scan(colInterface...)
		check(err)

		colNames = append(colNames, string(colRawBytes))
	}

	return colNames
}

//returns array of column names from table in database
func GetColumnTypes(tableName string) []string {
	var colTypes []string

	colRawBytes := make([]byte, 1)
	colInterface := make([]interface{}, 1)

	colInterface[0] = &colRawBytes

	rows, err := globalDB.Query("SELECT COLUMN_TYPE FROM information_schema.columns WHERE TABLE_NAME='" + tableName + "' and TABLE_SCHEMA='" + globalEnvironment + "'")
	check(err)

	for rows.Next() {
		err := rows.Scan(colInterface...)
		check(err)

		colTypes = append(colTypes, string(colRawBytes))
	}

	return colTypes
}

//returns array of column names from table in database
func GetColumnType(columnName string) string {
	var colTypes []string

	colRawBytes := make([]byte, 1)
	colInterface := make([]interface{}, 1)

	colInterface[0] = &colRawBytes

	rows, err := globalDB.Query("SELECT COLUMN_TYPE FROM information_schema.columns WHERE COLUMN_NAME='" + columnName + "'")
	check(err)

	for rows.Next() {
		err := rows.Scan(colInterface...)
		check(err)

		colTypes = append(colTypes, string(colRawBytes))
	}

	return colTypes[0]
}
