//takes in username, password, and database
//currently prints each instance in table to commnand line as
//json struct

package main

import (
	"./sqlParser"
	"./structConstructor"
	"./structs"
	"flag"
	"fmt"
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
	db          = sqlParser.ConnectToDatabase(username, password, environment)
)

//prints JSON of argument table name in database
func generateHandler(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Path[len("/"):]
	//	tableName := "deliveryservice"
	fmt.Print(tableName)
	rows := sqlParser.GetRows(db, tableName)
	fmt.Printf("%s", structs.MapTableToJson(tableName, rows))
}

func main() {
	//make initial files
	structConstructor.MakeStructFiles(db)

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
