package structFilter

import (
	"./../sqlParser"
	"./../structBuilder"
	"fmt"
	"github.com/jmoiron/sqlx"
	// "log"
	"net/http"
	// "os/exec"
)

func MapCustomTableToJson(tableName string, rows *sqlx.Rows, w http.ResponseWriter, fields []string) {
	//first make a new struct file custom
	MakeStructCustom(tableName, fields)

	// cmd := exec.Command("go", "build", "structFilter/structCustom.go")
	// err := cmd.Run()
	/* if err != nil {
		log.Fatal(err)
	} */
	//encode struct custom
	EncodeStructCustom(rows, w)
}

func MakeStructCustom(tableName string, fields []string) {
	structStr := "package structFilter\n"

	//get fieldTypes
	fieldTypes := GetFieldTypes(fields)
	structStr += structBuilder.MakeStructStr("Custom", fields, fieldTypes)
	structBuilder.WriteFile(structStr, "./structFilter/structCustom.go")
	fmt.Printf("%s\n", "Struct Custom built.")
}

func GetFieldTypes(fields []string) []string {
	ta := make([]string, len(fields))

	for idx, field := range fields {
		ta[idx] = sqlParser.GetColumnType(field)
	}

	return ta
}
