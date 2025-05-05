package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jeff3710/ndot/internal/pkg/core"
	"github.com/jeff3710/ndot/pkg/log"
)

func (h *UserHandler) ListUser(c *gin.Context) {
	log.Infow("list user function called")
	resp, err := h.service.ListUser(c)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}
