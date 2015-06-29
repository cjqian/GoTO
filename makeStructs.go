package makeStructs

import (
	"io/ioutil"
	"os"
)

//error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	testString := []byte("hello, world!")
	err := ioutil.WriteFile("structstmp.go", testString, 0644)
	check(err)
}
