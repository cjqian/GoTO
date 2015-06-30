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

//makes structs and struct interfaces
func MakeStructFiles(db sqlx.DB) {
	MakeStructs(db)
	MakeStructInterface(db)
	MakeStructToJson(db)
}

//makes structures from database
func MakeStructs(db sqlx.DB) {
	structStr := "package structs\n"
	tableList := sqlParser.GetTableNames(db)

	//add a struct for each table
	//writes in relation to home directory
	for _, table := range tableList {
		structStr += "type " + strings.Title(table) + " struct {\n"
		columnList := sqlParser.GetColumnNames(db, table)
		for _, column := range columnList {
			//they are all string types, hope that's cool
			structStr += "\t" + strings.Title(column) + "\t\t" + "string\n"
		}

		structStr += "}\n"
	}
	WriteFile(structStr, "./structs/structs.go")
}

//adds rows into a struct
func MakeStructInterface(db sqlx.DB) {
	structInterface := "package structs\n"
	structInterface += "import (\n"
	structInterface += "\t\"github.com/jmoiron/sqlx\"\n"
	structInterface += "\t\"encoding/json\"\n"
	structInterface += ")\n"
	tableList := sqlParser.GetTableNames(db)
	for _, table := range tableList {
		//function declaration
		structInterface += "func ByteArrayFrom" + strings.Title(table) + "(rows *sqlx.Rows)  []byte {\n"
		structInterface += "\tvar tStr []byte\n"
		//make new instance
		structInterface += "\tt := " + strings.Title(table) + "{}\n"
		//loops through all columns
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

//maps each string to a function
func MakeStructToJson(db sqlx.DB) {
	structMap := "package structs\n"
	structMap += "import \"github.com/jmoiron/sqlx\"\n"
	structMap += "func MapTableToJson(tableName string, rows *sqlx.Rows) []byte{\n"

	tableList := sqlParser.GetTableNames(db)
	for _, table := range tableList {
		structMap += "\tif tableName == \"" + table + "\"{\n"
		structMap += "\t\ttStr := ByteArrayFrom" + strings.Title(table) + "(rows)\n"
		structMap += "\t\treturn tStr\n"
		structMap += "\t}\n"
	}

	structMap += "return []byte(\"ERROR: TABLE NOT FOUND\")\n"
	structMap += "}\n"
	WriteFile(structMap, "./structs/structMap.go")
}

//writes string str to fileName
func WriteFile(str string, fileName string) {
	strByte := []byte(str)
	err := ioutil.WriteFile(fileName, strByte, 0644)
	check(err)
}
