package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OsInstallLogRouter struct {
}

// InitOsInstallLogRouter 初始化 操作系统安装日志 路由信息
func (s *OsInstallLogRouter) InitOsInstallLogRouter(Router *gin.RouterGroup) {
	osInstallLogRouter := Router.Group("osInstallLog").Use(middleware.OperationRecord())
	osInstallLogRouterWithoutRecord := Router.Group("osInstallLog")
	var osInstallLogApi = v1.ApiGroupApp.OsInstallApiGroup.OsInstallLogApi
	{
		osInstallLogRouter.POST("createOsInstallLog", osInstallLogApi.CreateOsInstallLog)             // 新建操作系统安装日志
		osInstallLogRouter.DELETE("deleteOsInstallLog", osInstallLogApi.DeleteOsInstallLog)           // 删除操作系统安装日志
		osInstallLogRouter.DELETE("deleteOsInstallLogByIds", osInstallLogApi.DeleteOsInstallLogByIds) // 批量删除操作系统安装日志
		osInstallLogRouter.PUT("updateOsInstallLog", osInstallLogApi.UpdateOsInstallLog)              // 更新操作系统安装日志
	}
	{
		osInstallLogRouterWithoutRecord.GET("findOsInstallLog", osInstallLogApi.FindOsInstallLog)       // 根据ID获取操作系统安装日志
		osInstallLogRouterWithoutRecord.GET("getOsInstallLogList", osInstallLogApi.GetOsInstallLogList) // 获取操作系统安装日志列表
	}
}

func (s *OsInstallLogRouter) InitOsInstallLogRouterPublic(Router *gin.RouterGroup) {
	osInstallLogRouterWithoutRecord := Router.Group("osInstallLog")
	var osInstallLogApi = v1.ApiGroupApp.OsInstallApiGroup.OsInstallLogApi
	{
		osInstallLogRouterWithoutRecord.GET("report", osInstallLogApi.Report) // 根据SN获取设备信息
	}
}
