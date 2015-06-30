package jsonFormatter

import (
	"./../structs"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//returns *Rows from queried database
func GetRows(username string, password string, environment string, tableName string) *sql.Rows {
	//opens database
	db, err := sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+environment)
	check(err)

	rows, err := db.Query("SELECT * from asn")
	check(err)

	return rows
}

//from given rows, returns array of byte arrays
func MakeJsonByteArray(rows *sql.Rows) [][]byte {
	byteArray := make([][]byte, 0)
	for rows.Next() {
		//currently hardcoded to "asn"
		var a string
		var b string
		var c string
		var d string
		err := rows.Scan(&a, &b, &c, &d)
		check(err)

		tableInstance := structs.Asn{a, b, c, d}
		tableInstanceJson, _ := json.MarshalIndent(tableInstance, "", "  ")
		tableInstanceJson = append(tableInstanceJson, ","...)
		byteArray = append(byteArray, tableInstanceJson)
	}

	return byteArray
}
