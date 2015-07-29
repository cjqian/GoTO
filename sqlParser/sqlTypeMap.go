package sqlParser

import (
	"errors"
	//	"fmt"
	"strconv"
	"strings"
)

func MapColumnTypes(SQLTypes []string) []string {
	ta := make([]string, 0)
	for _, t := range SQLTypes {
		ta = append(ta, MapColumnType(t))
	}

	return ta
}
func MapColumnType(SQLType string) string {
	//this is due to the (11), etc type affixed to certain sqlTypes
	SQLType = strings.Split(SQLType, "(")[0]
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
	"int":     "int64",
	"integer": "int64",
	"tinyint": "int64",
	"double":  "float64",
	//string types
	"varchar": "string",
}

//given a []byte s and type t, return s in t form
func StringToType(b []byte, t string) interface{} {
	t = strings.Split(t, "(")[0]
	//all unregistered types (datetime for now, etc) are type string
	s := string(b)
	if t == "bigint" || t == "int" || t == "integer" || t == "tinyint" {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return i
	} else if t == "double" {
		float, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err)
		}
		return float
	} else if t == "varchar" {
		return s
	} else {

		return string(b)
	}

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
