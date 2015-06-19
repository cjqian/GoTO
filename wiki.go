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
	Body []byte
}

/*
//saves a page
func (p *jsonFile) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
*/

//loads a page
func makeJson() *jsonFile {
	body := "hello, world"
	// else return page with no error
	return &jsonFile{Body: []byte(body)}
}

/*
//allows a user to view a wikipage
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := makeJson(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}*/
/*
//allows user to edit a wikipage
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := makeJson(title)
	if err != nil {
		p = &jsonFile{Title: title}
	}

	renderTemplate(w, "edit", p)
}
*/
func generateHandler(w http.ResponseWriter, r *http.Request) {
	file := makeJson()
	renderTemplate(w, "index", file)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	body := ""
	file := &jsonFile{Body: []byte(body)}
	renderTemplate(w, "index", file)
}

/*
//save pages
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &jsonFile{Title: title, Body: []byte(body)}
	err := p.save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}*/

var templates = template.Must(template.ParseFiles("index.html"))

//helper function
func renderTemplate(w http.ResponseWriter, tmpl string, p *jsonFile) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Parse/compile regext, panic if compilation fails (no err parameter)
var validPath = regexp.MustCompile("^/(generate|edit|save|view)/([a-zA-Z0-9]+)$")

/*
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the page title from the Request and call handler 'hn'
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
*/
func main() {
	flag.Parse()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/generate/", generateHandler)
	//	http.HandleFunc("/view/", makeHandler(viewHandler))
	//	http.HandleFunc("/edit/", makeHandler(editHandler))
	//	http.HandleFunc("/save/", makeHandler(saveHandler))

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
