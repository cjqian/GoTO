#GoTO (Golang Traffic Ops)

GoTO is a web API that returns JSON formatting for SQL database tables (specifically, the Comcast [Traffic Ops](http://traffic-control-cdn.net/docs/latest/development/traffic_ops.html) database). 

## Releases
* 0.2.0 (Upcoming)
	* Clean up output JSON formatting and object naming
	* Generate struct handler with OK/404 for URLs (should solve .favicon issue)
	* Return JSON format via curl, displayed as .json output in browser
* [0.1.2] (7/1/2015) Added variable struct generation via bash script. 
* [0.1.1](https://github.com/cjqian/GoTO/commit/11914007c8ccd3d1d0eb039cc25abc1a8decfc34) (7/1/2015)
	Documentation is updated and code is cleaned up. 
* [0.1.0](https://github.com/cjqian/jsonserver/commit/be727ea8bb4597126c3171d9f809a0437833b9a5) (6/30/2015)
	Basic packages are sketched out and incorporated with the main server; everything
	works, kind of. See the demo [here.](https://www.dropbox.com/s/7u48ihlxkuytmxn/demo_presentation.pdf?dl=0)

## Install/Usage
1. First, fork a copy of this sick repo. "GoTO" a directory of your choice and type in:
```
git clone https://github.com/cjqian/GoTO.git
```
2. Then, make a `dbInfo` file that follows this snytax, replacing the content in 
brackets with your own data:
```
USERNAME="[databaseUsername]"
PASSWORD="[databasePassword]"
ENVIRONMENT="[databaseName]"
```

3. Now, you can run the server by typing this into your terminal:
```
./runGoto
```
Your system should print out `Structs Generated.` during your first run, because you shouldn't have a `structs/` folder yet. 

In the future, if you wish to regenerate the structs package, run `./runGoto gs`. The `gs` argument will "generate structs."

## Known Issues
* Favicon.ico responses are breaking the program.
* Inaccurate JSON formatting.

## Packages
### Main.go

This is the main Go program that starts the web service and listens for requests. 

Requests are in the form:
```go
url/[table_in_database]
//for example,
http://localhost:8000/deliveryservice
```

Which will return the JSON for the "deliveryservice" table in the database.

The program takes in three parameters: the username, password and database. 

### SQL Parser

This package (sqlParser) contains the following public methods for i
interacting with the database. Also, I'm using the [SQLX library](http://jmoiron.github.io/sqlx/).

```go
// connects to and returns a pointer to the database
func ConnectToDatabase(username string, password string, environment string) sqlx.DB;

//returns an array of table name strings from queried database
func GetTableNames(db sqlx.DB) []string;

//returns *Rows from given table (name) from queried database
func GetRows(db sqlx.DB, tableName string) *sqlx.Rows;

//returns array of column names from table (name) in database
func GetColumnNames(db sqlx.DB, tableName string) []string; 
```

### Struct Constructor

This package (structConstructor) contructs the following package (see Structs) 
by generating three .go files. This is run by adding the `gs` argument (`./runGoto gs`) or when the `structs` package is not found.

```go
//writes struct, interface, and map files to structs package
func MakeStructFiles(db sqlx.DB);

//writes the struct file, which has an object for each database table, 
//with each table field as a member variable
func MakeStructs(db sqlx.DB);

//writes structInterface.go, which has functions that take in *Rows and
//return the byte array JSON format for each table in the database
func MakeStructInterface(db sqlx.DB);

//writes structMap.go, which has one function that maps each tableName string
//to its respective function in structInterface.go`
func MakeStructMap(db sqlx.DB);

//writes string str to fileName, helper function for the above three
func WriteFile(str string, fileName string);
```

### Structs

This package (structs) is dynamically generated on server start from Struct Constructor (see above). 
There are three files:
* structs.go
```go
type [TableName] struct{
	[Table Field]	[Field Type]
	...
}
```

* structInterface.go
```go
func ByteArrayFrom[Table Name](rows *sqlx.Rows) []byte{
	var tStr []byte

	//creates a new [Table Name] object (defined in Structs) and scans
	//contents of given rows into its fields. appends the JSON 
	//representation to a tStr byte array representing the entire table
	t := [Table Name]{}
	for rows.Next() {
		rows.StructScan(&t)
		tmpStr, _ := json.MarshalIndent(t, "", "  ")
		tStr = append(tStr[:], tmpStr[:]...)
	}

	return tStr
}
```

* structMap.go
```go
func MapTableToJson(tableName string, rows *sqlx.Rows) []byte{
	if tableName == [Table Name]{
		tStr := ByteArrayFromAsn(row)
		return tStr
	}
	...
}
```
=======

## Known Issues
* Favicon.ico responses are breaking the program.
* Inaccurate JSON formatting.

## Packages
### Main.go

This is the main Go program that starts the web service and listens for requests. 

Requests are in the form:
```go
url/[table_in_database]
//for example,
http://localhost:8000/deliveryservice
```

Which will return the JSON for the "deliveryservice" table in the database.

The program takes in three parameters: the username, password and database. 

### SQL Parser

This package (sqlParser) contains the following public methods for i
interacting with the database. Also, I'm using the [SQLX library](http://jmoiron.github.io/sqlx/).

```go
// connects to and returns a pointer to the database
func ConnectToDatabase(username string, password string, environment string) sqlx.DB;

//returns an array of table name strings from queried database
func GetTableNames(db sqlx.DB) []string;

//returns *Rows from given table (name) from queried database
func GetRows(db sqlx.DB, tableName string) *sqlx.Rows;

//returns array of column names from table (name) in database
func GetColumnNames(db sqlx.DB, tableName string) []string; 
```

### Struct Constructor

This package (structConstructor) contructs the following package (see Structs) 
by generating three .go files. Main.go runs the construction whenever the server is started.

```go
//writes struct, interface, and map files to structs package
func MakeStructFiles(db sqlx.DB);

//writes the struct file, which has an object for each database table, 
//with each table field as a member variable
func MakeStructs(db sqlx.DB);

//writes structInterface.go, which has functions that take in *Rows and
//return the byte array JSON format for each table in the database
func MakeStructInterface(db sqlx.DB);

//writes structMap.go, which has one function that maps each tableName string
//to its respective function in structInterface.go`
func MakeStructMap(db sqlx.DB);

//writes string str to fileName, helper function for the above three
func WriteFile(str string, fileName string);
```

### Structs

This package (structs) is dynamically generated on server start from Struct Constructor (see above). 
There are three files:
* structs.go
```go
type [TableName] struct{
	[Table Field]	[Field Type]
	...
}
```

* structInterface.go
```go
func ByteArrayFrom[Table Name](rows *sqlx.Rows) []byte{
	var tStr []byte

	//creates a new [Table Name] object (defined in Structs) and scans
	//contents of given rows into its fields. appends the JSON 
	//representation to a tStr byte array representing the entire table
	t := [Table Name]{}
	for rows.Next() {
		rows.StructScan(&t)
		tmpStr, _ := json.MarshalIndent(t, "", "  ")
		tStr = append(tStr[:], tmpStr[:]...)
	}

	return tStr
}
```

* structMap.go
```go
func MapTableToJson(tableName string, rows *sqlx.Rows) []byte{
	if tableName == [Table Name]{
		tStr := ByteArrayFromAsn(row)
		return tStr
	}
	...
}
```
