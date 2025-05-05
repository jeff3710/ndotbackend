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
	// 当db.Store是接口类型时（常见于sqlc生成的代码）
	// 保持当前的值传递方式是正确用法
	// 因为接口本身已经隐含指针语义

	// 当db.Store是结构体类型时（较少见）
	// 应该改为指针传递：db *db.Store
	// 以避免结构体复制的性能开销

	snmp := snmp.NewSNMPClient(config)
	
	deviceSvc := service.NewDeviceService(db, snmp)
	deviceHandler := device.NewDeviceHandler(deviceSvc)

	group.POST("/devicesadd", deviceHandler.CollectDevice)
	group.POST("/devicesid", deviceHandler.GetDeviceById)
	group.POST("/devicesip", deviceHandler.GetDeviceByIp)
}
