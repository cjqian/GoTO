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

package main

import (
	"./outputFormatter"
	"./sqlParser"
	"./urlParser"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

var (
	addr     = flag.Bool("addr", false, "find open address and print to final-port.txt")
	username = os.Args[1]
	password = os.Args[2]
	database = os.Args[3]

	//initializing the database connects and writes a column type map
	//(see sqlParser for more details)
	db = sqlParser.InitializeDatabase(username, password, database)
)

//handles all calls to the API
func apiHandler(w http.ResponseWriter, r *http.Request) {
	//url of type "/table?parameterA=valueA&parameterB=valueB/id
	path := r.URL.Path[1:]
	if r.URL.RawQuery != "" {
		path += "?" + r.URL.RawQuery
	}

	request := urlParser.ParseURL(path)

	//note: tableName could also refer to a view
	tableName := request.TableName
	tableParameters := request.Parameters

	if r.Method == "POST" {
		fileName := r.PostFormValue("filename")
		sqlParser.Post(tableName, fileName)
	} else if r.Method == "DELETE" {
		sqlParser.Delete(tableName, tableParameters)
	} else if r.Method == "PUT" {
		fileName := r.PostFormValue("filename")
		sqlParser.Put(tableName, tableParameters, fileName)
	}

	//GETS the request
	if tableName != "" {
		rows := sqlParser.Get(tableName, tableParameters)
		resp := outputFormatter.MakeWrapper(rows)

		//encoder writes the resultant "Response" struct (see outputFormatter) to writer
		enc := json.NewEncoder(w)
		enc.Encode(resp)
	}
}

func main() {
	fmt.Println("Starting server.")
	flag.Parse()

	http.HandleFunc("/", apiHandler)

	if *addr {
		//runs on home
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			panic(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}

	http.ListenAndServe(":8080", nil)
}
