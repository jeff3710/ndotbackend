package model

// SNMP 连接参数（HTTP 请求体）
type SNMPRequest struct {
	IP                     string `json:"ip" binding:"required,ip"`
	SNMPVersion            string `json:"snmp_version" binding:"required,oneof=v2c v3"`
	Community              string `json:"community,omitempty" binding:"required_if=SNMPVersion v2c"`
	Username               string `json:"username,omitempty" binding:"required_if=SNMPVersion v3"`
	AuthenticationProtocol string `json:"authentication_protocol,omitempty" binding:"required_if=SNMPVersion v3,oneof=MD5 SHA"`
	AuthenticationPassword string `json:"auth_password,omitempty" binding:"required_if=SNMPVersion v3"`
	PrivacyProtocol        string `json:"privacy_protocol,omitempty" binding:"required_if=SNMPVersion v3,oneof=DES AES"`
	PrivacyPassword        string `json:"privacy_password,omitempty" binding:"required_if=SNMPVersion v3"`
}

// 设备信息（数据库存储）
type DeviceInfo struct {
	IP         string `json:"ip" db:"ip"`
	Hostname   string `json:"hostname" db:"hostname"`
	Model      string `json:"model" db:"model"`
	Vendor     string `json:"vendor" db:"vendor"`
	DeviceType string `json:"device_type" db:"device_type"`
}
