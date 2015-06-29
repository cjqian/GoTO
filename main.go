//takes in username, password, and database
//currently prints each instance in table to commnand line as
//json struct

package main

import (
	"./jsonFormatter"
	"./sqlParser"
	"./structConstructor"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	addr        = flag.Bool("addr", false, "find open address and print to final-port.txt")
	username    = os.Args[1]
	password    = os.Args[2]
	environment = os.Args[3]
)

//prints JSON of argument table name in database
func generateHandler(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Path[len("/"):]
	db := sqlParser.ConnectToDatabase(username, password, environment)
	rows := sqlParser.GetRows(db, tableName)
	fmt.Printf("%s", jsonFormatter.MakeJsonByteArray(rows))
	fmt.Printf("%v", sqlParser.GetTableNames(db))
	fmt.Printf("%v", sqlParser.GetColumnNames(db, "asn"))

	structConstructor.WriteFile(structConstructor.MakeString(db), "structs/structs.go")
}

func main() {
	flag.Parse()

	http.HandleFunc("/", generateHandler)

	if *addr {
		//runs on home
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}
	http.ListenAndServe(":8080", nil)
}
