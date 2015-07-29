package genStructs
import (
	"github.com/jmoiron/sqlx"
	"encoding/json"
	"net/http"
)
func EncodeStructAsn(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Asn, 0)
	t := Asn{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructCachegroup(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Cachegroup, 0)
	t := Cachegroup{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructCachegroup_parameter(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Cachegroup_parameter, 0)
	t := Cachegroup_parameter{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructDeliveryservice(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Deliveryservice, 0)
	t := Deliveryservice{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructDeliveryservice_regex(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Deliveryservice_regex, 0)
	t := Deliveryservice_regex{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructDeliveryservice_server(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Deliveryservice_server, 0)
	t := Deliveryservice_server{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructDeliveryservice_tmuser(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Deliveryservice_tmuser, 0)
	t := Deliveryservice_tmuser{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructDivision(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Division, 0)
	t := Division{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructGoose_db_version(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Goose_db_version, 0)
	t := Goose_db_version{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructHwinfo(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Hwinfo, 0)
	t := Hwinfo{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructJob(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Job, 0)
	t := Job{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructJob_agent(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Job_agent, 0)
	t := Job_agent{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructJob_result(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Job_result, 0)
	t := Job_result{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructJob_status(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Job_status, 0)
	t := Job_status{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructLog(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Log, 0)
	t := Log{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructParameter(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Parameter, 0)
	t := Parameter{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructPhys_location(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Phys_location, 0)
	t := Phys_location{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructProfile(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Profile, 0)
	t := Profile{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructProfile_parameter(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Profile_parameter, 0)
	t := Profile_parameter{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructRegex(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Regex, 0)
	t := Regex{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructRegion(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Region, 0)
	t := Region{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructRole(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Role, 0)
	t := Role{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructServer(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Server, 0)
	t := Server{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructServercheck(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Servercheck, 0)
	t := Servercheck{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructStaticdnsentry(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Staticdnsentry, 0)
	t := Staticdnsentry{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructStatus(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Status, 0)
	t := Status{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructTm_user(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Tm_user, 0)
	t := Tm_user{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructTo_extension(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]To_extension, 0)
	t := To_extension{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructType(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Type, 0)
	t := Type{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructCondensedasn(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Condensedasn, 0)
	t := Condensedasn{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructCrystal(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Crystal, 0)
	t := Crystal{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructDeliveryservicesdenver(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Deliveryservicesdenver, 0)
	t := Deliveryservicesdenver{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
func EncodeStructHimark(rows *sqlx.Rows, w http.ResponseWriter) {
	sa := make([]Himark, 0)
	t := Himark{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	enc := json.NewEncoder(w)
	enc.Encode(sa)
}
