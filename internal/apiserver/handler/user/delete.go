package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jeff3710/ndot/internal/pkg/core"
)

func (h *UserHandler) DeleteUser(c *gin.Context) {
	username := c.Param("username")
	if err := h.service.DeleteUser(c, username); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, nil)

}
