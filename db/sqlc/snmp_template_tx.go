package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jeff3710/ndot/pkg/log"
)

// 定义版本常量
const (
	SNMPv2c = "v2c"
	SNMPv3  = "v3"
)

// SnmpTemplateUnion 是一个联合表，用于存储SNMP模板的基本信息和v2c和v3参数
type SnmpTemplateUnion struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	UserID      int32  `json:"user_id"`
	Protocol    string `json:"protocol"`
	Version     string `json:"version"`
	Description string `json:"description"`
	DeviceCount int32  `json:"device_count"`

	TemplateID     int32  `json:"template_id"`
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

	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func createBaseParams(template SnmpTemplateUnion) CreateSnmpTemplateBaseParams {
	return CreateSnmpTemplateBaseParams{
		Name:        template.Name,
		UserID:      template.UserID,
		Protocol:    template.Protocol,
		Version:     template.Version,
		Description: template.Description,
		DeviceCount: template.DeviceCount,
	}
}
func createV2Params(baseID int32, template SnmpTemplateUnion) CreateSnmpV2TemplateParams {
	return CreateSnmpV2TemplateParams{
		TemplateID:     baseID,
		Port:           template.Port,
		ReadCommunity:  template.ReadCommunity,
		WriteCommunity: template.WriteCommunity,
		TrapCommunity:  template.TrapCommunity,
		Timeout:        template.Timeout,
		PollInterval:   template.PollInterval,
		Retries:        template.Retries,
	}
}
func createV3Params(baseID int32, template SnmpTemplateUnion) CreateSnmpV3TemplateParams {
	return CreateSnmpV3TemplateParams{
		TemplateID:    baseID,
		Port:          template.Port,
		SecurityLevel: template.SecurityLevel,
		AuthProtocol:  template.AuthProtocol,
		AuthPassword:  template.AuthPassword,
		PrivProtocol:  template.PrivProtocol,
		PrivPassword:  template.PrivPassword,
		V3User:        template.V3User,
		EngineID:      template.EngineID,
		Timeout:       template.Timeout,
		PollInterval:  template.PollInterval,
		Retries:       template.Retries,
	}
}

// CreateSNMPTemplate 创建一个SNMP模板
// 使用了联合表，需要根据version字段来判断是v2c还是v3
// 通过执行数据库事务来完成创建
// 先插入一条base记录，获取到id,然后根据version字段来插入v2c或v3记录，并关联到base记录
func (store *SQLStore) CreateSNMPTempate(ctx context.Context, template SnmpTemplateUnion) error {
	req := createBaseParams(template)
	log.Infow("create snmp template base", "req", req)
	return store.execTx(ctx, func(q *Queries) error {
		baseID, err := q.CreateSnmpTemplateBase(ctx, req)
		log.Infow("create snmp template base", "baseID", baseID)
		if err != nil {
			return fmt.Errorf("failed to create base template: %w", err)
		}
		log.Infow("snmp version is ", template.Version)
		switch template.Version {
		case SNMPv2c:
			reqv2 := createV2Params(baseID, template)
			log.Infow("create snmp template v2", "reqv2", reqv2)
			return q.CreateSnmpV2Template(ctx, reqv2)
		case SNMPv3:
			reqv3 := createV3Params(baseID, template)
			return q.CreateSnmpV3Template(ctx, reqv3)
		default:
			return fmt.Errorf("invalid SNMP version")
		}

	})
}

