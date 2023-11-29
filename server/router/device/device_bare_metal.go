package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceBareMetalRouter struct {
}

// InitDeviceBareMetalRouter 初始化 裸金属设备 路由信息
func (s *DeviceBareMetalRouter) InitDeviceBareMetalRouter(Router *gin.RouterGroup) {
	deviceBareMetalRouter := Router.Group("deviceBareMetal").Use(middleware.OperationRecord())
	deviceBareMetalRouterWithoutRecord := Router.Group("deviceBareMetal")
	var deviceBareMetalApi = v1.ApiGroupApp.DeviceApiGroup.DeviceBareMetalApi
	{
		deviceBareMetalRouter.POST("createDeviceBareMetal", deviceBareMetalApi.CreateDeviceBareMetal)             // 新建裸金属设备
		deviceBareMetalRouter.DELETE("deleteDeviceBareMetal", deviceBareMetalApi.DeleteDeviceBareMetal)           // 删除裸金属设备
		deviceBareMetalRouter.DELETE("deleteDeviceBareMetalByIds", deviceBareMetalApi.DeleteDeviceBareMetalByIds) // 批量删除裸金属设备
		deviceBareMetalRouter.PUT("updateDeviceBareMetal", deviceBareMetalApi.UpdateDeviceBareMetal)              // 更新裸金属设备
	}
	{
		deviceBareMetalRouterWithoutRecord.GET("findDeviceBareMetal", deviceBareMetalApi.FindDeviceBareMetal)       // 根据ID获取裸金属设备
		deviceBareMetalRouterWithoutRecord.GET("getDeviceBareMetalList", deviceBareMetalApi.GetDeviceBareMetalList) // 获取裸金属设备列表

	}
}
