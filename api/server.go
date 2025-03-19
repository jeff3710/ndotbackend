package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jeff3710/ndot/config"
	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/internal/ndot/repository"
	"github.com/jeff3710/ndot/internal/ndot/service"
	"github.com/jeff3710/ndot/pkg/snmp"
)

type Server struct {
	config  config.Config
	service *service.DeviceService
	router  *gin.Engine
	snmp    *snmp.SNMPClient
}

func NewServer(config config.Config, store db.Store, snmp *snmp.SNMPClient) *Server {
	deviceService := service.NewDeviceService(repository.NewDeviceRepository(store), snmp)
	server := &Server{
		config:  config,
		service: deviceService,
	}
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/v1")
	{

		v1.POST("/devices", s.CollectDevice)
	}

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