// UpdateSNMPTemplate 用于更新SNMP模版的事务函数
func (store *SQLStore) UpdateSNMPTemplate(ctx context.Context, template SnmpTemplateUnion) error {
	req := UpdateSnmpTemplateBaseParams{
		ID:          template.TemplateID,
		Name:        template.Name,
		Protocol:    template.Protocol,
		Version:     template.Version,
		Description: template.Description,
	}
	return store.execTx(ctx, func(q *Queries) error {
		err := q.UpdateSnmpTemplateBase(ctx, req)
		if err != nil {
			return err
		}
		switch template.Version {
		case SNMPv2c:
			reqv2 := UpdateSnmpV2ParametersParams{
				Port:           template.Port,
				ReadCommunity:  template.ReadCommunity,
				WriteCommunity: template.WriteCommunity,
				TrapCommunity:  template.TrapCommunity,
				Timeout:        template.Timeout,
				PollInterval:   template.PollInterval,
				Retries:        template.Retries,
				TemplateID:     template.TemplateID,
			}
			return q.UpdateSnmpV2Parameters(ctx, reqv2)
		case SNMPv3:
			reqv3 := UpdateSnmpV3ParametersParams{
				Port:          template.Port,
				SecurityLevel: template.SecurityLevel,
				AuthProtocol:  template.AuthProtocol,
				AuthPassword:  template.AuthPassword,
				PrivProtocol:  template.PrivProtocol,
				PrivPassword:  template.PrivPassword,
				V3User:        template.V3User,
				EngineID:      template.EngineID,
				Timeout:       template.Timeout,
				PollInterval:  template.PollInterval,
				Retries:       template.Retries,
				TemplateID:    template.TemplateID,
			}
			return q.UpdateSnmpV3Parameters(ctx, reqv3)
		default:
			return fmt.Errorf("invalid SNMP version")
		}
	})
}

// DeleteSNMPTemplate 删除一个SNMP模板
func (store *SQLStore) DeleteSNMPTemplate(ctx context.Context, id int32) error {
	return store.execTx(ctx, func(q *Queries) error {
		err := q.DeleteSnmpTemplateBase(ctx, id)
		if err != nil {
			return err
		}

		// 根据模板版本删除对应的参数表
		template, err := q.GetSnmpTemplateBase(ctx, id)
		if err != nil {
			return err
		}

		switch template.Version {
		case SNMPv2c:
			return q.DeleteSnmpV2Parameters(ctx, id)
		case SNMPv3:
			return q.DeleteSnmpV3Parameters(ctx, id)
		default:
			return fmt.Errorf("invalid SNMP version")
		}
	})
}

// assignSnmpParams 用于根据版本为模板分配SNMP参数
// 从v2Map和v3Map中获取对应的参数，并赋值给template
// 必须使用指针类型，因为在函数内部修改了template的值
func assignSnmpParams(template *SnmpTemplateUnion, version string, v2Map map[int32]ListSnmpV2ParametersRow, v3Map map[int32]ListSnmpV3ParametersRow) {
	switch version {
	case SNMPv2c:
		if v2, ok := v2Map[template.ID]; ok {
			template.Port = v2.Port
			template.ReadCommunity = v2.ReadCommunity
			template.WriteCommunity = v2.WriteCommunity
			template.TrapCommunity = v2.TrapCommunity
			template.Timeout = v2.Timeout
			template.PollInterval = v2.PollInterval
			template.Retries = v2.Retries
		}
	case SNMPv3:
		if v3, ok := v3Map[template.ID]; ok {
			template.Port = v3.Port
			template.SecurityLevel = v3.SecurityLevel
			template.AuthProtocol = v3.AuthProtocol
			template.AuthPassword = v3.AuthPassword
			template.PrivProtocol = v3.PrivProtocol
			template.PrivPassword = v3.PrivPassword
			template.V3User = v3.V3User
			template.EngineID = v3.EngineID
			template.Timeout = v3.Timeout
			template.PollInterval = v3.PollInterval
			template.Retries = v3.Retries
		}
	}
}

