package main

import (
	"fmt"
	"net/http"
)

//Assemples HTTP server's response, sends data there
func handler(w http.ResponseWriter, r *http.Request) {
	//1: removes the slash thing
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	//handle requests to web root ("/")
	http.HandleFunc("/", handler)

	//listen to this port, block until program is terminated
	http.ListenAndServe(":8080", nil)
}
