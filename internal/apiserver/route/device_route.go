package route

import (
	"github.com/gin-gonic/gin"
	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/internal/apiserver/handler/device"
	"github.com/jeff3710/ndot/internal/apiserver/service"
	"github.com/jeff3710/ndot/pkg/config"
	"github.com/jeff3710/ndot/pkg/snmp"
)

func NewDeviceRouter(config *config.Config, db db.Store, group *gin.RouterGroup) {


	snmp := snmp.NewSNMPClient(config)
	
	deviceSvc := service.NewDeviceService(db, snmp)
	deviceHandler := device.NewDeviceHandler(deviceSvc)

	group.POST("/devicesadd", deviceHandler.CollectDevice)
	group.POST("/devicesid", deviceHandler.GetDeviceById)
	group.POST("/devicesip", deviceHandler.GetDeviceByIp)
	group.POST("/snmptemplate", deviceHandler.AddSnmpTemplate)
}
