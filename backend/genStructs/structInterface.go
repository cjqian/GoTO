package genStructs
import (
	"github.com/jmoiron/sqlx"
)
func EncodeStructAsn(rows *sqlx.Rows) interface{} {
	sa := make([]Asn, 0)
	t := Asn{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructCachegroup(rows *sqlx.Rows) interface{} {
	sa := make([]Cachegroup, 0)
	t := Cachegroup{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructCachegroup_parameter(rows *sqlx.Rows) interface{} {
	sa := make([]Cachegroup_parameter, 0)
	t := Cachegroup_parameter{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructDeliveryservice(rows *sqlx.Rows) interface{} {
	sa := make([]Deliveryservice, 0)
	t := Deliveryservice{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructDeliveryservice_regex(rows *sqlx.Rows) interface{} {
	sa := make([]Deliveryservice_regex, 0)
	t := Deliveryservice_regex{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructDeliveryservice_server(rows *sqlx.Rows) interface{} {
	sa := make([]Deliveryservice_server, 0)
	t := Deliveryservice_server{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructDeliveryservice_tmuser(rows *sqlx.Rows) interface{} {
	sa := make([]Deliveryservice_tmuser, 0)
	t := Deliveryservice_tmuser{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructDivision(rows *sqlx.Rows) interface{} {
	sa := make([]Division, 0)
	t := Division{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructGoose_db_version(rows *sqlx.Rows) interface{} {
	sa := make([]Goose_db_version, 0)
	t := Goose_db_version{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructHwinfo(rows *sqlx.Rows) interface{} {
	sa := make([]Hwinfo, 0)
	t := Hwinfo{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructJob(rows *sqlx.Rows) interface{} {
	sa := make([]Job, 0)
	t := Job{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructJob_agent(rows *sqlx.Rows) interface{} {
	sa := make([]Job_agent, 0)
	t := Job_agent{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructJob_result(rows *sqlx.Rows) interface{} {
	sa := make([]Job_result, 0)
	t := Job_result{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructJob_status(rows *sqlx.Rows) interface{} {
	sa := make([]Job_status, 0)
	t := Job_status{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructLog(rows *sqlx.Rows) interface{} {
	sa := make([]Log, 0)
	t := Log{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructParameter(rows *sqlx.Rows) interface{} {
	sa := make([]Parameter, 0)
	t := Parameter{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructPhys_location(rows *sqlx.Rows) interface{} {
	sa := make([]Phys_location, 0)
	t := Phys_location{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructProfile(rows *sqlx.Rows) interface{} {
	sa := make([]Profile, 0)
	t := Profile{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructProfile_parameter(rows *sqlx.Rows) interface{} {
	sa := make([]Profile_parameter, 0)
	t := Profile_parameter{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructRegex(rows *sqlx.Rows) interface{} {
	sa := make([]Regex, 0)
	t := Regex{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructRegion(rows *sqlx.Rows) interface{} {
	sa := make([]Region, 0)
	t := Region{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructRole(rows *sqlx.Rows) interface{} {
	sa := make([]Role, 0)
	t := Role{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructServer(rows *sqlx.Rows) interface{} {
	sa := make([]Server, 0)
	t := Server{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructServercheck(rows *sqlx.Rows) interface{} {
	sa := make([]Servercheck, 0)
	t := Servercheck{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructStaticdnsentry(rows *sqlx.Rows) interface{} {
	sa := make([]Staticdnsentry, 0)
	t := Staticdnsentry{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructStatus(rows *sqlx.Rows) interface{} {
	sa := make([]Status, 0)
	t := Status{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructTm_user(rows *sqlx.Rows) interface{} {
	sa := make([]Tm_user, 0)
	t := Tm_user{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructTo_extension(rows *sqlx.Rows) interface{} {
	sa := make([]To_extension, 0)
	t := To_extension{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructType(rows *sqlx.Rows) interface{} {
	sa := make([]Type, 0)
	t := Type{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructCondensedasn(rows *sqlx.Rows) interface{} {
	sa := make([]Condensedasn, 0)
	t := Condensedasn{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructCrystal(rows *sqlx.Rows) interface{} {
	sa := make([]Crystal, 0)
	t := Crystal{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructDeliveryservicesdenver(rows *sqlx.Rows) interface{} {
	sa := make([]Deliveryservicesdenver, 0)
	t := Deliveryservicesdenver{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
func EncodeStructHimark(rows *sqlx.Rows) interface{} {
	sa := make([]Himark, 0)
	t := Himark{}

	for rows.Next() {
		 rows.StructScan(&t)
		 sa = append(sa, t)
	}

	return sa
}
