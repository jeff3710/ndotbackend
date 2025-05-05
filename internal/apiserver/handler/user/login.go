package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jeff3710/ndot/internal/pkg/core"
	"github.com/jeff3710/ndot/internal/pkg/errno"
	"github.com/jeff3710/ndot/pkg/log"
)

type UserLoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

func (h *UserHandler) Login(c *gin.Context) {
	log.Infow("login function called")
	var req UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	resp, err := h.service.Login(c, req.Username, req.Password)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}
