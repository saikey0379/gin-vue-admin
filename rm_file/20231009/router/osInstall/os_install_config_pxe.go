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
	pxeRouter := Router.Group("pxe").Use(middleware.OperationRecord())
	pxeRouterWithoutRecord := Router.Group("pxe")
	var pxeApi = v1.ApiGroupApp.OsInstallApiGroup.OsInstallConfigPxeApi
	{
		pxeRouter.POST("createOsInstallConfigPxe", pxeApi.CreateOsInstallConfigPxe)   // 新建PXE配置
		pxeRouter.DELETE("deleteOsInstallConfigPxe", pxeApi.DeleteOsInstallConfigPxe) // 删除PXE配置
		pxeRouter.DELETE("deleteOsInstallConfigPxeByIds", pxeApi.DeleteOsInstallConfigPxeByIds) // 批量删除PXE配置
		pxeRouter.PUT("updateOsInstallConfigPxe", pxeApi.UpdateOsInstallConfigPxe)    // 更新PXE配置
	}
	{
		pxeRouterWithoutRecord.GET("findOsInstallConfigPxe", pxeApi.FindOsInstallConfigPxe)        // 根据ID获取PXE配置
		pxeRouterWithoutRecord.GET("getOsInstallConfigPxeList", pxeApi.GetOsInstallConfigPxeList)  // 获取PXE配置列表
	}
}
