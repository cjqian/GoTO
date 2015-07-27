package structBuilder

import (
	"./../sqlParser"
	//"fmt"
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
