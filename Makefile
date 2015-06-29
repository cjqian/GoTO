all:
	go build ./sqlToJson/sqlToJson.go
	go build ./structs/structs.go
	go build main.go 
	./main to_user twelve to_development
