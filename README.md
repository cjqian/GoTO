A web server written in Go that converts SQL data into a JSON file.

Progress:

1. Followed golang tutorial to make a wikipage edit/view server. (6/17)
2. Removed parts that won't be used. Now, when a button is pushed, output is generated and displayed. (6/19)
3. Given a database table name, returns a pseudo-json structure. (6/29)

Modules:

1. structs: currently hardcoded details on each table in the to_development database. Need to write a script to dynamically generate this file.
2. sqlToJson: two methods in this, one to return rows from a sql database, another to make a json byte array from a dict of rows. Not sure if I should just combine into one method.
