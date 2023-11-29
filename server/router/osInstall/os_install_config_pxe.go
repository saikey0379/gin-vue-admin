package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OsInstallConfigPxeRouter struct {
}

// InitOsInstallConfigPxeRouter 初始化 PXE配置 路由信息
func (s *OsInstallConfigPxeRouter) InitOsInstallConfigPxeRouter(Router *gin.RouterGroup) {
	osInstallConfigPxeRouter := Router.Group("osInstallConfigPxe").Use(middleware.OperationRecord())
	osInstallConfigPxeRouterWithoutRecord := Router.Group("osInstallConfigPxe")
	var osInstallConfigPxeApi = v1.ApiGroupApp.OsInstallApiGroup.OsInstallConfigPxeApi
	{
		osInstallConfigPxeRouter.POST("createOsInstallConfigPxe", osInstallConfigPxeApi.CreateOsInstallConfigPxe)   // 新建PXE配置
		osInstallConfigPxeRouter.DELETE("deleteOsInstallConfigPxe", osInstallConfigPxeApi.DeleteOsInstallConfigPxe) // 删除PXE配置
		osInstallConfigPxeRouter.DELETE("deleteOsInstallConfigPxeByIds", osInstallConfigPxeApi.DeleteOsInstallConfigPxeByIds) // 批量删除PXE配置
		osInstallConfigPxeRouter.PUT("updateOsInstallConfigPxe", osInstallConfigPxeApi.UpdateOsInstallConfigPxe)    // 更新PXE配置
	}
	{
		osInstallConfigPxeRouterWithoutRecord.GET("findOsInstallConfigPxe", osInstallConfigPxeApi.FindOsInstallConfigPxe)        // 根据ID获取PXE配置
		osInstallConfigPxeRouterWithoutRecord.GET("getOsInstallConfigPxeList", osInstallConfigPxeApi.GetOsInstallConfigPxeList)  // 获取PXE配置列表
	}
}
