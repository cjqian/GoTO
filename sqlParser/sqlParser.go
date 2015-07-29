//interacts with database
package sqlParser

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	//	"os"
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

//adds new row to db
func AddTableToDatabase(newRow interface{}, tableName string) {
	m := newRow.(map[string]interface{})

	//"INSERT INTO tableName (a, b) values (av, bv);"
	query := "INSERT INTO " + tableName + " ("
	//parses through all keys and values
	keyStr := ""
	valueStr := ""
	for k, v := range m {
		keyStr += k + ","
		valueStr += "'" + TypeToString(v) + "',"
	}

	//removes last comma
	keyStr = keyStr[:len(keyStr)-1]
	valueStr = valueStr[:len(valueStr)-1]

	query += keyStr + ") VALUES ( " + valueStr + " );"

	_, err := globalDB.Query(query)
	if err != nil {
		outputError := errors.New("SP-ATTD: Invalid row for table " + tableName)
		panic(outputError)
	}
}

//add array of rows to db
func AddTablesToDatabase(newRows []interface{}, tableName string) {
	for _, row := range newRows {
		AddTableToDatabase(row, tableName)
	}

	fmt.Printf("SUCCESS: %d rows added to %s!\n", len(newRows), tableName)
}

//connects to and returns a pointer to the database
func ConnectToDatabase(username string, password string, environment string) sqlx.DB {
	db, err := sqlx.Connect("mysql", username+":"+password+"@tcp(localhost:3306)/"+environment)
	if err != nil {
		outputError := errors.New("SP-CTD: Could not connect to DB with creds")
		panic(outputError)
	}

	//set globalEnvironment
	globalEnvironment = environment
	globalDB = *db

	return *db
}

//deletes given parameters
func DeleteFromTable(tableName string, parameters map[string]string) {
	//delete from tableName where x = a and y = b
	query := "delete from " + tableName

	if len(parameters) > 0 {
		query += " where "

		for k, v := range parameters {
			query += k + "=" + v + " and "
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

func UpdateTable(tableName string, parameters map[string]string, updateParameters map[string]string) {
	query := "update " + tableName

	//new changes
	if len(updateParameters) > 0 {
		query += " set "

		for k, v := range updateParameters {
			query += k + "='" + v + "', "
		}

		query = query[:len(query)-2]
	}

	//where
	if len(parameters) > 0 {
		query += " where "

		for k, v := range parameters {
			query += k + "='" + v + "' and "
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

//returns *Rows from given table (name) from queried database
func GetRows(tableName string) *sqlx.Rows {
	rows, err := globalDB.Queryx("SELECT * from " + tableName)
	check(err)
	return rows
}

func GetCustomRows(query string) *sqlx.Rows {
	rows, err := globalDB.Queryx(query)
	check(err)
	return rows
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
