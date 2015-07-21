package queryBuilder

import (
	"./../sqlParser"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Query struct {
	QueryStr string
	Tables   []string
	Fields   []string
	//either custom or auto
	IsJoinType bool
}

/***************************************************************************
 *CUSTOM STRUCT SECTION
 **************************************************************************/
func ParseCustom(qStr string) Query {
	qStr = strings.ToLower(qStr)
	selectLoc := strings.Index(qStr, "select")
	fromLoc := strings.Index(qStr, "from")
	joinLoc := strings.Index(qStr, "join")
	onLoc := strings.Index(qStr, " on ")

	fmt.Printf("Select: %d, From: %d, Join: %d, On:%d\n", selectLoc, fromLoc, joinLoc, onLoc)
	//valid
	if (selectLoc == -1 || fromLoc == -1) ||
		(joinLoc != -1 && onLoc == -1) ||
		(joinLoc == -1 && onLoc != -1) {
		err := errors.New("You obviously don't know how to write a SQL query")
		panic(err)
	}

	isJoinType := false
	if joinLoc != -1 {
		isJoinType = true
	}

	//get fields
	parameterStr := qStr[selectLoc+7 : fromLoc]
	qFields := strings.Split(parameterStr, ",")

	//get qTables
	var qTables []string
	if !isJoinType {
		qTables = append(qTables, strings.Fields(qStr[fromLoc+4:])...)
	} else {
		qTables = append(qTables, strings.Fields(qStr[fromLoc+4:joinLoc])...)
		//fmt.Println("Appended from to join", qTables)
		qTables = append(qTables, strings.Fields(qStr[joinLoc+5:onLoc])...)
		//fmt.Println("Appended join after", qTables)
	}

	//clean up tables and fields
	for idx, table := range qTables {
		qTables[idx] = strings.TrimSpace(table)
	}

	for idx, field := range qFields {
		qFields[idx] = strings.TrimSpace(field)
	}

	//get qFields
	if isJoinType {
		//check asterisks
		if len(qFields) == 1 && qFields[0] == "*" {
			//case "select * from x join y"
			qFields = sqlParser.GetTableColumnNames(qTables[0])
			qFields = append(qFields, sqlParser.GetTableColumnNames(qTables[1])...)
		} else if strings.Count(parameterStr, ".") != len(qFields) {
			//make sure not ambiguous
			err := errors.New("Ambiguous parameter types.")
			panic(err)
		}
		//at this point, all should have table.type structure.
		//we check for asterisks
		for idx, field := range qFields {
			split := strings.Split(field, ".")
			//if asterisk, remove asterisk and add all
			if split[1] == "*" {
				qFieldsEnd := qFields[idx+1:]
				qFields = append(qFields[:idx], sqlParser.GetTableColumnNames(split[0])...)
				qFields = append(qFields, qFieldsEnd...)
			}
		}
	} else {
		//check no headings in non-joins
		for idx, field := range qFields {
			if strings.Contains(field, ".") {
				//make sure heading is accurate
				split := strings.Split(field, ".")
				if split[0] == qTables[0] {
					qFields[idx] = split[1]
				} else {
					err := errors.New("What are you even doing that is not the table -_-")
					panic(err)
				}
			}

			//at this point, should have type structure
			//we check for asterisks
			if field == "*" {
				fmt.Println("asterisk detected")
				qFields = sqlParser.GetColumnNames(qTables[0])
				break
			}
		}
	}

	q := Query{qStr, qTables, qFields, isJoinType}
	return q
}

/***************************************************************************
 *AUTO STRUCT SECTION
 **************************************************************************/
func ParseAuto() Query {
	query := "select "

	table1 := QueryTable()
	cols1 := QueryCols(table1)

	for _, col := range cols1 {
		query += col + ", "
		//query += table1 + "." + col + ", "
	}

	table2 := QueryJoin(table1)
	cols2 := QueryCols(table2)

	for _, col := range cols2 {
		//query += table2 + "." + col + ", "
		query += col + ", "
	}

	//remove comma
	query = query[:len(query)-2] + " "

	//add join
	query += "FROM " + table1 + " JOIN " + table2

	//equiv
	query += " ON " + QueryEquiv(table1, table2) + ";"

	qTables := []string{table1, table2}
	qFields := append(cols1, cols2...)

	q := Query{query, qTables, qFields, true}
	fmt.Println(q)
	return q
}

//returns name of table
func QueryTable() string {
	qTables := sqlParser.GetTableNames()

	//print options
	fmt.Printf("%s\n", "Select a table:")
	for idx, table := range qTables {
		fmt.Printf("[%d] %s\n", idx, table)
	}

	//get option
	var optionNum int
	_, err := fmt.Scanln(&optionNum)
	if err != nil {
		log.Fatal(err)
	}

	//debug
	fmt.Println(qTables[optionNum])

	return qTables[optionNum]
}

//returns tableName.columnName
//note that this only works for joins (which works for now)
func QueryCols(tableName string) []string {
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
		colArray[idx] = tableName + "." + cols[optionNum]
	}

	return colArray
}

//can only join with one
func QueryJoin(tableName string) string {
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

//given the name of two tables, returns a "table1.column1 = table2.column2" statement
func QueryEquiv(table1 string, table2 string) string {
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
	return equivStmt

}

/***************************************************************************
 *END AUTO STRUCT SECTION
 **************************************************************************/

//writes to tmpQuery
func MakeQuery() Query {
	fmt.Println("Custom [0] or AutoQuery [1]?")
	var response int
	_, err := fmt.Scanln(&response)
	if err != nil || response > 1 {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)

	var q Query
	if response == 0 {
		fmt.Println("Enter custom here: (no need for end semi-colon)")
		qStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		q = ParseCustom(qStr)
	} else {
		q = ParseAuto()
	}

	fmt.Println("Query made!", q)
	return q
}

func main() {
	sqlParser.ConnectToDatabase("to_user", "twelve", "to_development")
	MakeQuery()
}
