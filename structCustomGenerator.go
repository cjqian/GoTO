package main

import (
	"./sqlParser"
	"./structGenerator"
	"bufio"
	"errors"
	"fmt"
	"os"
)

func CreateCustomView() {
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

func main() {
	//connect to database
	sqlParser.ConnectToDatabase(os.Args[1], os.Args[2], os.Args[3])

	//get command
	fmt.Println("Create [0], Delete All [1] Views?")
	var response int
	_, err := fmt.Scanln(&response)
	if err != nil {
		panic(err)
	}

	if response == 0 {
		CreateCustomView()
	} else if response == 1 {
		sqlParser.DeleteViews()
		structGenerator.InitStructFiles()
	} else {
		err := errors.New("Incorrect response.")
		panic(err)
	}
}
