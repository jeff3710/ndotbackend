// internal/repository/device_repo.go
package repository

import (
	"context"
	// "database/sql"
	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/internal/pkg/model"
)

type DeviceRepository struct {
	db db.Store
}

func NewDeviceRepository(db db.Store) *DeviceRepository {
	return &DeviceRepository{
		db: db,
	}
}

func (r *DeviceRepository) SaveDevice(ctx context.Context, dto *model.DeviceDTO) error {
	params := db.CreateDeviceParams{
		Ip:         dto.Ip,
		Hostname:   dto.Hostname,
		Model:      dto.Model,
		Vendor:     dto.Vendor,
		DeviceType: dto.DeviceType,
	}
	return r.db.CreateDevice(ctx, params)
}

func (r *DeviceRepository) GetDeviceById(ctx context.Context, id int32) (*model.DeviceDTO, error) {
	device, err := r.db.GetDeviceById(ctx, id)
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

func (r *DeviceRepository) GetDeviceByIp(ctx context.Context, ip string) (*model.DeviceDTO, error) {
	device, err := r.db.GetDeviceByIp(ctx, ip)
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
