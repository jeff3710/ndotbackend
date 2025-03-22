package repository

import (
	"context"

	"github.com/jeff3710/ndot/internal/pkg/model"
)

type DeviceRepositoryInterface interface {
	SaveDevice(ctx context.Context, dto *model.DeviceDTO) error
	GetDeviceById(ctx context.Context, id int32) (*model.DeviceDTO, error)
	GetDeviceByIp(ctx context.Context, ip string) (*model.DeviceDTO, error)
	GetAllDevices(ctx context.Context) ([]*model.DeviceDTO, error)
	UpdateDeviceAll(ctx context.Context, dto *model.DeviceDTO) error
}
