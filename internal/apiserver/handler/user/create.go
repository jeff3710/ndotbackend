package user

import (
	"github.com/gin-gonic/gin"
	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/internal/pkg/core"
	"github.com/jeff3710/ndot/internal/pkg/errno"
	"github.com/jeff3710/ndot/pkg/log"
)

type UserCreateRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Role     string `json:"role" binding:"required"`
	Active   bool   `json:"active" binding:"required"`
}

const defaultMethods = "(GET)|(POST)|(PUT)|(DELETE)"

func (h *UserHandler) CreateUser(c *gin.Context) {
	log.Infow("create user function called")
	var req UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	// 验证参数

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
		Active:   req.Active,
	}

	if err := h.service.CreateUser(c, arg); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	// 授权
	log.Infow("create user function success", "username", req.Username)

	core.WriteResponse(c, nil, nil)
}
