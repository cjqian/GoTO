//generates 'structs' package
package structGenerator

import (
	"./../sqlParser"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

//error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func AppendToStructFiles(table string) {
	AppendToStructs(table)
	AppendToStructInterface(table)
	AppendToStructMap(table)
	AppendToStructValidMap(table)
}

//writes struct, interface, and map files to structs package
func InitStructFiles() {
	InitStructs()
	InitStructInterface()
	InitStructMap()
	InitStructValidMap()
}

func AppendToStructs(table string) {
	columnList := sqlParser.GetColumnNames(table)
	columnTypes := sqlParser.GetColumnTypes(table)

	structStr := MakeStructStr(table, columnList, columnTypes)

	AddToFile("./genStructs/structs.go", structStr)
}

//writes the struct file (structs.go), which has an object for each
//database table, ith each table field as a member variable
func InitStructs() {
	structStr := "package genStructs\n"

	tableList := sqlParser.GetTableNames()
	tableList = append(tableList, sqlParser.GetViewNames()...)

	//add a struct for each table
	for _, table := range tableList {
		columnList := sqlParser.GetColumnNames(table)
		columnTypes := sqlParser.GetColumnTypes(table)

		structStr += MakeStructStr(table, columnList, columnTypes)
	}

	//writes in relation to home directory
	WriteFile(structStr, "./genStructs/structs.go")
}

func AppendToStructInterface(table string) {
	//function declaration
	structInterface := "func EncodeStruct" + strings.Title(table) + "(rows *sqlx.Rows) interface{} {\n" //make new array
	structInterface += "\tsa := make([]" + strings.Title(table) + ", 0)\n"
	//make new instance
	structInterface += "\tt := " + strings.Title(table) + "{}\n\n"
	//loops through all columns and translates to JSON
	structInterface += "\tfor rows.Next() {\n"
	structInterface += "\t\t rows.StructScan(&t)\n"
	structInterface += "\t\t sa = append(sa, t)\n"
	structInterface += "\t}\n\n"
	structInterface += "\treturn sa\n"
	structInterface += "}\n"

	AddToFile("./genStructs/structInterface.go", structInterface)
}

//writes structInterface.go, which has a function that takes in *Rows,
//makes them into an array of []TableName structs, and encodes this
//array into JSON format
func InitStructInterface() {
	//header, imports
	structInterface := "package genStructs\n"
	structInterface += "import (\n"
	structInterface += "\t\"github.com/jmoiron/sqlx\"\n"
	structInterface += ")\n"

	//makes a function for each object
	tableList := sqlParser.GetTableNames()
	tableList = append(tableList, sqlParser.GetViewNames()...)
	for _, table := range tableList {
		//function declaration
		structInterface += "func EncodeStruct" + strings.Title(table) + "(rows *sqlx.Rows) interface{} {\n" //make new array
		structInterface += "\tsa := make([]" + strings.Title(table) + ", 0)\n"
		//make new instance
		structInterface += "\tt := " + strings.Title(table) + "{}\n\n"
		//loops through all columns and translates to JSON
		structInterface += "\tfor rows.Next() {\n"
		structInterface += "\t\t rows.StructScan(&t)\n"
		structInterface += "\t\t sa = append(sa, t)\n"
		structInterface += "\t}\n\n"
		structInterface += "\treturn sa\n"
		structInterface += "}\n"
	}

	//writes in relation to home directory
	WriteFile(structInterface, "./genStructs/structInterface.go")
}

func AppendToStructValidMap(table string) {

	structValid := "\t\"" + table + "\" : true,\n"

	AddToMethodInFile("./genStructs/structValidMap.go", structValid)
}

//writes structValidMap.go, which maps each table in the database to the boolean "true,"
//used to confirm validity of URL
func InitStructValidMap() {
	structValid := "package genStructs\n"

	structValid += "var ValidStruct = map[string]bool {\n"

	tableList := sqlParser.GetTableNames()
	tableList = append(tableList, sqlParser.GetViewNames()...)
	for _, table := range tableList {
		structValid += "\t\"" + table + "\" : true,\n"
	}

	structValid += "}\n"

	WriteFile(structValid, "./genStructs/structValidMap.go")
}

func AppendToStructMap(table string) {
	structMap := "\tif tableName == \"" + table + "\"{\n"
	structMap += "\t\treturn EncodeStruct" + strings.Title(table) + "(rows)\n"
	structMap += "\t}\n"

	AddToMethodInFile("./genStructs/structMap.go", structMap)
}

//writes structMap.go, which has a function that maps each tableName string to
//its respective function in structInterface.go
func InitStructMap() {
	//declaration, imports
	structMap := "package genStructs\n"
	structMap += "import \"github.com/jmoiron/sqlx\"\n"
	structMap += "func MapTableToJson(tableName string, rows *sqlx.Rows) interface{}{\n"

	//each table has a case mapping name with structInterface function
	tableList := sqlParser.GetTableNames()
	tableList = append(tableList, sqlParser.GetViewNames()...)
	for _, table := range tableList {
		structMap += "\tif tableName == \"" + table + "\"{\n"
		structMap += "\t\treturn EncodeStruct" + strings.Title(table) + "(rows)\n"
		structMap += "\t}\n"
	}

	structMap += "\treturn \"\"\n"
	//if invalid table, returns error
	structMap += "}\n"

	//writes in relation to home directory
	WriteFile(structMap, "./genStructs/structMap.go")
}

/**************************************************************************
 *STRUCT BUILDER HELPER METHODS
 *************************************************************************/

//returns a struct tableName with fields as methods
func MakeStructStr(tableName string, fields []string, fieldTypes []string) string {
	structStr := "type " + strings.Title(tableName) + " struct {\n"

	for idx, field := range fields {
		//lowercase json output
		jsonDec := " `json:\"" + field + "\"`\n"

		//all before (
		re := regexp.MustCompile("^[^(]+")

		//field mapping
		goFieldType := sqlParser.MapColType(re.FindString(fieldTypes[idx]))

		//get string
		structStr += "\t" + strings.Title(field) + "\t\t" + goFieldType + jsonDec
	}

	structStr += "}\n"
	return structStr
}

func AddToFile(fileName string, addString string) {
	cur := ReadFile(fileName)
	cur += addString
	WriteFile(cur, fileName)
}

func AddToMethodInFile(fileName string, addString string) {
	cur := ReadFile(fileName)
	cur = cur[:len(cur)-2]
	cur += addString
	cur += "}\n"
	WriteFile(cur, fileName)
}

func ReadFile(fileName string) string {
	file, err := ioutil.ReadFile(fileName)
	check(err)

	return string(file)
}

//writes string str to fileName
func WriteFile(str string, fileName string) {
	strByte := []byte(str)
	err := ioutil.WriteFile(fileName, strByte, 0644)
	check(err)
}

//removes file fileName
func RemoveFile(fileName string) {
	err := os.Remove(fileName)
	check(err)
}
