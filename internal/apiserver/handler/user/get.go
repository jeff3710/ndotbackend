package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jeff3710/ndot/internal/pkg/core"
	"github.com/jeff3710/ndot/pkg/log"
)

func (h *UserHandler) GetUser(c *gin.Context) {
	log.Infow("get user function called")
	username := c.Param("username")
	resp, err := h.service.GetUser(c, username)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}

func (h *UserHandler) GetUserById(c *gin.Context) {

	userId := c.Param("userid")
	resp, err := h.service.GetUserById(c, userId)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}
