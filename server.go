/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at
  http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

//main.go
//takes in username, password, and database arguments
//runs server that handles url table searches

package main

import (
	//	"./outputFormatter"
	"./sqlParser"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	//	"net/url"
	"./urlParser"
	"os"
)

var (
	addr        = flag.Bool("addr", false, "find open address and print to final-port.txt")
	username    = os.Args[1]
	password    = os.Args[2]
	environment = os.Args[3]
	db          = sqlParser.InitializeDatabase(username, password, environment)
)

func getHandler(w http.ResponseWriter, tableName string, tableParameters []string) {
	rows := sqlParser.GetRows(tableName, tableParameters)
	enc := json.NewEncoder(w)
	enc.Encode(rows)
}

//returns JSON of argument table name in database
func viewHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if r.URL.RawQuery != "" {
		path += "?" + r.URL.RawQuery
	}

	request := urlParser.ParseURL(path)
	tableName := request.TableName
	tableParameters := request.Parameters

	if r.Method == "POST" {
		fileName := r.PostFormValue("filename")
		sqlParser.PostViews(fileName)

	} else if r.Method == "DELETE" {
		if tableName != "" {
			sqlParser.DeleteView(tableName)
		} else {
			sqlParser.DeleteViews()
		}
	}

	if tableName != "" && r.Method != "DELETE" {
		getHandler(w, tableName, tableParameters)
	}
}

//returns JSON of argument table name in database
func apiHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if r.URL.RawQuery != "" {
		path += "?" + r.URL.RawQuery
	}

	request := urlParser.ParseURL(path)
	tableName := request.TableName
	tableParameters := request.Parameters

	fmt.Println(request)
	if r.Method == "POST" {
		filename := r.PostFormValue("filename")
		sqlParser.AddRowsFromFile(tableName, filename)
	} else if r.Method == "DELETE" {
		sqlParser.DeleteFromTable(tableName, tableParameters)
	} else if r.Method == "PUT" {
		filename := r.PostFormValue("filename")
		sqlParser.PutJsonRow(tableName, tableParameters, filename)
	}

	if tableName != "" {
		getHandler(w, tableName, tableParameters)
	}
}
func main() {
	fmt.Println("Starting server.")
	flag.Parse()

	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/view/", viewHandler)

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