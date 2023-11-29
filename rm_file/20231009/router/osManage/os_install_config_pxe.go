package osManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PxeRouter struct {
}

// InitPxeRouter 初始化 PXE配置 路由信息
func (s *PxeRouter) InitPxeRouter(Router *gin.RouterGroup) {
	PxeConfigRouter := Router.Group("PxeConfig").Use(middleware.OperationRecord())
	PxeConfigRouterWithoutRecord := Router.Group("PxeConfig")
	var PxeConfigApi = v1.ApiGroupApp.OsManageApiGroup.PxeApi
	{
		PxeConfigRouter.POST("createPxe", PxeConfigApi.CreatePxe)   // 新建PXE配置
		PxeConfigRouter.DELETE("deletePxe", PxeConfigApi.DeletePxe) // 删除PXE配置
		PxeConfigRouter.DELETE("deletePxeByIds", PxeConfigApi.DeletePxeByIds) // 批量删除PXE配置
		PxeConfigRouter.PUT("updatePxe", PxeConfigApi.UpdatePxe)    // 更新PXE配置
	}
	{
		PxeConfigRouterWithoutRecord.GET("findPxe", PxeConfigApi.FindPxe)        // 根据ID获取PXE配置
		PxeConfigRouterWithoutRecord.GET("getPxeList", PxeConfigApi.GetPxeList)  // 获取PXE配置列表
	}
}
