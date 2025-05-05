package service

import (
	"github.com/gin-gonic/gin"
	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/pkg/token"
)

type UserServiceInterface interface {
	CreateUser(ctx *gin.Context, req db.CreateUserParams) error
	Login(ctx *gin.Context, username, password string) (*LoginUserResponse, error)
	GetUser(ctx *gin.Context, username string) (*GetUserResponse, error)
	GetUserById(ctx *gin.Context, userId string) (*GetUserResponse, error)
	ListUser(ctx *gin.Context) (*ListUserResponse, error)
	DeleteUser(ctx *gin.Context, username string) error
	RenewAccessToken(ctx *gin.Context, refreshToken string) (string, *token.Payload, error)
}
