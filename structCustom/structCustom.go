package structCustom

import (
	"./../structDirectory"
	"strconv"
)

func ValidCustomStruct(id string) bool {
	max := structDirectory.GetMaxId()
	cur, _ := strconv.Atoi(id)

	return cur <= max
}
