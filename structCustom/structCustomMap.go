package structCustom
import (
	"github.com/jmoiron/sqlx"
	"net/http"
)
func MapCustomTableToJson(id string, rows *sqlx.Rows, w http.ResponseWriter) {
	if id == "1"{
		EncodeStructCustom1(rows, w)
	}
	if id == "3"{
		EncodeStructCustom3(rows, w)
	}
}
