package outputFormatter

import (
	"./../sqlParser"
	//"fmt"
)

//given a []map[string]interface{}, makes the bytes into their according type.
func CleanRowArray(rows []map[string]interface{}) {
	//traverses each row
	for idx, row := range rows {
		//note that all values are []byte
		for k, v := range row {
			keyType := sqlParser.GetColumnType(k)
			newKeyType := sqlParser.MapColumnType(keyType)
			rows[idx][k] = sqlParser.StringToType(v.([]byte), newKeyType)
		}
	}
}
