package route

import (
	"github.com/gin-gonic/gin"
	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/internal/apiserver/handler/device"
	"github.com/jeff3710/ndot/internal/apiserver/service"
	mw "github.com/jeff3710/ndot/internal/pkg/middleware"
	"github.com/jeff3710/ndot/pkg/config"
	"github.com/jeff3710/ndot/pkg/snmp"
	"github.com/jeff3710/ndot/pkg/token"
)

func NewDeviceRouter(config *config.Config, db db.Store,token token.Maker, group *gin.RouterGroup) {


	snmp := snmp.NewSNMPClient(config)
	
	deviceSvc := service.NewDeviceService(db, snmp)
	deviceHandler := device.NewDeviceHandler(deviceSvc)

	group.POST("/devicesadd", deviceHandler.CollectDevice)
	group.POST("/devicesid", deviceHandler.GetDeviceById)
	group.POST("/devicesip", deviceHandler.GetDeviceByIp)
	group.POST("/snmptemplate", mw.AuthMiddleware(token),deviceHandler.AddSnmpTemplate)
}
