//takes in username, password, and devel

package main

import (
	"./structs"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
)

var (
	addr        = flag.Bool("addr", false, "find open address and print to final-port.txt")
	tables      = getSqlTables()
	username    = os.Args[1]
	password    = os.Args[2]
	environment = os.Args[3]
)

type jsonFile struct {
	Body   string
	Tables []string
}

//generates JSON-styled output string
func makeJson(dbName string) *jsonFile {
	body := fmtJson(getSql(dbName))
	// else return page with no error
	return &jsonFile{Body: body, Tables: tables}
}

//from a sql string, format json
func fmtJson(data string) string {
	return data
}

func getSqlTables() []string {
	var saTables []string

	//opens database
	db, err := sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+environment)
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT DISTINCT TABLE_NAME FROM information_schema.tables WHERE TABLE_TYPE='BASE TABLE'")
	if err != nil {
		panic(err.Error())
	}

	//slice string
	rawResult := make([]byte, 1)
	//interface slice
	dest := make([]interface{}, 1)

	dest[0] = &rawResult
	for rows.Next() {
		if err := rows.Scan(dest...); err != nil {
			log.Fatal(err)
		}

		result := string(rawResult)
		saTables = append(saTables, result)
	}

	return saTables
}

//pulls sql and returns sql object
func getSql(dbName string) string {
	//opens database
	db, err := sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+environment)
	if err != nil {
		panic(err.Error())
	}
	//asnObjs := structs.AsnStruct{12, 13, 12}

	rows, erra := db.Query("SELECT id, asn, cachegroup from asn")
	/*for _, each := range asnObjs {
		fmt.Printf("%#v\n", each)
	}*/
	if erra != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var a int
		var b int
		var c int
		erra2 := rows.Scan(&a, &b, &c)
		if erra2 != nil {
			log.Fatal(err)
		}
		obj := structs.AsnStruct{a, b, c}
		d, _ := json.MarshalIndent(obj, "", "  ")
		fmt.Printf("%s\n", d)

	}

	return ""
}

//when generate putton is pressed, JSON string is outputted
func generateHandler(w http.ResponseWriter, r *http.Request) {
	tableName := r.FormValue("tableName")

	file := makeJson(tableName)
	renderTemplate(w, "index", file)
}

//serves regular homepage
func homeHandler(w http.ResponseWriter, r *http.Request) {
	body := ""
	file := &jsonFile{Body: body, Tables: tables}
	renderTemplate(w, "index", file)
}

var templates = template.Must(template.ParseFiles("index.html"))

//error checks a template (index.html)
func renderTemplate(w http.ResponseWriter, tmpl string, p *jsonFile) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Parse/compile regext, panic if compilation fails (no err parameter)
var validPath = regexp.MustCompile("^/(generate|edit|save|view)/([a-zA-Z0-9]+)$")

func main() {
	flag.Parse()

	//two possible url cases
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/generate/", generateHandler)

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
