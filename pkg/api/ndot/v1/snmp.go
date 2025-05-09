package v1


type SnmpTemplateRequest struct {
	Name        string `json:"name" binding:"required"`
	Protocol    string `json:"protocol" default:"snmp"`
	Version     string `json:"version" binding:"required"`
	Description string `json:"description"`
	DeviceCount int32  `json:"device_count"`

	Port           int32  `json:"port"`
	ReadCommunity  string `json:"read_community"`
	WriteCommunity string `json:"write_community"`
	TrapCommunity  string `json:"trap_community"`
	Timeout        int32  `json:"timeout"`
	PollInterval   int32  `json:"poll_interval"`
	Retries        int32  `json:"retries"`

	SecurityLevel string `json:"security_level"`
	AuthProtocol  string `json:"auth_protocol"`
	AuthPassword  string `json:"auth_password"`
	PrivProtocol  string `json:"priv_protocol"`
	PrivPassword  string `json:"priv_password"`
	V3User        string `json:"v3_user"`
	EngineID      string `json:"engine_id"`
}