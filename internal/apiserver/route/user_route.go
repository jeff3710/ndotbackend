package route

import (
	"github.com/gin-gonic/gin"
	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/internal/apiserver/handler/user"
	"github.com/jeff3710/ndot/internal/apiserver/service"
	"github.com/jeff3710/ndot/pkg/config"
	"github.com/jeff3710/ndot/pkg/token"
	mw "github.com/jeff3710/ndot/internal/pkg/middleware"
)

func NewUserRouter(config *config.Config, db db.Store, token token.Maker, group *gin.RouterGroup) {

	userSvc := service.NewUserService(db, token, config)
	userHandler := user.NewUserHandler(userSvc)

	group.POST("/create", userHandler.CreateUser)

	group.POST("/login", userHandler.Login)
	// group.POST("/devicesip", userHandler.)
	group.GET("/:username", mw.AuthMiddleware(token), userHandler.GetUser)
	group.GET("/id/:userid", mw.AuthMiddleware(token), userHandler.GetUserById)

	group.GET("/", mw.AuthMiddleware(token), userHandler.ListUser)

	group.DELETE("/:username", mw.AuthMiddleware(token), userHandler.DeleteUser)

	group.POST("/renew_access", mw.AuthMiddleware(token), userHandler.RenewAccessToken)
}
