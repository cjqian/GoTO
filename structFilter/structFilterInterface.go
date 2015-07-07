package structFilter

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func EncodeStructCustom(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Custom, 0)
	t := Custom{}

	for rows.Next() {
		rows.StructScan(&t)
		sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
