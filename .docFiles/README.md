#GoTO (Golang Traffic Ops)

GoTO is a web API that returns JSON formatting for SQL database tables (specifically, the Comcast [Traffic Ops](http://traffic-control-cdn.net/docs/latest/development/traffic_ops.html) database). 

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
  ./run
  ```
  Your system should print out `Structs Generated.` during your first run, 
  because you shouldn't have a `structs/` folder yet. 

  In the future, if you wish to regenerate the structs package, run `./run gs`. 
  
  The `gs` argument will "generate structs."

  If you wish to initiate the custom struct generator, run `./run cs`. 


## Debugging
If you're getting errors in the Install process or you happen to be Mark, make sure you can answer "yes" to
the following questions. If you're still having issues, that really sucks.
* Do you have the most recent version of Go [installed](https://golang.org/doc/install)? Try uninstalling/reinstalling.
* Did you make a `dbInfo` file? (See step two of the [Install](http://github.com/cjqian/GoTO#installusage) notes.)
* Are you running `./run` from your `GoTO/` folder and not a subfolder?
* Alternatively, if you're not using the `./run` command, did you make sure to add arguments when running the build? 
  
  ```
  ./[program] [username] [password] [environment]
  ``` 
  
  See `./run` for execution examples. Also, are your database credentials correct?
* Is your `mysql` up and running? Type `mysql` into your terminal to verify.
* Do you have the latest version of this code? Run `git pull` to get an update. 
* Also, make sure you've checked out `master` branch and not a development branch.

## Syntax 
Note: not yet implemented!
* `read/` should display a list of all tables in the database.
* `read/[tableName]` should display table fields and corresponding type in JSON.	
* `encode/[tableName]` should output JSON of all rows in table.
* CRUD functionality to come (ex. `create/`, `delete/`, potentially?)

##Scripts
###clean
Running `./clean` runs the `cleanDirectory` program, removes the list of queries from `structDirectory/queries`, and resets the initial query/ID count.
###run
This starts up the server on default port :8080.
    * `./run gs` will regenerate default structs.
    * `./run cs` will run the QueryBuilder script beforehand, which will create a new query and custom struct.
###show
Running `./show` displays each query filepath along with its specific query.

##Packages
###sqlParser
This package contains all methods that interact with the database.

###structBuilder
This package contains basic struct building methods used to generate both default and custom structs.

###structCustom
This package is mostly generated by structCustomGenerator.go, and will contain all the structs made by user queries.

###structDirectory
This package contains a directory of all queries and functions to help update the current max ID.###structs
This package is generated by structGenerator.go, and contains structs for each database table as wel as a map and an encoder of each respective function.
###urlParser
This package contains functions for interacting with the database.

##Programs
###cleanDirectory.go
This rewrites the structDirector/structCustom directories to default and resets the IDs.

###main.go
This runs the web server.

###structCustomGenerator.go
This runs a command-line query builder and builds a custom struct.

###structGenerator.go
This generates all default structs from the database.