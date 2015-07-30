//package main

package sqlParser

/**********************************************************************************
 * DIRECTORY:
 * 1. DB INITIALIZE
 * 2. MAINSERVER: DELETE
 * 3. MAINSERVER: GET
 * 4. MAINSERVER: POST
 * 5. MAINSERVER: PUT
 * 6. VIEWSERVER: POST
 * 7: VIEWSERVER: DELETE
 *********************************************************************************/

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"strings"
)

var (
	globalDB sqlx.DB
	colMap   map[string]string
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*********************************************************************************
 * DB INITIALIZE: Connects given DB creds, creates ColMap FOR SESSION
 ********************************************************************************/
func InitializeDatabase(username string, password string, environment string) sqlx.DB {
	db, err := sqlx.Connect("mysql", username+":"+password+"@tcp(localhost:3306)/"+environment)
	check(err)

	globalDB = *db

	//set global colMap
	colMap = GetColMap()
	return *db
}

//returns interface from given table (name) from queried database
func GetColMap() map[string]string {
	colMap := make(map[string]string, 0)
	cols, err := globalDB.Queryx("SELECT DISTINCT COLUMN_NAME, COLUMN_TYPE FROM information_schema.columns")
	check(err)

	for cols.Next() {
		var colName string
		var colType string

		err = cols.Scan(&colName, &colType)

		colMap[colName] = strings.Split(colType, "(")[0]
	}

	return colMap
}

/*********************************************************************************
 * MAINSERVER: DELETE FUNCTIONALITY
 ********************************************************************************/
//deletes given parameters
func DeleteFromTable(tableName string, parameters []string) {
	//delete from tableName where x = a and y = b
	query := "delete from " + tableName

	if len(parameters) > 0 {
		query += " where "

		for _, v := range parameters {
			query += v + " and "
		}
		//removes last "and"
		query = query[:len(query)-4]
	}

	_, err := globalDB.Query(query)
	if err != nil {
		outputError := errors.New("SP-DFT: Invalid delete query: " + query)
		panic(outputError)
	}

	fmt.Printf("SUCCESS: %s\n", query)
}

/*********************************************************************************
 * MAINSERVER: GET FUNCTIONALITY
 ********************************************************************************/
//returns interface from given table (name) from queried database
func GetRows(tableName string, tableParams []string) []map[string]interface{} {

	whereStmt := ""
	if len(tableParams) > 0 {
		whereStmt += " where "

		for _, v := range tableParams {
			whereStmt += v + " and "
		}

		whereStmt = whereStmt[:len(whereStmt)-4]
	}
	rows, err := globalDB.Queryx("SELECT * from " + tableName + whereStmt)
	check(err)

	rowArray := make([]map[string]interface{}, 0)

	for rows.Next() {
		results := make(map[string]interface{}, 0)
		err = rows.MapScan(results)
		for k, v := range results {
			if b, ok := v.([]byte); ok {
				results[k] = StringToType(b, colMap[k])
			}
		}

		rowArray = append(rowArray, results)
	}
	return rowArray
}

/*********************************************************************************
 * MAINSERVER: POST FUNCTIONALITY
 ********************************************************************************/
//adds new row to table
func AddRowToTable(newRow interface{}, tableName string) {
	m := newRow.(map[string]interface{})
	//insert into table (colA, colB) values (valA, valB);
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
	fmt.Println(query)
	_, err := globalDB.Query(query)
	check(err)
}

func AddRowsToTable(newRows []interface{}, tableName string) {
	for _, row := range newRows {
		AddRowToTable(row, tableName)
	}
}

//adds JSON from FILENAME to TABLE
func AddRowsFromFile(tableName string, fileName string) {
	fileStr, err := ioutil.ReadFile(fileName)
	check(err)

	var f []interface{}
	err2 := json.Unmarshal(fileStr, &f)
	check(err2)

	AddRowsToTable(f, tableName)
	fmt.Println("SP-ADDROWS: SUCCESS")
}

/*********************************************************************************
 * MAINSERVER: PUT FUNCTIONALITY
 ********************************************************************************/
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

func UpdateRow(newRow interface{}, tableName string, parameters []string) {
	query := "update " + tableName

	updateParameters := newRow.(map[string]interface{})
	//new changes
	if len(updateParameters) > 0 {
		query += " set "

		for k, v := range updateParameters {
			query += k + "='" + TypeToString(v) + "', "
		}

		query = query[:len(query)-2]
	}

	//where
	if len(parameters) > 0 {
		query += " where "

		for _, v := range parameters {
			query += v + " and "
		}

		query = query[:len(query)-4]
	}

	_, err := globalDB.Query(query)
	if err != nil {
		outputError := errors.New("SP-UT: Invalid update query: " + query)
		panic(outputError)
	}

	fmt.Printf("SUCCESS: %s\n", query)
}

/*********************************************************************************
 * VIEWSERVER: DELETE FUNCTIONALITY
 ********************************************************************************/

func DeleteView(viewName string) {
	qStr := "drop view " + viewName
	_, err := globalDB.Query(qStr)
	check(err)
}

func DeleteViews() {
	for _, view := range GetViewNames() {
		DeleteView(view)
	}

	fmt.Printf("SUCCESS: DELETED VIEWS")
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

/*********************************************************************************
 * VIEWSERVER: POST FUNCTIONALITY
 ********************************************************************************/
type View struct {
	Name  string
	Query string
}

//adds JSON from FILENAME to TABLE
func PostViews(fileName string) {
	fileStr, err := ioutil.ReadFile(fileName)
	check(err)

	var views []View
	err2 := json.Unmarshal(fileStr, &views)
	check(err2)

	for _, view := range views {
		MakeView(view.Name, view.Query)
	}

	fmt.Println("SUCCESS: VIEWS MADE")
}

func MakeView(viewName string, view string) {
	qStr := "create view " + viewName + " as " + view
	_, err := globalDB.Query(qStr)
	check(err)
}
