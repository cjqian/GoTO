#GoTO (Golang Traffic Ops)

GoTO is a web API that returns JSON formatting for SQL database tables (specifically, the Comcast [Traffic Ops](http://traffic-control-cdn.net/docs/latest/development/traffic_ops.html) database). 

## Releases
* 0.2.1 (Upcoming)
	* Dynamic queries in URL
* [0.2.0](https://github.com/cjqian/GoTO/commit/7b0ed143dc8287a6405947eb80f8c1338d54f3de) (7/6/2015)
	* Clean up output JSON formatting and object naming
	* Generate struct handler with OK/404 for URLs (should solve .favicon issue) [Done!]
	* Return JSON format via curl, displayed as .json output in browser [Done!]
* [0.1.2](https://github.com/cjqian/GoTO/commit/e6a30c4010ebe15cf0e5adf01691a6f434484731) (7/1/2015) Added variable struct generation via bash script. 
* [0.1.1](https://github.com/cjqian/GoTO/commit/11914007c8ccd3d1d0eb039cc25abc1a8decfc34) (7/1/2015)
	Documentation is updated and code is cleaned up. 
* [0.1.0](https://github.com/cjqian/jsonserver/commit/be727ea8bb4597126c3171d9f809a0437833b9a5) (6/30/2015)
	Basic packages are sketched out and incorporated with the main server; everything
	works, kind of. See the demo [here.](https://www.dropbox.com/s/7u48ihlxkuytmxn/demo_presentation.pdf?dl=0)

## Install/Usage

1. First, fork a copy of this sick repo. "GoTO" a directory of your choice and type in

  ```
  git clone https://github.com/cjqian/GoTO.git
  ```
2. Then, make a `dbInfo` file that follows this syntax, 
	replacing the content in brackets with your own data:
  ```
  USERNAME="[databaseUsername]"
  PASSWORD="[databasePassword]"
  ENVIRONMENT="[databaseName]"
  ```

3. Now, you can run the server by typing this into your terminal:
  ```
  ./runGoto
  ```
  Your system should print out `Structs Generated.` during your first run, 
  because you shouldn't have a `structs/` folder yet. 

  In the future, if you wish to regenerate the structs package, run `./runGoto gs`. 
  
  The `gs` argument will "generate structs."


## Debugging
If you're getting errors in the Install process or you happen to be Mark, make sure you can answer "yes" to
the following questions. If you're still having issues, that really sucks.
* Do you have the most recent version of Go [installed](https://golang.org/doc/install)? Try uninstalling/reinstalling.
* Did you make a `dbInfo` file? (See step two of the [Install](http://github.com/cjqian/GoTO#installusage) notes.)
* Are you running `./runGoto` from your `GoTO/` folder and not a subfolder?
* Alternatively, if you're not using the `./runGoto` command, did you make sure to add arguments when running the build? 
  ```
  ./[program] [username] [password] [environment]
  ``` 
  See `./runGoto` for execution examples. Also, are your database credentials correct?
* Is your `mysql` up and running? Type `mysql` into your terminal to verify.
* Do you have the latest version of this code? Run `git pull` to get an update. 
* Also, make sure you've checked out `master` branch and not a development branch.

## Syntax 
Note: not yet implemented!
* `read/` should display a list of all tables in the database.
* `read/[tableName]` should display table fields and corresponding type in JSON.	
* `encode/[tableName]` should output JSON of all rows in table.
* CRUD functionality to come (ex. `create/`, `delete/`, potentially?)
 
## Packages
### Main.go

This is the main Go program that starts the web service and listens for requests. 

Requests are currently in the form:
```
url/[table_in_database]	//syntax
localhost:8000/deliveryservice //example
```

Which will return the JSON for the "deliveryservice" table in the database.

If the table queried exists in the database (checked against structValidMap), the program will print "[table name] found."
Else, the program will print "[table name] is not found."

This will be changed after (Syntax)[http://www.github.com/cjqian/GoTO/#Syntax] is implemented.

The program takes in three parameters: the username, password and database. 

### SQL Parser

This package (sqlParser) contains the following public methods for interacting with the database. 
Also, I'm using the [SQLX library](http://jmoiron.github.io/sqlx/).

There are two files:
* `sqlParser.go` has the following SQL database API.
```go
// connects to and returns a pointer to the database
func ConnectToDatabase(username string, password string, environment string) sqlx.DB {
	...
}

//returns an array of table name strings from queried database
func GetTableNames(db sqlx.DB) []string {
	...
}

//returns *Rows from given table (name) from queried database
func GetRows(db sqlx.DB, tableName string) *sqlx.Rows {
	...
}

//returns array of column names from table (name) in database
func GetColumnNames(db sqlx.DB, tableName string) []string {
	...
} 

//returns array of column types from table (name) in database
func GetColumnTypes(db sqlx.DB, tableName string) []string {
	...
}
```

* `SQLTypeMap.go` contains mappings from SQL data types to Golang datatypes
```go
//given a SQL data type, returns the name of the equivalent Golang data type
//unless it's not in the SQLTypeMap, in which case it returns string
//note: timestamp is intentionally mapped to a string
func MapColType(SQLType string) string {
	...
}

//map of SQL types to Golang types
var SQLTypeMap = map[string]string{
	//format
	[SQL Type]	: 	[Golang Type]
	//example
	"bigint"	:	"int64"
}
```

### Struct Generator

This package (structGenerator) contructs the following package (see [Structs](https://github.com/cjqian/GoTO/#structs)) 
by generating .go files. This is run by adding the `gs` argument (`./runGoto gs`) or when the `structs` package is not found.

Alternatively, this could be built on its own by running (from the home directory)
```
go build structGenerator/structGenerator.go 
./structGenerator/structGenerator [databaseUsername] [databasePassword] [databaseName]
```

```go
//writes struct, interface, valid map and map files to structs package
func MakeStructFiles() {
	...
}

//writes the struct file, which has an object for each database table, 
//with each table field as a member variable
func MakeStructs() {
	...
}

//writes structInterface.go, which has functions that take in *Rows and
//parses them into an array of structs, writing the resulting JSON 
//encoding to the writer argument. has one function for each table.
func MakeStructInterface() {
	...
}

//writes structValidMap.go, which maps each table in the database to the boolean "true,"
//used to confirm validity of URL
func MakeStructValidMap() {
	...
}

//writes structMap.go, which has one function that maps each tableName string
//to its respective function in structInterface.go`
func MakeStructMap() {
	...
}

//writes string str to fileName, helper function for the above three
func WriteFile(str string, fileName string) {
	...
}
```
### Structs

This package (structs) is dynamically generated on server start from [Struct Generator](https://github.com/cjqian/GoTO/#struct-generator). 
These files are made:

* `structs.go`
```go
type [TableName] struct{
	[Table Field]	[Field Type]
}
	...
```

* `structInterface.go`
```go
func EncodeStruct[Table Name](rows *sqlx.Rows, w http.ResponseWriter) {
	//Makes an array of structs of type [Table Name]
	sa := make([][Table Name], 0)

	//creates a new [Table Name] object (defined in Structs) and scans
	//contents of given rows into its fields, appending the object
	//to the object array
	t := [Table Name]{}
	for rows.Next() {
		rows.StructScan(&t)
		sa = append(sa, t)
	}

	//encodes the resultant array JSON representation to the writer
	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
```

* `structMap.go`
```go
func MapTableToJson(tableName string, rows *sqlx.Rows, w http.ResponseWriter) []byte{
	if tableName == [Table Name]{
		EncodeStruct[Table Name](rows, w)
	}
	...
}
```

* `structValidMap.go`
```go
var ValidStruct = map[string]bool {
	//format
	[Table Name] : true,
}
```
