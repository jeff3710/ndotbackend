package route

import (
	"github.com/gin-gonic/gin"
	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/internal/pkg/core"
	"github.com/jeff3710/ndot/internal/pkg/errno"
	"github.com/jeff3710/ndot/pkg/config"
	"github.com/jeff3710/ndot/pkg/token"
)

func Setup(config *config.Config, db db.Store, token token.Maker, g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})
	g.GET("/ping", func(c *gin.Context) {
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})
	publicRouter := g.Group("/v1/device")
	NewDeviceRouter(config, db, publicRouter)
	userRouter := g.Group("/v1/user")
	NewUserRouter(config, db, token, userRouter)
}
