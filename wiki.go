package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
)

var (
	addr   = flag.Bool("addr", false, "find open address and print to final-port.txt")
	tables = getSqlTables()
)

type jsonFile struct {
	Body   string
	Tables []string
}

//generates JSON-styled output string
func makeJson() *jsonFile {
	body := fmtJson(getSql())
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
	db, err := sql.Open("mysql", "to_user:twelve@tcp(localhost:3306)/to_development")
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

//pulls sql and returns string
func getSql() string {
	sqlString := ""
	//opens database
	db, err := sql.Open("mysql", "to_user:twelve@tcp(localhost:3306)/to_development")
	if err != nil {
		panic(err.Error())
	}
	//gets rows and columns
	table := "deliveryservice"
	rows, err := db.Query("SELECT * FROM " + table)
	if err != nil {
		panic(err.Error())
	}

	cols, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	//slice string
	rawResult := make([][]byte, len(cols))
	var result string
	//interface slice
	dest := make([]interface{}, len(cols))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		if err := rows.Scan(dest...); err != nil {
			log.Fatal(err)
		}

		for _, raw := range rawResult {
			if raw == nil {
				result += "\\N"
			} else {
				result += string(raw)
			}
			result += "\t"
		}

		sqlString += result + "\n" + "\n"
	}

	return sqlString
}

//when generate putton is pressed, JSON string is outputted
func generateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Print("%+v", r)
	tableName := r.FormValue("tableName")
	fmt.Println(tableName)
	file := makeJson()
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
