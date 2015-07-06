package sqlParser

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
