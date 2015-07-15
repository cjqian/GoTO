package structDirectory

import (
	"./../structBuilder"
	//	"fmt"
	"io/ioutil"
	"strconv"
)

func GetMaxId() int {
	maxID, err := ioutil.ReadFile("./structDirectory/curCustomNumber")
	if err != nil {
		panic(err)
	}

	str, err2 := strconv.Atoi(string(maxID))
	if err2 != nil {
		panic(err2)
	}

	return str
}

func UpdateCustom() string {
	customStr, err := ioutil.ReadFile("./structDirectory/curCustomNumber")
	if err != nil {
		panic(err)
	}

	customNum, _ := strconv.Atoi(string(customStr))
	customNum += 1

	newCustomStr := strconv.Itoa(customNum)
	structBuilder.WriteFile(newCustomStr, "./structDirectory/curCustomNumber")
	return newCustomStr
}
