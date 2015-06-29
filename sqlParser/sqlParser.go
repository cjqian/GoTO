package sqlParser

import (
	"database/sql"
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
