// internal/service/device_service.go
package service

import (
	"context"

	"github.com/jeff3710/ndot/internal/pkg/model"
	"github.com/jeff3710/ndot/pkg/snmp"

	"github.com/jeff3710/ndot/internal/ndot/repository"
)

type DeviceService struct {
    repo      *repository.DeviceRepository
    snmpClient *snmp.SNMPClient
}

func NewDeviceService(repo *repository.DeviceRepository, client *snmp.SNMPClient) *DeviceService {
    return &DeviceService{repo: repo, snmpClient: client}
}

func (s *DeviceService) CollectAndSave(ctx context.Context, req *model.SNMPRequest) (*model.DeviceInfo, error) {
    // 获取设备信息
    info, err := s.snmpClient.GetDeviceInfo(req)
    if err != nil {
        return nil, err
    }

    // 存储到数据库
    if err := s.repo.SaveDevice(ctx, info); err != nil {
        return nil, err
    }

    return info, nil
}

func (s *DeviceService) GetDeviceById(ctx context.Context, id int32) (*model.DeviceInfo, error) {
    return s.repo.GetDeviceById(ctx, id)
}

func (s *DeviceService) GetDeviceByIp(ctx context.Context, ip string) (*model.DeviceInfo, error) {
    return s.repo.GetDeviceByIp(ctx, ip)
}
