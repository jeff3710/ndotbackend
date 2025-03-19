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

func (r *DeviceRepository) SaveDevice(ctx context.Context, info *model.DeviceInfo) error {
    params := db.CreateDeviceParams{
        Ip:         info.IP,
        Hostname:   info.Hostname,
        Model:      info.Model,
        Vendor:     info.Vendor,
        DeviceType: info.DeviceType,
    }
    return r.db.CreateDevice(ctx, params)
}

func (r *DeviceRepository) GetDeviceById(ctx context.Context, id int32) (*model.DeviceInfo, error) {
    device, err := r.db.GetDeviceById(ctx, id)
    if err != nil {
        return nil, err
    }
    info := &model.DeviceInfo{
        IP:         device.Ip,
        Hostname:   device.Hostname,
        Model:      device.Model,
        Vendor:     device.Vendor,
        DeviceType: device.DeviceType,
    }
    return info, nil
}

func (r *DeviceRepository) GetDeviceByIp(ctx context.Context, ip string) (*model.DeviceInfo, error) {
    device, err := r.db.GetDeviceByIp(ctx, ip)
    if err!= nil {
        return nil, err
    }
    info := &model.DeviceInfo{
        IP:         device.Ip,
        Hostname:   device.Hostname,
        Model:      device.Model,
        Vendor:     device.Vendor,
        DeviceType: device.DeviceType,
    }
    return info, nil
}
