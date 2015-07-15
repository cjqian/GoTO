package structCustom

import (
	"./../structDirectory"
	"strconv"
)

func ValidCustomStruct(id string) bool {

	//not custom tables
	if id == "" {
		return true
	}
	max := structDirectory.GetMaxId()
	cur, _ := strconv.Atoi(id)

	return cur <= max
}
