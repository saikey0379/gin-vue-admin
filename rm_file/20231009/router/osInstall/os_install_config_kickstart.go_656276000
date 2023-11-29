package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OsInstallConfigKickstartRouter struct {
}

// InitOsInstallConfigKickstartRouter 初始化 Kickstart配置 路由信息
func (s *OsInstallConfigKickstartRouter) InitOsInstallConfigKickstartRouter(Router *gin.RouterGroup) {
	ksRouter := Router.Group("ks").Use(middleware.OperationRecord())
	ksRouterWithoutRecord := Router.Group("ks")
	var ksApi = v1.ApiGroupApp.OsInstallApiGroup.OsInstallConfigKickstartApi
	{
		ksRouter.POST("createOsInstallConfigKickstart", ksApi.CreateOsInstallConfigKickstart)   // 新建Kickstart配置
		ksRouter.DELETE("deleteOsInstallConfigKickstart", ksApi.DeleteOsInstallConfigKickstart) // 删除Kickstart配置
		ksRouter.DELETE("deleteOsInstallConfigKickstartByIds", ksApi.DeleteOsInstallConfigKickstartByIds) // 批量删除Kickstart配置
		ksRouter.PUT("updateOsInstallConfigKickstart", ksApi.UpdateOsInstallConfigKickstart)    // 更新Kickstart配置
	}
	{
		ksRouterWithoutRecord.GET("findOsInstallConfigKickstart", ksApi.FindOsInstallConfigKickstart)        // 根据ID获取Kickstart配置
		ksRouterWithoutRecord.GET("getOsInstallConfigKickstartList", ksApi.GetOsInstallConfigKickstartList)  // 获取Kickstart配置列表
	}
}