// ListSNMPTemplates 获取分页SNMP模板列表
func (store *SQLStore) ListSNMPTemplates(ctx context.Context, limit int32, offset int32) ([]SnmpTemplateUnion, error) {
	var templates []SnmpTemplateUnion

	err := store.execTx(ctx, func(q *Queries) error {
		// 1. 获取基础模板列表
		bases, err := q.ListSnmpTemplatesBase(ctx, ListSnmpTemplatesBaseParams{
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return fmt.Errorf("failed to list snmp templates: %w", err)
		}

		// 2. 收集所有模板ID
		var templateIDs []int32
		for _, base := range bases {
			templateIDs = append(templateIDs, base.ID)
		}

		// 3. 批量获取v2和v3参数
		v2Params, err := q.ListSnmpV2Parameters(ctx, templateIDs)
		if err != nil {
			return fmt.Errorf("failed to list snmp v2 parameters: %w", err)
		}
		v3Params, err := q.ListSnmpV3Parameters(ctx, templateIDs)
		if err != nil {
			return fmt.Errorf("failed to list snmp v3 parameters: %w", err)
		}

		// 4. 创建参数映射表
		v2Map := make(map[int32]ListSnmpV2ParametersRow)
		for _, param := range v2Params {
			v2Map[param.TemplateID] = param
		}

		v3Map := make(map[int32]ListSnmpV3ParametersRow)
		for _, param := range v3Params {
			v3Map[param.TemplateID] = param
		}

		// 5. 组装完整模板数据
		for _, base := range bases {
			template := SnmpTemplateUnion{
				ID:          base.ID,
				Name:        base.Name,
				UserID:      base.UserID,
				Protocol:    base.Protocol,
				Version:     base.Version,
				Description: base.Description,
				DeviceCount: base.DeviceCount,
				CreatedAt:   base.CreatedAt,
				UpdatedAt:   base.UpdatedAt,
			}

			assignSnmpParams(&template, base.Version, v2Map, v3Map)

			templates = append(templates, template)
		}

		return nil
	})

	return templates, err
}

// getV2TemplateParams 获取v2模板参数
func (store *SQLStore) getV2TemplateParams(ctx context.Context, q *Queries, template *SnmpTemplateUnion, id int32) error {
	v2, err := q.GetSnmpV2Parameters(ctx, id)
	if err != nil {
		return err
	}
	template.Port = v2.Port
	template.ReadCommunity = v2.ReadCommunity
	template.WriteCommunity = v2.WriteCommunity
	template.TrapCommunity = v2.TrapCommunity
	template.Timeout = v2.Timeout
	template.PollInterval = v2.PollInterval
	template.Retries = v2.Retries
	return nil
}

// getV3TemplateParams 获取v3模板参数
func (store *SQLStore) getV3TemplateParams(ctx context.Context, q *Queries, template *SnmpTemplateUnion, id int32) error {
	v3, err := q.GetSnmpV3Parameters(ctx, id)
	if err != nil {
		return err
	}
	template.Port = v3.Port
	template.SecurityLevel = v3.SecurityLevel
	template.AuthProtocol = v3.AuthProtocol
	template.AuthPassword = v3.AuthPassword
	template.PrivProtocol = v3.PrivProtocol
	template.PrivPassword = v3.PrivPassword
	template.V3User = v3.V3User
	template.EngineID = v3.EngineID
	template.Timeout = v3.Timeout
	template.PollInterval = v3.PollInterval
	template.Retries = v3.Retries
	return nil
}

// GetSNMPTemplate 获取SNMP模板
func (store *SQLStore) GetSNMPTemplate(ctx context.Context, id int32) (SnmpTemplateUnion, error) {
	var template SnmpTemplateUnion
	err := store.execTx(ctx, func(q *Queries) error {
		base, err := q.GetSnmpTemplateBase(ctx, id)
		if err != nil {
			return err
		}
		template.TemplateID = base.ID
		template.Name = base.Name
		template.UserID = base.UserID
		template.Protocol = base.Protocol
		template.Version = base.Version
		template.Description = base.Description
		template.DeviceCount = base.DeviceCount

		switch base.Version {
		case SNMPv2c:
			return store.getV2TemplateParams(ctx, q, &template, base.ID)
		case SNMPv3:
			return store.getV3TemplateParams(ctx, q, &template, base.ID)
		default:
			return fmt.Errorf("invalid SNMP version")
		}
	})
	return template, err
}
