package structs
type Asn struct {
	Id		string
	Asn		string
	Cachegroup		string
	Last_updated		string
}
type Cachegroup struct {
	Id		string
	Name		string
	Short_name		string
	Latitude		string
	Longitude		string
	Parent_cachegroup_id		string
	Type		string
	Last_updated		string
}
type Cachegroup_parameter struct {
	Cachegroup		string
	Parameter		string
	Last_updated		string
}
type Deliveryservice struct {
	Id		string
	Xml_id		string
	Active		string
	Dscp		string
	Signed		string
	Qstring_ignore		string
	Geo_limit		string
	Http_bypass_fqdn		string
	Dns_bypass_ip		string
	Dns_bypass_ip6		string
	Dns_bypass_ttl		string
	Org_server_fqdn		string
	Type		string
	Profile		string
	Ccr_dns_ttl		string
	Global_max_mbps		string
	Global_max_tps		string
	Long_desc		string
	Long_desc_1		string
	Long_desc_2		string
	Max_dns_answers		string
	Info_url		string
	Miss_lat		string
	Miss_long		string
	Check_path		string
	Last_updated		string
	Protocol		string
	Ssl_key_version		string
	Ipv6_routing_enabled		string
	Range_request_handling		string
	Edge_header_rewrite		string
	Origin_shield		string
	Mid_header_rewrite		string
	Regex_remap		string
	Cacheurl		string
	Remap_text		string
	Routing_name		string
}
type Deliveryservice_regex struct {
	Deliveryservice		string
	Regex		string
	Set_number		string
}
type Deliveryservice_server struct {
	Deliveryservice		string
	Server		string
	Last_updated		string
}
type Deliveryservice_tmuser struct {
	Deliveryservice		string
	Tm_user_id		string
	Last_updated		string
}
type Division struct {
	Id		string
	Name		string
	Last_updated		string
}
type Goose_db_version struct {
	Id		string
	Version_id		string
	Is_applied		string
	Tstamp		string
}
type Hwinfo struct {
	Id		string
	Serverid		string
	Description		string
	Val		string
	Last_updated		string
}
type Job struct {
	Id		string
	Agent		string
	Object_type		string
	Object_name		string
	Keyword		string
	Parameters		string
	Asset_url		string
	Asset_type		string
	Status		string
	Start_time		string
	Entered_time		string
	Job_user		string
	Last_updated		string
	Job_deliveryservice		string
}
type Job_agent struct {
	Id		string
	Name		string
	Description		string
	Active		string
	Last_updated		string
}
type Job_result struct {
	Id		string
	Job		string
	Agent		string
	Result		string
	Description		string
	Last_updated		string
}
type Job_status struct {
	Id		string
	Name		string
	Description		string
	Last_updated		string
}
type Log struct {
	Id		string
	Level		string
	Message		string
	Tm_user		string
	Ticketnum		string
	Last_updated		string
}
type Parameter struct {
	Id		string
	Name		string
	Config_file		string
	Value		string
	Last_updated		string
}
type Phys_location struct {
	Id		string
	Name		string
	Short_name		string
	Address		string
	City		string
	State		string
	Zip		string
	Poc		string
	Phone		string
	Email		string
	Comments		string
	Region		string
	Last_updated		string
}
type Profile struct {
	Id		string
	Name		string
	Description		string
	Last_updated		string
}
type Profile_parameter struct {
	Profile		string
	Parameter		string
	Last_updated		string
}
type Regex struct {
	Id		string
	Pattern		string
	Type		string
	Last_updated		string
}
type Region struct {
	Id		string
	Name		string
	Division		string
	Last_updated		string
}
type Role struct {
	Id		string
	Name		string
	Description		string
	Priv_level		string
}
type Server struct {
	Id		string
	Host_name		string
	Domain_name		string
	Tcp_port		string
	Xmpp_id		string
	Xmpp_passwd		string
	Interface_name		string
	Ip_address		string
	Ip_netmask		string
	Ip_gateway		string
	Ip6_address		string
	Ip6_gateway		string
	Interface_mtu		string
	Phys_location		string
	Rack		string
	Cachegroup		string
	Type		string
	Status		string
	Upd_pending		string
	Profile		string
	Mgmt_ip_address		string
	Mgmt_ip_netmask		string
	Mgmt_ip_gateway		string
	Ilo_ip_address		string
	Ilo_ip_netmask		string
	Ilo_ip_gateway		string
	Ilo_username		string
	Ilo_password		string
	Router_host_name		string
	Router_port_name		string
	Last_updated		string
}
type Servercheck struct {
	Id		string
	Server		string
	Aa		string
	Ab		string
	Ac		string
	Ad		string
	Ae		string
	Af		string
	Ag		string
	Ah		string
	Ai		string
	Aj		string
	Ak		string
	Al		string
	Am		string
	An		string
	Ao		string
	Ap		string
	Aq		string
	Ar		string
	As		string
	At		string
	Au		string
	Av		string
	Aw		string
	Ax		string
	Ay		string
	Az		string
	Ba		string
	Bb		string
	Bc		string
	Bd		string
	Be		string
	Last_updated		string
}
type Staticdnsentry struct {
	Id		string
	Host		string
	Address		string
	Type		string
	Ttl		string
	Deliveryservice		string
	Cachegroup		string
	Last_updated		string
}
type Status struct {
	Id		string
	Name		string
	Description		string
	Last_updated		string
}
type Tm_user struct {
	Id		string
	Username		string
	Role		string
	Uid		string
	Gid		string
	Local_passwd		string
	Confirm_local_passwd		string
	Last_updated		string
	Company		string
	Email		string
	Full_name		string
	New_user		string
	Address_line1		string
	Address_line2		string
	City		string
	State_or_province		string
	Phone_number		string
	Postal_code		string
	Country		string
	Local_user		string
	Token		string
	Registration_sent		string
}
type To_extension struct {
	Id		string
	Name		string
	Version		string
	Info_url		string
	Script_file		string
	Isactive		string
	Additional_config_json		string
	Description		string
	Servercheck_short_name		string
	Servercheck_column_name		string
	Type		string
	Last_updated		string
}
type Type struct {
	Id		string
	Name		string
	Description		string
	Use_in_table		string
	Last_updated		string
}
