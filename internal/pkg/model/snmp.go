package model



// SNMP 连接参数（HTTP 请求体）
type SNMPRequest struct {
    IP              string `json:"ip" binding:"required,ipv4"`
    SNMPVersion     string `json:"snmp_version" binding:"required,oneof=v3"`
    Username        string `json:"username"`       // SNMPv3 用户名
    AuthPassword    string `json:"auth_password"`  // 认证密码
    AuthProtocol    string `json:"auth_protocol"`  // 认证协议（SHA/AES）
    PrivPassword    string `json:"priv_password"`  // 加密密码
    PrivProtocol    string `json:"priv_protocol"`  // 加密协议
}

// 设备信息（数据库存储）
type DeviceInfo struct {
    IP         string `json:"ip" db:"ip"`
    Hostname   string `json:"hostname" db:"hostname"`
    Model      string `json:"model" db:"model"`
    Vendor     string `json:"vendor" db:"vendor"`
    DeviceType string `json:"device_type" db:"device_type"`
}