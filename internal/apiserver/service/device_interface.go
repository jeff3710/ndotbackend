package service

import (
	"context"

	"github.com/jeff3710/ndot/internal/pkg/model"
	v1 "github.com/jeff3710/ndot/pkg/api/ndot/v1"
)

type DeviceServiceInterface interface {
	CollectAndSave(ctx context.Context, req *model.SNMPRequest) (*model.DeviceDTO, error)
	GetDeviceById(ctx context.Context, id int32) (*model.DeviceDTO, error)
	GetDeviceByIp(ctx context.Context, ip string) (*model.DeviceDTO, error)
	CreateSnmpTemplate(ctx context.Context,req v1.SnmpTemplateRequest) error
}
