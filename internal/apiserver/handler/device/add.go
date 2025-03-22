package device

import (
	"github.com/gin-gonic/gin"
	"github.com/jeff3710/ndot/internal/pkg/core"
	"github.com/jeff3710/ndot/internal/pkg/errno"
	"github.com/jeff3710/ndot/internal/pkg/model"
)

func (dh *DeviceHandler) CollectDevice(c *gin.Context) {
	var req model.SNMPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	info, err := dh.service.CollectAndSave(c.Request.Context(), &req)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, info)
}
