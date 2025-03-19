package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeff3710/ndot/internal/ndot/service"
	"github.com/jeff3710/ndot/internal/pkg/model"
)


type DeviceHandler struct {
    service *service.DeviceService
}

func NewDeviceHandler(service *service.DeviceService) *DeviceHandler {
    return &DeviceHandler{service: service}
}

func (s *Server) CollectDevice(c *gin.Context) {
    var req model.SNMPRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

    info, err := s.service.CollectAndSave(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, errorResponse(err))
        return
    }

    c.JSON(http.StatusOK, info)
}
