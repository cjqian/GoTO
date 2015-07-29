package genStructs
type Asn struct {
	Id		int32 `json:"id"`
	Asn		int32 `json:"asn"`
	Cachegroup		int32 `json:"cachegroup"`
	Last_updated		string `json:"last_updated"`
}
type Cachegroup struct {
	Id		int32 `json:"id"`
	Name		string `json:"name"`
	Short_name		string `json:"short_name"`
	Latitude		float64 `json:"latitude"`
	Longitude		float64 `json:"longitude"`
	Parent_cachegroup_id		int32 `json:"parent_cachegroup_id"`
	Type		int32 `json:"type"`
	Last_updated		string `json:"last_updated"`
}
type Cachegroup_parameter struct {
	Cachegroup		int32 `json:"cachegroup"`
	Parameter		int32 `json:"parameter"`
	Last_updated		string `json:"last_updated"`
}
type Deliveryservice struct {
	Id		int32 `json:"id"`
	Xml_id		string `json:"xml_id"`
	Active		uint8 `json:"active"`
	Dscp		int32 `json:"dscp"`
	Signed		uint8 `json:"signed"`
	Qstring_ignore		uint8 `json:"qstring_ignore"`
	Geo_limit		uint8 `json:"geo_limit"`
	Http_bypass_fqdn		string `json:"http_bypass_fqdn"`
	Dns_bypass_ip		string `json:"dns_bypass_ip"`
	Dns_bypass_ip6		string `json:"dns_bypass_ip6"`
	Dns_bypass_ttl		int32 `json:"dns_bypass_ttl"`
	Org_server_fqdn		string `json:"org_server_fqdn"`
	Type		int32 `json:"type"`
	Profile		int32 `json:"profile"`
	Ccr_dns_ttl		int32 `json:"ccr_dns_ttl"`
	Global_max_mbps		int32 `json:"global_max_mbps"`
	Global_max_tps		int32 `json:"global_max_tps"`
	Long_desc		string `json:"long_desc"`
	Long_desc_1		string `json:"long_desc_1"`
	Long_desc_2		string `json:"long_desc_2"`
	Max_dns_answers		int32 `json:"max_dns_answers"`
	Info_url		string `json:"info_url"`
	Miss_lat		float64 `json:"miss_lat"`
	Miss_long		float64 `json:"miss_long"`
	Check_path		string `json:"check_path"`
	Last_updated		string `json:"last_updated"`
	Protocol		uint8 `json:"protocol"`
	Ssl_key_version		int32 `json:"ssl_key_version"`
	Ipv6_routing_enabled		uint8 `json:"ipv6_routing_enabled"`
	Range_request_handling		uint8 `json:"range_request_handling"`
	Edge_header_rewrite		string `json:"edge_header_rewrite"`
	Origin_shield		string `json:"origin_shield"`
	Mid_header_rewrite		string `json:"mid_header_rewrite"`
	Regex_remap		string `json:"regex_remap"`
	Cacheurl		string `json:"cacheurl"`
	Remap_text		string `json:"remap_text"`
	Multi_site_origin		uint8 `json:"multi_site_origin"`
}
type Deliveryservice_regex struct {
	Deliveryservice		int32 `json:"deliveryservice"`
	Regex		int32 `json:"regex"`
	Set_number		int32 `json:"set_number"`
}
type Deliveryservice_server struct {
	Deliveryservice		int32 `json:"deliveryservice"`
	Server		int32 `json:"server"`
	Last_updated		string `json:"last_updated"`
}
type Deliveryservice_tmuser struct {
	Deliveryservice		int32 `json:"deliveryservice"`
	Tm_user_id		int32 `json:"tm_user_id"`
	Last_updated		string `json:"last_updated"`
}
type Division struct {
	Id		int32 `json:"id"`
	Name		string `json:"name"`
	Last_updated		string `json:"last_updated"`
}
type Goose_db_version struct {
	Id		int64 `json:"id"`
	Version_id		int64 `json:"version_id"`
	Is_applied		uint8 `json:"is_applied"`
	Tstamp		string `json:"tstamp"`
}
type Hwinfo struct {
	Id		int32 `json:"id"`
	Serverid		int32 `json:"serverid"`
	Description		string `json:"description"`
	Val		string `json:"val"`
	Last_updated		string `json:"last_updated"`
}
type Job struct {
	Id		int32 `json:"id"`
	Agent		int32 `json:"agent"`
	Object_type		string `json:"object_type"`
	Object_name		string `json:"object_name"`
	Keyword		string `json:"keyword"`
	Parameters		string `json:"parameters"`
	Asset_url		string `json:"asset_url"`
	Asset_type		string `json:"asset_type"`
	Status		int32 `json:"status"`
	Start_time		string `json:"start_time"`
	Entered_time		string `json:"entered_time"`
	Job_user		int32 `json:"job_user"`
	Last_updated		string `json:"last_updated"`
	Job_deliveryservice		int32 `json:"job_deliveryservice"`
}
type Job_agent struct {
	Id		int32 `json:"id"`
	Name		string `json:"name"`
	Description		string `json:"description"`
	Active		int32 `json:"active"`
	Last_updated		string `json:"last_updated"`
}
type Job_result struct {
	Id		int32 `json:"id"`
	Job		int32 `json:"job"`
	Agent		int32 `json:"agent"`
	Result		string `json:"result"`
	Description		string `json:"description"`
	Last_updated		string `json:"last_updated"`
}
type Job_status struct {
	Id		int32 `json:"id"`
	Name		string `json:"name"`
	Description		string `json:"description"`
	Last_updated		string `json:"last_updated"`
}
type Log struct {
	Id		int32 `json:"id"`
	Level		string `json:"level"`
	Message		string `json:"message"`
	Tm_user		int32 `json:"tm_user"`
	Ticketnum		string `json:"ticketnum"`
	Last_updated		string `json:"last_updated"`
}
type Parameter struct {
	Id		int32 `json:"id"`
	Name		string `json:"name"`
	Config_file		string `json:"config_file"`
	Value		string `json:"value"`
	Last_updated		string `json:"last_updated"`
}
type Phys_location struct {
	Id		int32 `json:"id"`
	Name		string `json:"name"`
	Short_name		string `json:"short_name"`
	Address		string `json:"address"`
	City		string `json:"city"`
	State		string `json:"state"`
	Zip		string `json:"zip"`
	Poc		string `json:"poc"`
	Phone		string `json:"phone"`
	Email		string `json:"email"`
	Comments		string `json:"comments"`
	Region		int32 `json:"region"`
	Last_updated		string `json:"last_updated"`
}
type Profile struct {
	Id		int32 `json:"id"`
	Name		string `json:"name"`
	Description		string `json:"description"`
	Last_updated		string `json:"last_updated"`
}
type Profile_parameter struct {
	Profile		int32 `json:"profile"`
	Parameter		int32 `json:"parameter"`
	Last_updated		string `json:"last_updated"`
}
type Regex struct {
	Id		int32 `json:"id"`
	Pattern		string `json:"pattern"`
	Type		int32 `json:"type"`
	Last_updated		string `json:"last_updated"`
}
type Region struct {
	Id		int32 `json:"id"`
	Name		string `json:"name"`
	Division		int32 `json:"division"`
	Last_updated		string `json:"last_updated"`
}
type Role struct {
	Id		int32 `json:"id"`
	Name		string `json:"name"`
	Description		string `json:"description"`
	Priv_level		int32 `json:"priv_level"`
}
type Server struct {
	Id		int32 `json:"id"`
	Host_name		string `json:"host_name"`
	Domain_name		string `json:"domain_name"`
	Tcp_port		int32 `json:"tcp_port"`
	Xmpp_id		string `json:"xmpp_id"`
	Xmpp_passwd		string `json:"xmpp_passwd"`
	Interface_name		string `json:"interface_name"`
	Ip_address		string `json:"ip_address"`
	Ip_netmask		string `json:"ip_netmask"`
	Ip_gateway		string `json:"ip_gateway"`
	Ip6_address		string `json:"ip6_address"`
	Ip6_gateway		string `json:"ip6_gateway"`
	Interface_mtu		int32 `json:"interface_mtu"`
	Phys_location		int32 `json:"phys_location"`
	Rack		string `json:"rack"`
	Cachegroup		int32 `json:"cachegroup"`
	Type		int32 `json:"type"`
	Status		int32 `json:"status"`
	Upd_pending		uint8 `json:"upd_pending"`
	Profile		int32 `json:"profile"`
	Mgmt_ip_address		string `json:"mgmt_ip_address"`
	Mgmt_ip_netmask		string `json:"mgmt_ip_netmask"`
	Mgmt_ip_gateway		string `json:"mgmt_ip_gateway"`
	Ilo_ip_address		string `json:"ilo_ip_address"`
	Ilo_ip_netmask		string `json:"ilo_ip_netmask"`
	Ilo_ip_gateway		string `json:"ilo_ip_gateway"`
	Ilo_username		string `json:"ilo_username"`
	Ilo_password		string `json:"ilo_password"`
	Router_host_name		string `json:"router_host_name"`
	Router_port_name		string `json:"router_port_name"`
	Last_updated		string `json:"last_updated"`
}
type Servercheck struct {
	Id		int32 `json:"id"`
	Server		int32 `json:"server"`
	Aa		int32 `json:"aa"`
	Ab		int32 `json:"ab"`
	Ac		int32 `json:"ac"`
	Ad		int32 `json:"ad"`
	Ae		int32 `json:"ae"`
	Af		int32 `json:"af"`
	Ag		int32 `json:"ag"`
	Ah		int32 `json:"ah"`
	Ai		int32 `json:"ai"`
	Aj		int32 `json:"aj"`
	Ak		int32 `json:"ak"`
	Al		int32 `json:"al"`
	Am		int32 `json:"am"`
	An		int32 `json:"an"`
	Ao		int32 `json:"ao"`
	Ap		int32 `json:"ap"`
	Aq		int32 `json:"aq"`
	Ar		int32 `json:"ar"`
	As		int32 `json:"as"`
	At		int32 `json:"at"`
	Au		int32 `json:"au"`
	Av		int32 `json:"av"`
	Aw		int32 `json:"aw"`
	Ax		int32 `json:"ax"`
	Ay		int32 `json:"ay"`
	Az		int32 `json:"az"`
	Ba		int32 `json:"ba"`
	Bb		int32 `json:"bb"`
	Bc		int32 `json:"bc"`
	Bd		int32 `json:"bd"`
	Be		int32 `json:"be"`
	Last_updated		string `json:"last_updated"`
}
type Staticdnsentry struct {
	Id		int32 `json:"id"`
	Host		string `json:"host"`
	Address		string `json:"address"`
	Type		int32 `json:"type"`
	Ttl		int32 `json:"ttl"`
	Deliveryservice		int32 `json:"deliveryservice"`
	Cachegroup		int32 `json:"cachegroup"`
	Last_updated		string `json:"last_updated"`
}
type Status struct {
	Id		int32 `json:"id"`
	Name		string `json:"name"`
	Description		string `json:"description"`
	Last_updated		string `json:"last_updated"`
}
type Tm_user struct {
	Id		int32 `json:"id"`
	Username		string `json:"username"`
	Role		int32 `json:"role"`
	Uid		int32 `json:"uid"`
	Gid		int32 `json:"gid"`
	Local_passwd		string `json:"local_passwd"`
	Confirm_local_passwd		string `json:"confirm_local_passwd"`
	Last_updated		string `json:"last_updated"`
	Company		string `json:"company"`
	Email		string `json:"email"`
	Full_name		string `json:"full_name"`
	New_user		uint8 `json:"new_user"`
	Address_line1		string `json:"address_line1"`
	Address_line2		string `json:"address_line2"`
	City		string `json:"city"`
	State_or_province		string `json:"state_or_province"`
	Phone_number		string `json:"phone_number"`
	Postal_code		string `json:"postal_code"`
	Country		string `json:"country"`
	Local_user		uint8 `json:"local_user"`
	Token		string `json:"token"`
	Registration_sent		string `json:"registration_sent"`
}
type To_extension struct {
	Id		int32 `json:"id"`
	Name		string `json:"name"`
	Version		string `json:"version"`
	Info_url		string `json:"info_url"`
	Script_file		string `json:"script_file"`
	Isactive		uint8 `json:"isactive"`
	Additional_config_json		string `json:"additional_config_json"`
	Description		string `json:"description"`
	Servercheck_short_name		string `json:"servercheck_short_name"`
	Servercheck_column_name		string `json:"servercheck_column_name"`
	Type		int32 `json:"type"`
	Last_updated		string `json:"last_updated"`
}
type Type struct {
	Id		int32 `json:"id"`
	Name		string `json:"name"`
	Description		string `json:"description"`
	Use_in_table		string `json:"use_in_table"`
	Last_updated		string `json:"last_updated"`
}
type Condensedasn struct {
	Id		int32 `json:"id"`
	Asn		int32 `json:"asn"`
	Cachegroup		int32 `json:"cachegroup"`
	Last_updated		string `json:"last_updated"`
}
type Crystal struct {
	Host_name		string `json:"host_name"`
	Id		int32 `json:"id"`
}
type Deliveryservicesdenver struct {
	Host_name		string `json:"host_name"`
	Xml_id		string `json:"xml_id"`
}
type Himark struct {
	Xml_id		string `json:"xml_id"`
	Host_name		string `json:"host_name"`
}
