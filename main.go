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
	"./jsonParser"
	"./sqlParser"
	"./structCustom"
	"./structs"
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
	db          = sqlParser.ConnectToDatabase(username, password, environment)
)

//returns JSON of argument table name in database
func generateHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	request := urlParser.ParseURL(path)

	if r.Method == "GET" {
		tableName := request.TableName
		tableID := request.Id

		//checks table name
		if !structs.ValidStruct[tableName] && tableName != "custom" {
			http.NotFound(w, r)
		} else if tableName == "custom" {
			if structCustom.ValidCustomStruct(tableID) {
				queryStr, err := ioutil.ReadFile("structDirectory/queries/queryCustom_" + request.Id)
				fmt.Printf("Reading query:\n%s\n ", queryStr)
				if err != nil {
					panic(err)
				}
				rows := sqlParser.GetCustomRows(string(queryStr))
				w.Header().Set("Content-Type", "application/json")
				structCustom.MapCustomTableToJson(request.Id, rows, w)
			} else {
				http.NotFound(w, r)
			}
		} else {
			rows := sqlParser.GetRows(tableName)
			w.Header().Set("Content-Type", "application/json")

			structs.MapTableToJson(tableName, rows, w)

		}
	} else if r.Method == "POST" {
		filename := r.PostFormValue("filename")
		fmt.Println("Filename is: ", filename)
		jsonParser.AddJsonCols(request.TableName, filename)
	}
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
