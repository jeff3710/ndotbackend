package api

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/jeff3710/ndot/internal/pkg/model"
)

// type Server struct {
// 	config  server.Config
// 	service service.DeviceServiceInterface
// 	router  *gin.Engine
// 	snmp    snmp.SNMPClientInterface
// }

// func NewServer(config server.Config, deviceService service.DeviceServiceInterface, snmp snmp.SNMPClientInterface) *Server {
// 	server := &Server{
// 		config:  config,
// 		service: deviceService,
// 	}
// 	server.setupRoutes()
// 	return server
// }

// func (s *Server) setupRoutes() {
// 	router := gin.Default()

// 	router.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	v1 := router.Group("/v1")
// 	{

// 		v1.POST("/devices", s.CollectDevice)
// 	}

// 	s.router = router
// }

// func (s *Server) Start(address string) error {
// 	return s.router.Run(address)
// }

func errorResponse(err error) *model.ErrorResponse {
	// 根据错误类型细化状态码
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return &model.ErrorResponse{HTTPStatus: 504, Code: "GATEWAY_TIMEOUT", Message: "请求超时"}
	case errors.As(err, &validator.ValidationErrors{}):
		return &model.ErrorResponse{HTTPStatus: 422, Code: "VALIDATION_ERROR", Message: "参数校验失败", Details: err.Error()}
	default:
		return &model.ErrorResponse{HTTPStatus: 500, Code: "INTERNAL_ERROR", Message: "服务器内部错误"}
	}
}
