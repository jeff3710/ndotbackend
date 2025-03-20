package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeff3710/ndot/server"
)


func NewDeviceRouter(config *server.Config, pool *pgxpool.Pool, group *gin.RouterGroup)  {
	

	group.POST("/devices", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	} )
}