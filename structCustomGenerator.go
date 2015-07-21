package main

import (
	"./queryBuilder"
	"./sqlParser"
	"./structBuilder"
	"./structDirectory"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func WriteQuery(q queryBuilder.Query, id string) {
	structBuilder.WriteFile(q.QueryStr, "./structDirectory/queries/queryCustom_"+id)
}
func WriteStructCustom(q queryBuilder.Query, id string) {
	structStr := "//GENERATED PACKAGE\n"
	structStr += "package structCustom\n"

	fieldTypes := GetFieldTypes(q.Fields)
	structStr += structBuilder.MakeStructStr("Custom"+id, q.Fields, fieldTypes)
	structBuilder.WriteFile(structStr, "./structCustom/structCustom_"+id+".go")
}

func WriteJoinStructCustom(q queryBuilder.Query, id string) {
	structStr := "//GENERATED PACKAGE\n"
	structStr += "package structCustom\n"

	fields := make([]string, len(q.Fields))
	fieldEnum := make([]string, len(q.Fields))
	for idx, field := range q.Fields {
		sArray := strings.Split(field, ".")
		fields[idx] = sArray[1]
		fieldEnum[idx] = sArray[0] + "_" + sArray[1]
	}

	fieldTypes := GetFieldTypes(fields)

	structStr += structBuilder.MakeStructStr("Custom"+id, fieldEnum, fieldTypes)
	structBuilder.WriteFile(structStr, "./structCustom/structCustom_"+id+".go")
}

func WriteEncodeStructCustom(id string) {
	encodeBytes, _ := ioutil.ReadFile("./structCustom/structCustomInterface.go")
	encodeStr := string(encodeBytes)

	encodeStr += "func EncodeStructCustom" + id + "(rows *sqlx.Rows, w http.ResponseWriter) {\n"
	encodeStr += "\tsa := make([]Custom" + id + ", 0)\n"
	encodeStr += "\tt := Custom" + id + "{}\n\n"
	encodeStr += "\tfor rows.Next() {\n"
	encodeStr += "\t\trows.StructScan(&t)\n\t\tsa = append(sa, t)\n\t}\n"
	encodeStr += "\tenc := json.NewEncoder(w)\n\tenc.Encode(sa)\n}\n"

	structBuilder.WriteFile(encodeStr, "./structCustom/structCustomInterface.go")
}

func UpdateStructCustomMap(id string) {
	encodeBytes, _ := ioutil.ReadFile("./structCustom/structCustomMap.go")
	encodeStr := string(encodeBytes[:len(encodeBytes)-2])

	encodeStr += "\tif id == \"" + id + "\"{\n"
	encodeStr += "\t\tEncodeStructCustom" + id + "(rows, w)\n\t}\n"

	encodeStr += "}\n"

	structBuilder.WriteFile(encodeStr, "./structCustom/structCustomMap.go")
}

func GetFieldTypes(fields []string) []string {
	ta := make([]string, len(fields))

	for idx, field := range fields {
		ta[idx] = sqlParser.GetColumnType(field)
	}

	return ta
}

func GenerateCustomStruct() {
	id := structDirectory.UpdateCustom()
	q := queryBuilder.MakeQuery()
	WriteQuery(q, id)
	if q.IsJoinType {
		WriteJoinStructCustom(q, id)
	} else {
		WriteStructCustom(q, id)
	}
	WriteEncodeStructCustom(id)
	UpdateStructCustomMap(id)
	fmt.Printf("Custom struct generated.\n Struct ID is: %s\n", id)
}

func main() {
	//connect to database
	sqlParser.ConnectToDatabase(os.Args[1], os.Args[2], os.Args[3])
	GenerateCustomStruct()
}
