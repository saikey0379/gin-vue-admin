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
	osInstallConfigKickstartRouter := Router.Group("osInstallConfigKickstart").Use(middleware.OperationRecord())
	osInstallConfigKickstartRouterWithoutRecord := Router.Group("osInstallConfigKickstart")
	var osInstallConfigKickstartApi = v1.ApiGroupApp.OsInstallApiGroup.OsInstallConfigKickstartApi
	{
		osInstallConfigKickstartRouter.POST("createOsInstallConfigKickstart", osInstallConfigKickstartApi.CreateOsInstallConfigKickstart)             // 新建Kickstart配置
		osInstallConfigKickstartRouter.DELETE("deleteOsInstallConfigKickstart", osInstallConfigKickstartApi.DeleteOsInstallConfigKickstart)           // 删除Kickstart配置
		osInstallConfigKickstartRouter.DELETE("deleteOsInstallConfigKickstartByIds", osInstallConfigKickstartApi.DeleteOsInstallConfigKickstartByIds) // 批量删除Kickstart配置
		osInstallConfigKickstartRouter.PUT("updateOsInstallConfigKickstart", osInstallConfigKickstartApi.UpdateOsInstallConfigKickstart)              // 更新Kickstart配置
	}
	{
		osInstallConfigKickstartRouterWithoutRecord.GET("findOsInstallConfigKickstart", osInstallConfigKickstartApi.FindOsInstallConfigKickstart)       // 根据ID获取Kickstart配置
		osInstallConfigKickstartRouterWithoutRecord.GET("getOsInstallConfigKickstartList", osInstallConfigKickstartApi.GetOsInstallConfigKickstartList) // 获取Kickstart配置列表
	}
}

func (s *OsInstallConfigKickstartRouter) InitOsInstallConfigKickstartRouterPublic(Router *gin.RouterGroup) {
	osInstallConfigKickstartRouterWithoutRecord := Router.Group("osInstallConfigKickstart")
	var osInstallConfigKickstartApi = v1.ApiGroupApp.OsInstallApiGroup.OsInstallConfigKickstartApi
	{
		osInstallConfigKickstartRouterWithoutRecord.GET("getContentBySn", osInstallConfigKickstartApi.GetContentBySn) // 根据Sn获取Content
	}
}
