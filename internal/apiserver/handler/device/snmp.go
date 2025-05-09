package device

import (
	"github.com/gin-gonic/gin"
	"github.com/jeff3710/ndot/internal/pkg/core"
	"github.com/jeff3710/ndot/internal/pkg/errno"
	v1 "github.com/jeff3710/ndot/pkg/api/ndot/v1"
)




func (h *DeviceHandler) AddSnmpTemplate(c *gin.Context){
	var req v1.SnmpTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c,errno.ErrBind,nil)
		return
	}
	if err:=h.service.CreateSnmpTemplate(c, req);err!=nil {
		core.WriteResponse(c,err,nil)
		return
	}
	core.WriteResponse(c,nil,nil)
}
