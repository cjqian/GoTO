package jsonParser

import (
	"./../sqlParser"
	//"./structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
 *
 *program takes in arguments "tablename" and JSON file
 *var (
 *    username    = os.Args[1]
 *    password    = os.Args[2]
 *    environment = os.Args[3]
 *    tableName   = os.Args[4]
 *    filename    = os.Args[5]
 *    db          = sqlParser.ConnectToDatabase(username, password, environment)
 *)
 */

//adds to sql database
func check(err error) {
	if err != nil {
		panic(err)
	}
}

//reads file, interprets as json
func AddJsonCols(tableName string, filename string) {
	fmt.Println(filename)
	fileStr, err := ioutil.ReadFile(filename)
	fmt.Println(string(fileStr))
	check(err)

	var f []interface{}
	err2 := json.Unmarshal(fileStr, &f)
	check(err2)

	sqlParser.AddTablesToDatabase(f, tableName)
	fmt.Println("Added!")
}
