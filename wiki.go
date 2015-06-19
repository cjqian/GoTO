package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

type jsonFile struct {
	Body string
}

//generates JSON-styled output string
func makeJson() *jsonFile {
	body := "hello, world"
	// else return page with no error
	return &jsonFile{Body: body}
}

//when generate putton is pressed, JSON string is outputted
func generateHandler(w http.ResponseWriter, r *http.Request) {
	file := makeJson()
	renderTemplate(w, "index", file)
}

//serves regular homepage
func homeHandler(w http.ResponseWriter, r *http.Request) {
	body := ""
	file := &jsonFile{Body: body}
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
