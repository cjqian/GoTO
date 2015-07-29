package genStructs
import (
	"github.com/jmoiron/sqlx"
	"net/http"
)
func MapTableToJson(tableName string, rows *sqlx.Rows, w http.ResponseWriter) {
	if tableName == "asn"{
		EncodeStructAsn(rows, w)
	}
	if tableName == "cachegroup"{
		EncodeStructCachegroup(rows, w)
	}
	if tableName == "cachegroup_parameter"{
		EncodeStructCachegroup_parameter(rows, w)
	}
	if tableName == "deliveryservice"{
		EncodeStructDeliveryservice(rows, w)
	}
	if tableName == "deliveryservice_regex"{
		EncodeStructDeliveryservice_regex(rows, w)
	}
	if tableName == "deliveryservice_server"{
		EncodeStructDeliveryservice_server(rows, w)
	}
	if tableName == "deliveryservice_tmuser"{
		EncodeStructDeliveryservice_tmuser(rows, w)
	}
	if tableName == "division"{
		EncodeStructDivision(rows, w)
	}
	if tableName == "goose_db_version"{
		EncodeStructGoose_db_version(rows, w)
	}
	if tableName == "hwinfo"{
		EncodeStructHwinfo(rows, w)
	}
	if tableName == "job"{
		EncodeStructJob(rows, w)
	}
	if tableName == "job_agent"{
		EncodeStructJob_agent(rows, w)
	}
	if tableName == "job_result"{
		EncodeStructJob_result(rows, w)
	}
	if tableName == "job_status"{
		EncodeStructJob_status(rows, w)
	}
	if tableName == "log"{
		EncodeStructLog(rows, w)
	}
	if tableName == "parameter"{
		EncodeStructParameter(rows, w)
	}
	if tableName == "phys_location"{
		EncodeStructPhys_location(rows, w)
	}
	if tableName == "profile"{
		EncodeStructProfile(rows, w)
	}
	if tableName == "profile_parameter"{
		EncodeStructProfile_parameter(rows, w)
	}
	if tableName == "regex"{
		EncodeStructRegex(rows, w)
	}
	if tableName == "region"{
		EncodeStructRegion(rows, w)
	}
	if tableName == "role"{
		EncodeStructRole(rows, w)
	}
	if tableName == "server"{
		EncodeStructServer(rows, w)
	}
	if tableName == "servercheck"{
		EncodeStructServercheck(rows, w)
	}
	if tableName == "staticdnsentry"{
		EncodeStructStaticdnsentry(rows, w)
	}
	if tableName == "status"{
		EncodeStructStatus(rows, w)
	}
	if tableName == "tm_user"{
		EncodeStructTm_user(rows, w)
	}
	if tableName == "to_extension"{
		EncodeStructTo_extension(rows, w)
	}
	if tableName == "type"{
		EncodeStructType(rows, w)
	}
	if tableName == "condensedasn"{
		EncodeStructCondensedasn(rows, w)
	}
	if tableName == "deliveryservicesdenver"{
		EncodeStructDeliveryservicesdenver(rows, w)
	}
	if tableName == "himark"{
		EncodeStructHimark(rows, w)
	}
}
