//temporary hard-coded structs (one for each table in to_development database)
//eventually hope to have a script generate this file at runtime

package structs

func PrintHi() string {
	return "Hello, world!"
}

type AsnStruct struct {
	Id         int
	Asn        int
	Cachegroup int
	//originally type timestamp
	//last_updated string
}

type CachegroupStruct struct {
	id                   int
	name                 string
	short_name           string
	latitude             float64
	longitude            float64
	parent_cachegroup_id int
	//this is the "type" field in the database
	group_type   int
	last_updated string
}
