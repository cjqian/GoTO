package main

import (
	"./sqlParser"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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

	viewName = strings.ToLower(viewName)
	sqlParser.MakeView(viewName, query)
	//if strings.Contains(query, " join ") {
	//structGenerator.JoinAppendToStructFiles(viewName)
	//} else {
	//structGenerator.AppendToStructFiles(viewName)
	//}
}

func main() {
	//connect to database
	sqlParser.ConnectToDatabase(os.Args[1], os.Args[2], os.Args[3])

	//get command
	response := os.Args[4]

	if response == "1" {
		CreateView()
	} else if response == "0" {
		sqlParser.DeleteViews()
	} else {
		err := errors.New("Incorrect response.")
		panic(err)
	}
}
