package db

import (
	"context"
	"fmt"
)

type SnmpTemplateUnion struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	UserID      int32  `json:"user_id"`
	Protocol    string `json:"protocol"`
	Version     string `json:"version"`
	Description string `json:"description"`
	DeviceCount int32  `json:"device_count"`

	TemplateID     int32  `json:"template_id"`
	Port           string `json:"port"`
	ReadCommunity  string `json:"read_community"`
	WriteCommunity string `json:"write_community"`
	TrapCommunity  string `json:"trap_community"`

	SecurityLevel string `json:"security_level"`
	AuthProtocol  string `json:"auth_protocol"`
	AuthPassword  string `json:"auth_password"`
	PrivProtocol  string `json:"priv_protocol"`
	PrivPassword  string `json:"priv_password"`
	V3User        string `json:"v3_user"`
	EngineID      string `json:"engine_id"`
}

// CreateSNMPTemplate 创建一个SNMP模板
// 使用了联合表，需要根据version字段来判断是v2c还是v3
// 通过执行数据库事务来完成创建
// 先插入一条base记录，获取到id,然后根据version字段来插入v2c或v3记录，并关联到base记录
func (store *SQLStore) CreateSNMPTempate(ctx context.Context, template SnmpTemplateUnion) error {
	req := CreateSnmpTemplateBaseParams{
		Name:        template.Name,
		UserID:      template.UserID,
		Protocol:    template.Protocol,
		Version:     template.Version,
		Description: template.Description,
		DeviceCount: template.DeviceCount,
	}

	return store.execTx(ctx, func(q *Queries) error {
		baseID, err := q.CreateSnmpTemplateBase(ctx, req)
		if err != nil {
			return err
		}

		reqv2 := CreateSnmpTemplateWithV2ParametersParams{
			TemplateID:     int32(baseID),
			Port:           template.Port,
			ReadCommunity:  template.ReadCommunity,
			WriteCommunity: template.WriteCommunity,
			TrapCommunity:  template.TrapCommunity,
		}
		reqv3 := CreateSnmpTemplateWithV3ParametersParams{
			TemplateID:    int32(baseID),
			Port:          template.Port,
			SecurityLevel: template.SecurityLevel,
			AuthProtocol:  template.AuthProtocol,
			AuthPassword:  template.AuthPassword,
			PrivProtocol:  template.PrivProtocol,
			PrivPassword:  template.PrivPassword,
			V3User:        template.V3User,
			EngineID:      template.EngineID,
		}
		switch template.Version {
		case "v2c":
			if template.ReadCommunity == "" || template.WriteCommunity == "" || template.TrapCommunity == "" {
				return fmt.Errorf("missing required fields for v2c template")
			}
			return q.CreateSnmpTemplateWithV2Parameters(ctx, reqv2)
		case "v3":
			if template.SecurityLevel == "" || template.AuthProtocol == "" || template.AuthPassword == "" || template.PrivProtocol == "" || template.PrivPassword == "" || template.V3User == "" || template.EngineID == "" {
				return fmt.Errorf("missing required fields for v3 template")
			}
			return q.CreateSnmpTemplateWithV3Parameters(ctx, reqv3)
		default:
			return fmt.Errorf("invalid SNMP version")
		}

	})
}

// UpdateSNMPTemplate 用于更新SNMP模版的事务函数
func (store *SQLStore) UpdateSNMPTemplate(ctx context.Context, template SnmpTemplateUnion) error {
	return store.execTx(ctx, func(q *Queries) error {
		err:=q.UpdateSnmpTemplateBase(ctx, template.ID)
		if err!= nil {
			return err
		}
		switch template.Version {
		case "v2c":
			reqv2 := UpdateSnmpTemplateWithV2ParametersParams{
				Port:           template.Port,
				ReadCommunity:  template.ReadCommunity,
				WriteCommunity: template.WriteCommunity,
				TrapCommunity:  template.TrapCommunity,
			}
			return q.UpdateSnmpTemplateWithV2Parameters(ctx, reqv2)
			case "v3":
				reqv3 := UpdateSnmpTemplateWithV3ParametersParams{
					Port:          template.Port,
					SecurityLevel: template.SecurityLevel,
					AuthProtocol:  template.AuthProtocol,
					AuthPassword:  template.AuthPassword,
					PrivProtocol:  template.PrivProtocol,
					PrivPassword:  template.PrivPassword,
					V3User:        template.V3User,
					EngineID:      template.EngineID,
				}
				return q.UpdateSnmpTemplateWithV3Parameters(ctx, reqv3)
				default:
					return fmt.Errorf("invalid SNMP version")
		}
	})
}


// DeleteSNMPTemplate 删除一个SNMP模板
// GetSNMPTemplate 获取SNMP模板