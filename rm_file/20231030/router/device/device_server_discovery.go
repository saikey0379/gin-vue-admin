package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceServerDiscoveryRouter struct {
}

// InitDeviceServerDiscoveryRouter 初始化 Server发现 路由信息
func (s *DeviceServerDiscoveryRouter) InitDeviceServerDiscoveryRouter(Router *gin.RouterGroup) {
	deviceServerDiscoveryRouter := Router.Group("deviceServerDiscovery").Use(middleware.OperationRecord())
	deviceServerDiscoveryRouterWithoutRecord := Router.Group("deviceServerDiscovery")
	var deviceServerDiscoveryApi = v1.ApiGroupApp.DeviceApiGroup.DeviceServerDiscoveryApi
	{
		deviceServerDiscoveryRouter.POST("createDeviceServerDiscovery", deviceServerDiscoveryApi.CreateDeviceServerDiscovery)   // 新建Server发现
		deviceServerDiscoveryRouter.DELETE("deleteDeviceServerDiscovery", deviceServerDiscoveryApi.DeleteDeviceServerDiscovery) // 删除Server发现
		deviceServerDiscoveryRouter.DELETE("deleteDeviceServerDiscoveryByIds", deviceServerDiscoveryApi.DeleteDeviceServerDiscoveryByIds) // 批量删除Server发现
		deviceServerDiscoveryRouter.PUT("updateDeviceServerDiscovery", deviceServerDiscoveryApi.UpdateDeviceServerDiscovery)    // 更新Server发现
	}
	{
		deviceServerDiscoveryRouterWithoutRecord.GET("findDeviceServerDiscovery", deviceServerDiscoveryApi.FindDeviceServerDiscovery)        // 根据ID获取Server发现
		deviceServerDiscoveryRouterWithoutRecord.GET("getDeviceServerDiscoveryList", deviceServerDiscoveryApi.GetDeviceServerDiscoveryList)  // 获取Server发现列表
	}
}
