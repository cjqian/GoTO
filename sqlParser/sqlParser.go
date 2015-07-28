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
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

var (
	globalEnvironment = ""
	globalDB          sqlx.DB
)

func check(e error) {
	if e != nil {
		fmt.Printf("Exiting: %v \n", e)
		os.Exit(1)
	}
}

func AddTableToDatabase(newCol interface{}, tableName string) {
	m := newCol.(map[string]interface{})
	query := "INSERT INTO " + tableName + " ("
	keyStr := ""
	valueStr := ""

	for k, v := range m {
		keyStr += k + ","
		valueStr += "'" + TypeToString(v) + "',"
	}

	keyStr = keyStr[:len(keyStr)-1]
	valueStr = valueStr[:len(valueStr)-1]

	query += keyStr + ") VALUES ( " + valueStr + " );"
	_, err := globalDB.Query(query)
	check(err)
}

func AddTablesToDatabase(newCols []interface{}, tableName string) {
	for _, col := range newCols {
		AddTableToDatabase(col, tableName)
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

//deletes given parameters
func DeleteFromTable(tableName string, parameters map[string]string) {
	query := "delete from " + tableName

	if len(parameters) > 0 {
		query += " where "

		for k, v := range parameters {
			query += k + "=" + v + " and "
		}

		query = query[:len(query)-4]
	}

	_, err := globalDB.Query(query)
	check(err)

	fmt.Println(query)
}

func UpdateTable(tableName string, parameters map[string]string, updateParameters map[string]string) {
	query := "update " + tableName

	if len(updateParameters) > 0 {
		query += " set "

		for k, v := range updateParameters {
			query += k + "='" + v + "', "
		}

		query = query[:len(query)-2]
	}

	if len(parameters) > 0 {
		query += " where "

		for k, v := range parameters {
			query += k + "='" + v + "' and "
		}

		query = query[:len(query)-4]
	}
	//_, err := globalDB.Query(query)
	//check(err)
	_, err := globalDB.Query(query)
	check(err)

	fmt.Println(query)

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

func GetViewNames() []string {
	var tableNames []string

	tableRawBytes := make([]byte, 1)
	tableInterface := make([]interface{}, 1)

	tableInterface[0] = &tableRawBytes

	rows, err := globalDB.Query("SELECT TABLE_NAME FROM information_schema.views")
	check(err)

	for rows.Next() {
		err := rows.Scan(tableInterface...)
		check(err)

		tableNames = append(tableNames, string(tableRawBytes))
	}

	return tableNames

}

//returns interface from given table (name) from queried database
func GetRowArray(tableName string) []map[string]interface{} {
	rows, err := globalDB.Queryx("SELECT * from " + tableName)
	check(err)

	rowArray := make([]map[string]interface{}, 0)

	for rows.Next() {
		results := make(map[string]interface{}, 0)
		err = rows.MapScan(results)

		for k, v := range results {
			if b, ok := v.([]byte); ok {
				results[k] = string(b)
			}
		}

		rowArray = append(rowArray, results)
	}
	return rowArray
}

//returns interface from given table (name) from queried database
func GetCustomRowArray(query string) []map[string]interface{} {
	rows, err := globalDB.Queryx(query)
	check(err)

	rowArray := make([]map[string]interface{}, 0)

	for rows.Next() {
		results := make(map[string]interface{}, 0)
		err = rows.MapScan(results)

		for k, v := range results {
			if b, ok := v.([]byte); ok {
				results[k] = string(b)
			}
		}
		rowArray = append(rowArray, results)
	}
	return rowArray
}

//returns array of column names from table in database
func GetCustomColumnNames(query string) []string {
	var colNames []string

	colRawBytes := make([]byte, 1)
	colInterface := make([]interface{}, 1)

	colInterface[0] = &colRawBytes

	rows, err := globalDB.Query(query)
	check(err)

	for rows.Next() {
		err := rows.Scan(colInterface...)
		check(err)

		colNames = append(colNames, string(colRawBytes))
	}

	return colNames
}

//returns array of column names from table in database
func GetTableColumnNames(tableName string) []string {
	var colNames []string

	colRawBytes := make([]byte, 1)
	colInterface := make([]interface{}, 1)

	colInterface[0] = &colRawBytes

	rows, err := globalDB.Query("SELECT COLUMN_NAME FROM information_schema.columns WHERE TABLE_NAME='" + tableName + "' and TABLE_SCHEMA='" + globalEnvironment + "'")
	check(err)

	for rows.Next() {
		err := rows.Scan(colInterface...)
		check(err)

		colName := tableName + "." + string(colRawBytes)
		colNames = append(colNames, colName)
	}

	return colNames
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
//MUST be of type table.column!!
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

func MakeView(queryName string, query string) {
	qStr := "create view " + queryName + " as " + query
	_, err := globalDB.Query(qStr)
	check(err)
}

func DeleteView(viewName string) {
	qStr := "drop view " + viewName
	_, err := globalDB.Query(qStr)
	check(err)
}

func DeleteViews() {
	for _, view := range GetViewNames() {
		fmt.Println(view)
		DeleteView(view)
	}
}
