package api

import (
	"github.com/jeff3710/ndot/internal/ndot/service"
)

type DeviceHandler struct {
	service service.DeviceServiceInterface
}

func NewDeviceHandler(service service.DeviceServiceInterface) *DeviceHandler {
	return &DeviceHandler{service: service}
}

// func (s *Server) CollectDevice(c *gin.Context) {
// 	var req model.SNMPRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	info, err := s.service.CollectAndSave(c.Request.Context(), &req)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, info)
// }
