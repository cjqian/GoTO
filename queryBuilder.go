package main

import (
	"./sqlParser"
	"./structBuilder"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func queryTable() string {
	tables := sqlParser.GetTableNames()

	//print options
	fmt.Printf("%s\n", "Select a table:")
	for idx, table := range tables {
		fmt.Printf("[%d] %s\n", idx, table)
	}

	//get option
	var optionNum int
	_, err := fmt.Scanln(&optionNum)
	if err != nil {
		log.Fatal(err)
	}

	//debug
	fmt.Println(tables[optionNum])

	return tables[optionNum]
}

func queryCols(tableName string) []string {
	cols := sqlParser.GetColumnNames(tableName)

	//print options
	fmt.Printf("Select columns from %s (separated by comma)\n", tableName)
	for idx, col := range cols {
		fmt.Printf("[%d] %s\n", idx, col)
	}

	//get option
	var optionStr string
	_, err := fmt.Scanln(&optionStr)
	if err != nil {
		log.Fatal(err)
	}

	optionsStr := strings.Split(optionStr, ",")

	var colArray = make([]string, len(optionsStr))
	for idx, opt := range optionsStr {
		optionNum, _ := strconv.Atoi(opt)
		colArray[idx] = cols[optionNum]
	}

	//debug
	fmt.Println(colArray)

	return colArray
}

//can only join with one
func queryJoin(tableName string) string {
	joinQuery := "select referenced_table_name from information_schema.referential_constraints where table_name='" + tableName + "';"
	neighborTables := sqlParser.GetCustomColumnNames(joinQuery)

	fmt.Println("Select table to join:")
	for idx, table := range neighborTables {
		fmt.Printf("[%d] %s\n", idx, table)
	}

	var joinNum int
	_, err := fmt.Scanln(&joinNum)
	if err != nil {
		log.Fatal(err)
	}

	//debug
	fmt.Println(neighborTables[joinNum])

	return neighborTables[joinNum]
}

func queryEquiv(table1 string, table2 string) string {
	table1cols := sqlParser.GetColumnNames(table1)
	table2cols := sqlParser.GetColumnNames(table2)

	//select 1
	fmt.Printf("%s\n", "Select equivalent statement part 1:")
	for idx, col := range table1cols {
		fmt.Printf("[%d] %s.%s\n", idx, table1, col)
	}

	var option1 int
	_, err := fmt.Scanln(&option1)
	if err != nil {
		log.Fatal(err)
	}

	//select 2
	fmt.Printf("%s\n", "Select equivalent statement part 2:")
	for idx, col := range table2cols {
		fmt.Printf("[%d] %s.%s\n", idx, table2, col)
	}

	var option2 int
	_, err2 := fmt.Scanln(&option2)
	if err2 != nil {
		log.Fatal(err)
	}

	equivStmt := table1 + "." + table1cols[option1] + "=" + table2 + "." + table2cols[option2]

	//debug
	fmt.Println(equivStmt)

	return equivStmt

}

//writes to tmpQuery
func writeQuery() string {
	query := "select "

	table1 := queryTable()
	cols1 := queryCols(table1)

	for _, col := range cols1 {
		query += table1 + "." + col + ", "
	}

	table2 := queryJoin(table1)
	cols2 := queryCols(table2)

	for _, col := range cols2 {
		query += table2 + "." + col + ", "
	}

	//remove comma
	query = query[:len(query)-2] + " "

	//add join
	query += "FROM " + table1 + " JOIN " + table2

	//equiv
	query += " ON " + queryEquiv(table1, table2) + ";"

	fmt.Println(query)

	structBuilder.WriteFile(query, "./tmpQuery")
	return query

}

func main() {
	//connect to database
	sqlParser.ConnectToDatabase(os.Args[1], os.Args[2], os.Args[3])
	writeQuery()
}
