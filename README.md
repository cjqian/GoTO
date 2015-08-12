#GoTO (Golang Traffic Ops)
##A web API for SQL databases

GoTO is a server/some other stuff written in Go that allows for RESTful interaction with SQL databases through an Angular web API.

This is written for the Comcast [Traffic Ops](http://traffic-control-cdn.net/docs/latest/development/traffic_ops.html) database, but I'm pretty sure it should probably work for all databases.

## Install

1. First, fork a copy of this sick repo. "GoTO" a directory of your choice and type in

```
git clone https://github.com/cjqian/goto.git

```
2. Then, make a `.dbInfo` file that follows this syntax, 
  replacing the content in brackets with your own data:
  ```
  USERNAME="[databaseUsername]"
  PASSWORD="[databasePassword]"
  DATABASE="[databaseName]"
  ```
  For example, if you want to work with the `foo` database with username `johndoe` and password `password`, 
  your `.dbInfo file should look like this:
  ```
  USERNAME="johndoe"
  PASSWORD="password"
  DATABASE="foo"
  ```
  3. In the head of `index.html`, uncomment all of the script/link js/css files and comment out my local ones, minus `main.js.`

  4. Now, you can run the server by typing this into your terminal:
  ```
  ./run
  ```

  Then, start up the Angular front-end by running
  ```
  python -m SimpleHTTPServer
  ```

  Should be up and running on :8000! Make sure ./run is still going concurrently.

## Debugging
  If you're getting errors in the Install process or you happen to be Mark, make sure you can answer "yes" to
  the following questions. If you're still having issues, that really sucks.
  * Do you have the most recent version of Go [installed](https://golang.org/doc/install)? Try uninstalling/reinstalling.
  * Did you make a `.dbInfo` file? (See step two of the [Install](http://github.com/cjqian/GoTO#install) notes.)
  * Are you running `./run` from your `GoTO/` folder and not a subfolder?

  See `./run` for execution examples. Also, are your database credentials correct?
  * Is your `mysql` up and running? Type `mysql` into your terminal to verify.
  * Do you have the latest version of this code? Run `git pull` to get an update. 
  * Also, make sure you've checked out `master` branch and not a development branch.

##Packages
###Local
  * sqlParser processes all interactions with the database. It contains `sqlParser.go`, which contains most of the CRUD methods, and `sqlTypeMap`, which has functions mapping values of type interface{} to string and vice-versa.
  * urlParser parses the url into a Request.
  * outputFormatter wraps the query into an encodable struct.

  There are more details in the comments of each of these packages.
###Other
  * I'm also using AngularJS, jQuery, Bootstrap.
  * `jmoiron/sqlx` has been super useful. Thanks!
  * `ng-react-grid`	and Facebook's `react` was necessary to speed up JSON to table rendering on the API side.
