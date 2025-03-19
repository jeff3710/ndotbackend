package snmp

import (
	"fmt"
	"strings"
	"time"

	"github.com/gosnmp/gosnmp"
	"github.com/jeff3710/ndot/config"
	"github.com/jeff3710/ndot/internal/pkg/model"
)

type SNMPClient struct {
	config *config.Config
}

func NewSNMPClient(cfg *config.Config) *SNMPClient {
	return &SNMPClient{
		config: cfg,
	}
}

// ConvertSNMPVersion 转换字符串类型的SNMP版本到gosnmp.SnmpVersion类型
func ConvertSNMPVersion(version string) (gosnmp.SnmpVersion, error) {
	switch version {
	case "v1":
		return gosnmp.Version1, nil
	case "v2c":
		return gosnmp.Version2c, nil
	case "v3":
		return gosnmp.Version3, nil
	default:
		return gosnmp.Version2c, fmt.Errorf("不支持的SNMP版本: %s", version)
	}
}

// ConvertAuthProtocol 转换SNMPv3认证协议字符串到枚举值
// 支持的协议: MD5, SHA系列算法
// 返回错误当协议不受支持时
func ConvertAuthProtocol(proto string) (gosnmp.SnmpV3AuthProtocol, error) {
	switch strings.ToUpper(proto) {
	case "MD5":
		return gosnmp.MD5, nil
	case "SHA":
		return gosnmp.SHA, nil
	case "SHA256":
		return gosnmp.SHA256, nil
	case "SHA384":
		return gosnmp.SHA384, nil
	case "SHA512":
		return gosnmp.SHA512, nil
	default:
		return gosnmp.NoAuth, fmt.Errorf("不支持的认证协议: %s，支持的协议有: MD5, SHA, SHA256, SHA384, SHA512", proto)
	}
}

// ConvertPrivProtocol 转换SNMPv3加密协议字符串到枚举值
// 支持DES和AES系列加密算法
// 返回错误当协议无效时
func ConvertPrivProtocol(proto string) (gosnmp.SnmpV3PrivProtocol, error) {
	switch strings.ToUpper(proto) {
	case "DES":
		return gosnmp.DES, nil
	case "AES":
		return gosnmp.AES, nil
	case "AES192":
		return gosnmp.AES192, nil
	case "AES256":
		return gosnmp.AES256, nil
	case "AES192C":
		return gosnmp.AES192C, nil
	case "AES256C":
		return gosnmp.AES256C, nil
	default:
		return gosnmp.NoPriv, fmt.Errorf("不支持的加密协议: %s，支持的协议有: DES, AES, AES192, AES256, AES192C, AES256C", proto)
	}
}

func detectVendor(descr string, sysObjectID string,vendorOIDs map[string]string) string {
	if v := detectVendorFromOID(vendorOIDs,sysObjectID); v != "Unknown" {
		return v
	}
	return detectVendorFromDescription(descr)
}

func detectVendorFromOID(vendorOIDs map[string]string,sysObjectID string) string {
	for prefix, vendor := range vendorOIDs {
		if strings.HasPrefix(sysObjectID, prefix) {
			return vendor
		}
	}
	return "Unknown"
}

func extractModel(descr string) string {
	// 典型sysDescr格式示例：
	// "Cisco IOS Software, C3750E Software (C3750E-UNIVERSALK9-M), Version 15.0(2)SE11"
	parts := strings.Split(descr, ",")
	if len(parts) >= 2 {
		return strings.TrimSpace(parts[1])
	}
	return descr
}

func detectVendorFromDescription(descr string) string {
	descr = strings.ToLower(descr)
	switch {
	case strings.Contains(descr, "cisco") || strings.Contains(descr, "nexus"):
		return "Cisco"
	case strings.Contains(descr, "huawei") || strings.Contains(descr, "hwe"):
		return "Huawei"
	case strings.Contains(descr, "junos") || strings.Contains(descr, "jnpr"):
		return "Juniper"
	case strings.Contains(descr, "h3c") || strings.Contains(descr, "comware"):
		return "H3C"
    case strings.Contains(descr,"sangfor"):
		return "Sangfor"
	case strings.Contains(descr, "fortinet"):
		return "Fortinet"
	default:
		return ""
	}
}

func (c *SNMPClient) GetDeviceInfo(req *model.SNMPRequest) (*model.DeviceInfo, error) {
	// 配置 SNMPv3 参数
	version, err := ConvertSNMPVersion(req.SNMPVersion)
	if err != nil {
		return nil, err
	}
	authProtocol, err := ConvertAuthProtocol(req.AuthProtocol)
	if err != nil {
		return nil, err
	}
	privProtocol, err := ConvertPrivProtocol(req.PrivProtocol)
	if err != nil {
		return nil, err
	}
	snmp := &gosnmp.GoSNMP{
		Target:        req.IP,
		Port:          161,
		Version:       version,
		SecurityModel: gosnmp.UserSecurityModel,
		MsgFlags:      gosnmp.AuthPriv,
		SecurityParameters: &gosnmp.UsmSecurityParameters{
			UserName:                 req.Username,
			AuthenticationProtocol:   authProtocol,
			AuthenticationPassphrase: req.AuthPassword,
			PrivacyProtocol:          privProtocol,
			PrivacyPassphrase:        req.PrivPassword,
		},
		Timeout: 10 * time.Second,
	}

	// 连接设备
	if err := snmp.Connect(); err != nil {
		return nil, fmt.Errorf("SNMP connect failed: %v", err)
	}
	defer snmp.Conn.Close()

	oids := make([]string, 0, len(c.config.SystemOIDs))
	for oid := range c.config.SystemOIDs {
		oids = append(oids, oid)
	}
	fmt.Println(oids)

	result, err := snmp.Get(oids)
	if err != nil {
		return nil, fmt.Errorf("SNMP get failed: %v", err)
	}

    // vendorOIDs 厂商OID前缀映射表，从配置文件加载
// 键为OID前缀，值为厂商名称
 vendorOIDs := c.config.Vendors // 从配置文件加载

	// 解析结果（需根据实际设备调整）
	info := &model.DeviceInfo{IP: req.IP}
    oidMap:= c.config.SystemOIDs
	for _, v := range result.Variables {
		switch v.Name {
		case oidMap["sysName"]:
			info.Hostname = string(v.Value.([]byte))
		case oidMap["sysDescr"]:
			// 从描述中提取型号、厂商等信息（示例）
            sysDescr := string(v.Value.([]byte))
			info.Vendor = detectVendor(sysDescr, v.Value.(string), vendorOIDs)
			info.Model = extractModel(sysDescr)
            case oidMap["sysObjectID"]:
			sysObjectID := v.Value.(string)
			if info.Vendor == "" {
				info.Vendor = detectVendorFromOID(vendorOIDs,sysObjectID)
			}
         
            // case oidMap["serialNumber"]:
            //     serialNumber:= string(v.Value.([]byte))
            default:
			// 其他OID的处理
		}
	}

	return info, nil
}
