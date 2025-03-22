package service

import (
	"context"

	"github.com/jeff3710/ndot/internal/pkg/model"
)

type DeviceServiceInterface interface {
	CollectAndSave(ctx context.Context, req *model.SNMPRequest) (*model.DeviceDTO, error)
	GetDeviceById(ctx context.Context, id int32) (*model.DeviceDTO, error)
	GetDeviceByIp(ctx context.Context, ip string) (*model.DeviceDTO, error)
}
