package device

import (
	"github.com/gin-gonic/gin"
	"github.com/jeff3710/ndot/internal/pkg/core"
	"github.com/jeff3710/ndot/internal/pkg/errno"
	"github.com/jeff3710/ndot/internal/pkg/model"
)

func (dh *DeviceHandler) GetDeviceById(c *gin.Context) {
	id := int32(10)
	device, err := dh.service.GetDeviceById(c.Request.Context(), id)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, device)
}

func (dh *DeviceHandler) GetDeviceByIp(c *gin.Context) {
	var req model.DeviceIp
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	device, err := dh.service.GetDeviceByIp(c.Request.Context(), req.IP)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, device)
}
