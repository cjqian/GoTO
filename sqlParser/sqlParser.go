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
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//connects to and returns a pointer to the database
func ConnectToDatabase(username string, password string, environment string) sqlx.DB {
	db, err := sqlx.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+environment)
	check(err)

	//set globalEnvironment
	globalEnvironment = environment

	return *db
}

//returns array of table name strings from queried database
func GetTableNames(db sqlx.DB) []string {
	var tableNames []string

	tableRawBytes := make([]byte, 1)
	tableInterface := make([]interface{}, 1)

	tableInterface[0] = &tableRawBytes

	rows, err := db.Query("SELECT TABLE_NAME FROM information_schema.tables WHERE TABLE_TYPE='BASE TABLE' and TABLE_SCHEMA='" + globalEnvironment + "'")
	check(err)

	for rows.Next() {
		err := rows.Scan(tableInterface...)
		check(err)

		tableNames = append(tableNames, string(tableRawBytes))
	}

	return tableNames
}

//returns *Rows from given table (name) from queried database
func GetRows(db sqlx.DB, tableName string) *sqlx.Rows {
	rows, err := db.Queryx("SELECT * from " + tableName)
	check(err)

	return rows
}

//returns array of column names from table in database
func GetColumnNames(db sqlx.DB, tableName string) []string {
	var colNames []string

	colRawBytes := make([]byte, 1)
	colInterface := make([]interface{}, 1)

	colInterface[0] = &colRawBytes

	rows, err := db.Query("SELECT COLUMN_NAME FROM information_schema.columns WHERE TABLE_NAME='" + tableName + "' and TABLE_SCHEMA='" + globalEnvironment + "'")
	check(err)

	for rows.Next() {
		err := rows.Scan(colInterface...)
		check(err)

		colNames = append(colNames, string(colRawBytes))
	}

	return colNames
}

//returns array of column names from table in database
func GetColumnTypes(db sqlx.DB, tableName string) []string {
	var colTypes []string

	colRawBytes := make([]byte, 1)
	colInterface := make([]interface{}, 1)

	colInterface[0] = &colRawBytes

	rows, err := db.Query("SELECT COLUMN_TYPE FROM information_schema.columns WHERE TABLE_NAME='" + tableName + "' and TABLE_SCHEMA='" + globalEnvironment + "'")
	check(err)

	for rows.Next() {
		err := rows.Scan(colInterface...)
		check(err)

		colTypes = append(colTypes, string(colRawBytes))
	}

	return colTypes
}
