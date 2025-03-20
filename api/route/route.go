package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeff3710/ndot/internal/pkg/core"
	"github.com/jeff3710/ndot/internal/pkg/errno"
	"github.com/jeff3710/ndot/server"
)

func Setup(config *server.Config, pool *pgxpool.Pool,g *gin.Engine){
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c,errno.ErrPageNotFound,nil)
	})
	g.GET("/ping", func(c *gin.Context) {
		core.WriteResponse(c,nil,map[string]string{"status":"ok"})
	})
	publicRouter:=g.Group("/v1")
	NewDeviceRouter(config,pool,publicRouter)
}