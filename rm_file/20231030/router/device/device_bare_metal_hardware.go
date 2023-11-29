package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceBareMetalHardwareRouter struct {
}

// InitDeviceBareMetalHardwareRouter 初始化 裸金属设备硬件信息 路由信息
func (s *DeviceBareMetalHardwareRouter) InitDeviceBareMetalHardwareRouter(Router *gin.RouterGroup) {
	deviceBareMetalHardwareRouter := Router.Group("deviceBareMetalHardware").Use(middleware.OperationRecord())
	deviceBareMetalHardwareRouterWithoutRecord := Router.Group("deviceBareMetalHardware")
	var deviceBareMetalHardwareApi = v1.ApiGroupApp.DeviceApiGroup.DeviceBareMetalHardwareApi
	{
		deviceBareMetalHardwareRouter.POST("createDeviceBareMetalHardware", deviceBareMetalHardwareApi.CreateDeviceBareMetalHardware)             // 新建裸金属设备硬件信息
		deviceBareMetalHardwareRouter.DELETE("deleteDeviceBareMetalHardware", deviceBareMetalHardwareApi.DeleteDeviceBareMetalHardware)           // 删除裸金属设备硬件信息
		deviceBareMetalHardwareRouter.DELETE("deleteDeviceBareMetalHardwareByIds", deviceBareMetalHardwareApi.DeleteDeviceBareMetalHardwareByIds) // 批量删除裸金属设备硬件信息
		deviceBareMetalHardwareRouter.PUT("updateDeviceBareMetalHardware", deviceBareMetalHardwareApi.UpdateDeviceBareMetalHardware)              // 更新裸金属设备硬件信息
	}
	{
		deviceBareMetalHardwareRouterWithoutRecord.GET("findDeviceBareMetalHardware", deviceBareMetalHardwareApi.FindDeviceBareMetalHardware)       // 根据ID获取裸金属设备硬件信息
		deviceBareMetalHardwareRouterWithoutRecord.GET("getDeviceBareMetalHardwareList", deviceBareMetalHardwareApi.GetDeviceBareMetalHardwareList) // 获取裸金属设备硬件信息列表
	}
}
