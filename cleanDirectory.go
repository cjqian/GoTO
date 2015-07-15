package main

import (
	"./structBuilder"
	"fmt"
)

func cleanInterface() {
	interfaceStr := "package structCustom\n"
	interfaceStr += "import (\n"
	interfaceStr += "\t\"encoding/json\"\n"
	interfaceStr += "\t\"github.com/jmoiron/sqlx\"\n"
	interfaceStr += "\t\"net/http\"\n"
	interfaceStr += ")\n"

	structBuilder.WriteFile(interfaceStr, "./structCustom/structCustomInterface.go")
}

func cleanMap() {
	mapStr := "package structCustom\n"
	mapStr += "import (\n"
	mapStr += "\t\"github.com/jmoiron/sqlx\"\n"
	mapStr += "\t\"net/http\"\n"
	mapStr += ")\n"

	mapStr += "func MapCustomTableToJson(id string, rows *sqlx.Rows, w http.ResponseWriter) {\n"
	mapStr += "}\n"

	structBuilder.WriteFile(mapStr, "./structCustom/structCustomMap.go")
}

func cleanCounter() {
	counter := "0"
	structBuilder.WriteFile(counter, "./structDirectory/curCustomNumber")
}

func main() {
	cleanInterface()
	cleanMap()
	cleanCounter()

	fmt.Println("Generated files refreshed.")
}
