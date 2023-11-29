package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ConfigKickstartRouter struct {
}

// InitConfigKickstartRouter 初始化 Kickstart配置 路由信息
func (s *ConfigKickstartRouter) InitConfigKickstartRouter(Router *gin.RouterGroup) {
	ksRouter := Router.Group("ks").Use(middleware.OperationRecord())
	ksRouterWithoutRecord := Router.Group("ks")
	var ksApi = v1.ApiGroupApp.OsInstallApiGroup.ConfigKickstartApi
	{
		ksRouter.POST("createConfigKickstart", ksApi.CreateConfigKickstart)   // 新建Kickstart配置
		ksRouter.DELETE("deleteConfigKickstart", ksApi.DeleteConfigKickstart) // 删除Kickstart配置
		ksRouter.DELETE("deleteConfigKickstartByIds", ksApi.DeleteConfigKickstartByIds) // 批量删除Kickstart配置
		ksRouter.PUT("updateConfigKickstart", ksApi.UpdateConfigKickstart)    // 更新Kickstart配置
	}
	{
		ksRouterWithoutRecord.GET("findConfigKickstart", ksApi.FindConfigKickstart)        // 根据ID获取Kickstart配置
		ksRouterWithoutRecord.GET("getConfigKickstartList", ksApi.GetConfigKickstartList)  // 获取Kickstart配置列表
	}
}
