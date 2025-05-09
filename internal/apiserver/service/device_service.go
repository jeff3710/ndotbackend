// internal/service/device_service.go
package service

import (
	"context"

	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/internal/pkg/model"
	v1 "github.com/jeff3710/ndot/pkg/api/ndot/v1"
	"github.com/jeff3710/ndot/pkg/snmp"
)

type DeviceService struct {
	db         db.Store
	snmpClient snmp.SNMPClientInterface
}

func NewDeviceService(db db.Store, client snmp.SNMPClientInterface) DeviceServiceInterface {
	return &DeviceService{db: db, snmpClient: client}
}

func (s *DeviceService) CollectAndSave(ctx context.Context, req *model.SNMPRequest) (*model.DeviceDTO, error) {
	// 获取设备信息
	info, err := s.snmpClient.GetDeviceInfo(req)
	if err != nil {
		return nil, err
	}

	arg := db.CreateDeviceParams{
		Ip:         info.IP,
		Hostname:   info.Hostname,
		Model:      info.Model,
		Vendor:     info.Vendor,
		DeviceType: info.DeviceType,
	}

	// 转换为DTO
	dto := &model.DeviceDTO{
		Ip:         info.IP,
		Hostname:   info.Hostname,
		Model:      info.Model,
		Vendor:     info.Vendor,
		DeviceType: info.DeviceType,
	}

	// 存储到数据库
	if err := s.db.CreateDevice(ctx, arg); err != nil {
		return nil, err
	}

	return dto, nil
}

func (s *DeviceService) GetDeviceById(ctx context.Context, id int32) (*model.DeviceDTO, error) {
	device, err := s.db.GetDeviceById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.DeviceDTO{
		Ip:         device.Ip,
		Hostname:   device.Hostname,
		Model:      device.Model,
		Vendor:     device.Vendor,
		DeviceType: device.DeviceType,
	}, nil
}

func (s *DeviceService) GetDeviceByIp(ctx context.Context, ip string) (*model.DeviceDTO, error) {

	device, err := s.db.GetDeviceByIp(ctx, ip)
	if err != nil {
		return nil, err
	}
	return &model.DeviceDTO{
		Ip:         device.Ip,
		Hostname:   device.Hostname,
		Model:      device.Model,
		Vendor:     device.Vendor,
		DeviceType: device.DeviceType,
	}, nil
}

func (s *DeviceService) CreateSnmpTemplate(ctx context.Context, req v1.SnmpTemplateRequest, userId int32) error {

	r := db.SnmpTemplateUnion{
		UserID:         userId,
		Name:           req.Name,
		Protocol:       req.Protocol,
		Version:        req.Version,
		Description:    req.Description,
		Port:           req.Port,
		ReadCommunity:  req.ReadCommunity,
		WriteCommunity: req.WriteCommunity,
		TrapCommunity:  req.TrapCommunity,
		Timeout:        req.Timeout,
		PollInterval:   req.PollInterval,
		Retries:        req.Retries,
		SecurityLevel:  req.SecurityLevel,
		AuthProtocol:   req.AuthProtocol,
		AuthPassword:   req.AuthPassword,
		PrivProtocol:   req.PrivProtocol,
		PrivPassword:   req.PrivPassword,
		V3User:         req.V3User,
		EngineID:       req.EngineID,
	}

	if err := s.db.CreateSNMPTempate(ctx, r); err != nil {
		return err
	}
	return nil
}
