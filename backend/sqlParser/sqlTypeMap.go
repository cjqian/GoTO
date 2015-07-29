package sqlParser

import (
	"errors"
	"strconv"
)

func MapColType(SQLType string) string {
	//all unregistered types (datetime for now, etc) are type string
	if GoType, ok := SQLTypeMap[SQLType]; ok {
		return GoType
	} else {
		return "string"
	}
}

var SQLTypeMap = map[string]string{
	//numerical types
	"bigint":  "int64",
	"int":     "int32",
	"integer": "int32",
	"tinyint": "uint8",
	"double":  "float64",
	//string types
	"varchar": "string",
}

func TypeToString(data interface{}) string {
	if bigint, ok := data.(int64); ok {
		return strconv.Itoa(int(bigint))
	} else if intv, ok := data.(int32); ok {
		return strconv.Itoa(int(intv))
	} else if tinyint, ok := data.(uint8); ok {
		return strconv.Itoa(int(tinyint))
	} else if double, ok := data.(float64); ok {
		return strconv.FormatFloat(double, 'f', 2, 32)
	} else if str, ok := data.(string); ok {
		return str
	} else {
		err := errors.New("SQLPARSER: Whoa, what is this type?")
		panic(err)
	}

	return ""
}
