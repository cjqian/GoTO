package genStructs
import "github.com/jmoiron/sqlx"
func MapTableToJson(tableName string, rows *sqlx.Rows) interface{}{
	if tableName == "asn"{
		return EncodeStructAsn(rows)
	}
	if tableName == "cachegroup"{
		return EncodeStructCachegroup(rows)
	}
	if tableName == "cachegroup_parameter"{
		return EncodeStructCachegroup_parameter(rows)
	}
	if tableName == "deliveryservice"{
		return EncodeStructDeliveryservice(rows)
	}
	if tableName == "deliveryservice_regex"{
		return EncodeStructDeliveryservice_regex(rows)
	}
	if tableName == "deliveryservice_server"{
		return EncodeStructDeliveryservice_server(rows)
	}
	if tableName == "deliveryservice_tmuser"{
		return EncodeStructDeliveryservice_tmuser(rows)
	}
	if tableName == "division"{
		return EncodeStructDivision(rows)
	}
	if tableName == "goose_db_version"{
		return EncodeStructGoose_db_version(rows)
	}
	if tableName == "hwinfo"{
		return EncodeStructHwinfo(rows)
	}
	if tableName == "job"{
		return EncodeStructJob(rows)
	}
	if tableName == "job_agent"{
		return EncodeStructJob_agent(rows)
	}
	if tableName == "job_result"{
		return EncodeStructJob_result(rows)
	}
	if tableName == "job_status"{
		return EncodeStructJob_status(rows)
	}
	if tableName == "log"{
		return EncodeStructLog(rows)
	}
	if tableName == "parameter"{
		return EncodeStructParameter(rows)
	}
	if tableName == "phys_location"{
		return EncodeStructPhys_location(rows)
	}
	if tableName == "profile"{
		return EncodeStructProfile(rows)
	}
	if tableName == "profile_parameter"{
		return EncodeStructProfile_parameter(rows)
	}
	if tableName == "regex"{
		return EncodeStructRegex(rows)
	}
	if tableName == "region"{
		return EncodeStructRegion(rows)
	}
	if tableName == "role"{
		return EncodeStructRole(rows)
	}
	if tableName == "server"{
		return EncodeStructServer(rows)
	}
	if tableName == "servercheck"{
		return EncodeStructServercheck(rows)
	}
	if tableName == "staticdnsentry"{
		return EncodeStructStaticdnsentry(rows)
	}
	if tableName == "status"{
		return EncodeStructStatus(rows)
	}
	if tableName == "tm_user"{
		return EncodeStructTm_user(rows)
	}
	if tableName == "to_extension"{
		return EncodeStructTo_extension(rows)
	}
	if tableName == "type"{
		return EncodeStructType(rows)
	}
	if tableName == "condensedasn"{
		return EncodeStructCondensedasn(rows)
	}
	if tableName == "crystal"{
		return EncodeStructCrystal(rows)
	}
	if tableName == "deliveryservicesdenver"{
		return EncodeStructDeliveryservicesdenver(rows)
	}
	if tableName == "himark"{
		return EncodeStructHimark(rows)
	}
	return ""
}
