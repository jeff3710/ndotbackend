package route

import (
	"github.com/gin-gonic/gin"
	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/internal/apiserver/handler/user"
	"github.com/jeff3710/ndot/internal/apiserver/service"
	"github.com/jeff3710/ndot/pkg/config"
	"github.com/jeff3710/ndot/pkg/token"
)

func NewUserRouter(config *config.Config, db db.Store, token token.Maker, group *gin.RouterGroup) {

	userSvc := service.NewUserService(db, token, config)
	userHandler := user.NewUserHandler(userSvc)

	group.POST("/create", userHandler.CreateUser)

	group.POST("/login", userHandler.Login)
	// group.POST("/devicesip", userHandler.)
	group.GET("/:username", userHandler.GetUser)
	group.GET("/id/:userid", userHandler.GetUserById)

	group.GET("/", userHandler.ListUser)

	group.DELETE("/:username", userHandler.DeleteUser)

	group.POST("/renew_access", userHandler.RenewAccessToken)
}
