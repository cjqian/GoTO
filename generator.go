/*
 *This program takes in database creds and handles views.
 *If "-n" argument is added, creates a new view.
 *If "-d" argument is added, deletes all views.
 */
package main

import (
	"./sqlParser"
	"./structGenerator"
	"bufio"
	"fmt"
	"os"
)

var (
	//username creds
	username    = os.Args[1]
	password    = os.Args[2]
	environment = os.Args[3]
	db          = sqlParser.ConnectToDatabase(username, password, environment)

	//"n" for new, "d" for delete, else just generate
	genType = os.Args[4]
)

func CreateView() {
	fmt.Println("View Name:")
	var viewName string
	_, err := fmt.Scanln(&viewName)
	if err != nil {
		panic(err)
	}

	fmt.Println("Query:")
	in := bufio.NewReader(os.Stdin)
	query, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	sqlParser.MakeView(viewName, query)
	structGenerator.AppendToStructFiles(viewName)
}

func DeleteView() {
	sqlParser.DeleteViews()
	structGenerator.InitStructFiles()
}

func main() {
	//connect to database
	sqlParser.ConnectToDatabase(os.Args[1], os.Args[2], os.Args[3])

	//handles cases
	if genType == "n" {
		CreateView()
	} else if genType == "d" {
		DeleteView()
	} else {
		structGenerator.InitStructFiles()
	}
}
