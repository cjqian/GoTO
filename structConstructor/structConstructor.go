//structConstructor.go
//generates 'structs' package
package structConstructor

import (
	"./../sqlParser"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"strings"
)

//error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//writes struct, interface, and map files to structs package
func MakeStructFiles(db sqlx.DB) {
	MakeStructs(db)
	MakeStructInterface(db)
	MakeStructMap(db)
}

//writes the struct file (structs.go), which has an object for each
//database table, ith each table field as a member variable
func MakeStructs(db sqlx.DB) {
	structStr := "package structs\n"
	tableList := sqlParser.GetTableNames(db)

	//add a struct for each table
	for _, table := range tableList {
		structStr += "type " + strings.Title(table) + " struct {\n"
		columnList := sqlParser.GetColumnNames(db, table)
		for _, column := range columnList {
			//they are all string types, hope that's cool
			structStr += "\t" + strings.Title(column) + "\t\t" + "string\n"
		}

		structStr += "}\n"
	}

	//writes in relation to home directory
	WriteFile(structStr, "./structs/structs.go")
}

//writes structInterface.go, which has a function that takes in *Rows and
//returns the byte array JSON format for each table in the database
func MakeStructInterface(db sqlx.DB) {
	//header, imports
	structInterface := "package structs\n"
	structInterface += "import (\n"
	structInterface += "\t\"github.com/jmoiron/sqlx\"\n"
	structInterface += "\t\"encoding/json\"\n"
	structInterface += ")\n"

	//makes a function for each object
	tableList := sqlParser.GetTableNames(db)
	for _, table := range tableList {
		//function declaration
		structInterface += "func ByteArrayFrom" + strings.Title(table) + "(rows *sqlx.Rows)  []byte {\n"
		structInterface += "\tvar tStr []byte\n"
		//make new instance
		structInterface += "\tt := " + strings.Title(table) + "{}\n"
		//loops through all columns and translates to JSON
		structInterface += "\tfor rows.Next() {\n"
		structInterface += "\t\t rows.StructScan(&t)\n"
		structInterface += "\t\t tmpStr, _ := json.MarshalIndent(t, \"\", \"  \")\n"
		structInterface += "\t\t tStr = append(tStr[:], tmpStr[:]...)\n"
		structInterface += "\t}\n\n"
		structInterface += "\treturn tStr\n"
		structInterface += "}\n"
	}

	//writes in relation to home directory
	WriteFile(structInterface, "./structs/structInterface.go")
}

//writes structMap.go, which has a function that maps each tableName string to
//its respective function in structInterface.go
func MakeStructMap(db sqlx.DB) {
	//declaration, imports
	structMap := "package structs\n"
	structMap += "import \"github.com/jmoiron/sqlx\"\n"
	structMap += "func MapTableToJson(tableName string, rows *sqlx.Rows) []byte{\n"

	//each table has a case mapping name with structInterface function
	tableList := sqlParser.GetTableNames(db)
	for _, table := range tableList {
		structMap += "\tif tableName == \"" + table + "\"{\n"
		structMap += "\t\ttStr := ByteArrayFrom" + strings.Title(table) + "(rows)\n"
		structMap += "\t\treturn tStr\n"
		structMap += "\t}\n"
	}

	//if invalid table, returns error
	structMap += "return []byte(\"ERROR: TABLE NOT FOUND\")\n"
	structMap += "}\n"

	//writes in relation to home directory
	WriteFile(structMap, "./structs/structMap.go")
}

//writes string str to fileName
func WriteFile(str string, fileName string) {
	strByte := []byte(str)
	err := ioutil.WriteFile(fileName, strByte, 0644)
	check(err)
}
