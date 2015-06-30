package sqlParser

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ConnectToDatabase(username string, password string, environment string) sqlx.DB {
	db, err := sqlx.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+environment)
	check(err)

	return *db
}

//returns array of table names from queried database
func GetTableNames(db sqlx.DB) []string {
	var tableNames []string

	tableRawBytes := make([]byte, 1)
	tableInterface := make([]interface{}, 1)

	tableInterface[0] = &tableRawBytes

	rows, err := db.Query("SELECT DISTINCT TABLE_NAME FROM information_schema.tables WHERE TABLE_TYPE='BASE TABLE'")
	check(err)

	for rows.Next() {
		err := rows.Scan(tableInterface...)
		check(err)

		tableNames = append(tableNames, string(tableRawBytes))
	}

	return tableNames
}

//returns *Rows from queried database
func GetRows(db sqlx.DB, tableName string) *sqlx.Rows {
	rows, err := db.Queryx("SELECT * from " + tableName)
	check(err)

	return rows
}

//returns array of column names from table in database
//returns array of table names from queried database
func GetColumnNames(db sqlx.DB, tableName string) []string {
	var colNames []string

	colRawBytes := make([]byte, 1)
	colInterface := make([]interface{}, 1)

	colInterface[0] = &colRawBytes

	rows, err := db.Query("SELECT DISTINCT COLUMN_NAME FROM information_schema.columns WHERE TABLE_NAME='" + tableName + "'")
	check(err)

	for rows.Next() {
		err := rows.Scan(colInterface...)
		check(err)

		colNames = append(colNames, string(colRawBytes))
	}

	return colNames
}
