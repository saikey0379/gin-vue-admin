package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OsInstallQueueRouter struct {
}

// InitOsInstallQueueRouter 初始化 操作系统安装队列 路由信息
func (s *OsInstallQueueRouter) InitOsInstallQueueRouter(Router *gin.RouterGroup) {
	osInstallQueueRouter := Router.Group("osInstallQueue").Use(middleware.OperationRecord())
	osInstallQueueRouterWithoutRecord := Router.Group("osInstallQueue")
	var osInstallQueueApi = v1.ApiGroupApp.OsInstallApiGroup.OsInstallQueueApi
	{
		osInstallQueueRouter.POST("createOsInstallQueue", osInstallQueueApi.CreateOsInstallQueue)             // 新建操作系统安装队列
		osInstallQueueRouter.DELETE("deleteOsInstallQueue", osInstallQueueApi.DeleteOsInstallQueue)           // 删除操作系统安装队列
		osInstallQueueRouter.DELETE("deleteOsInstallQueueByIds", osInstallQueueApi.DeleteOsInstallQueueByIds) // 批量删除操作系统安装队列
		osInstallQueueRouter.PUT("updateOsInstallQueue", osInstallQueueApi.UpdateOsInstallQueue)              // 更新操作系统安装队列
		osInstallQueueRouter.POST("osInstallStart", osInstallQueueApi.OsInstallStart)                         // 操作系统安装
	}
	{
		osInstallQueueRouterWithoutRecord.GET("findOsInstallQueue", osInstallQueueApi.FindOsInstallQueue)       // 根据ID获取操作系统安装队列
		osInstallQueueRouterWithoutRecord.GET("getOsInstallQueueList", osInstallQueueApi.GetOsInstallQueueList) // 获取操作系统安装队列列表
	}
}

func (s *OsInstallQueueRouter) InitOsInstallQueueRouterPublic(Router *gin.RouterGroup) {
	osInstallQueueRouterWithoutRecord := Router.Group("osInstallQueue")
	var osInstallQueueApi = v1.ApiGroupApp.OsInstallApiGroup.OsInstallQueueApi
	{
		osInstallQueueRouterWithoutRecord.GET("getDeviceInfoBySn", osInstallQueueApi.GetDeviceInfoBySn) // 根据SN获取设备信息
	}
}
