package main

import (
	"io/ioutil"
)

//error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	testString := []byte("hello, world!")
	err := ioutil.WriteFile("structs/structstmp.go", testString, 0644)
	check(err)
}
