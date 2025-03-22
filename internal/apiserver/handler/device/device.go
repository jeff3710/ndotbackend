package device

import (
	"github.com/jeff3710/ndot/internal/apiserver/service"
)

type DeviceHandler struct {
	service service.DeviceServiceInterface
}

func NewDeviceHandler(service service.DeviceServiceInterface) *DeviceHandler {
	return &DeviceHandler{service: service}
}
