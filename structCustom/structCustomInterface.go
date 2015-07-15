package structCustom
import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"net/http"
)
func EncodeStructCustom1(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Custom1, 0)
	t := Custom1{}

	for rows.Next() {
		rows.StructScan(&t)
		sa = append(sa, t)
	}
	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructCustom3(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Custom3, 0)
	t := Custom3{}

	for rows.Next() {
		rows.StructScan(&t)
		sa = append(sa, t)
	}
	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
