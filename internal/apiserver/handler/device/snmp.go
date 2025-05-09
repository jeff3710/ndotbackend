package device

import (
	"github.com/gin-gonic/gin"
	"github.com/jeff3710/ndot/internal/pkg/core"
	"github.com/jeff3710/ndot/internal/pkg/errno"
	v1 "github.com/jeff3710/ndot/pkg/api/ndot/v1"
	"github.com/jeff3710/ndot/pkg/log"
)

func (h *DeviceHandler) AddSnmpTemplate(c *gin.Context) {
	var req v1.SnmpTemplateRequest
	value, exists := c.Get("user_id")
	if !exists {
		log.Errorw("user ID not found in context")
		core.WriteResponse(c, errno.ErrUserIdNotFound, nil)
		return
	}
	userId, ok := value.(int32)
	if !ok {
		log.Errorw("user ID type assertion failed", "value", value)
		core.WriteResponse(c, errno.ErrUserTypeAssertionFailed, nil)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	if err := h.service.CreateSnmpTemplate(c, req, userId); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, nil)
}
