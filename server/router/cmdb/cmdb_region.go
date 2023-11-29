package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmdbRegionRouter struct {
}

// InitCmdbRegionRouter 初始化 区域信息 路由信息
func (s *CmdbRegionRouter) InitCmdbRegionRouter(Router *gin.RouterGroup) {
	cmdbRegionRouter := Router.Group("cmdbRegion").Use(middleware.OperationRecord())
	cmdbRegionRouterWithoutRecord := Router.Group("cmdbRegion")
	var cmdbRegionApi = v1.ApiGroupApp.CmdbApiGroup.CmdbRegionApi
	{
		cmdbRegionRouter.POST("createCmdbRegion", cmdbRegionApi.CreateCmdbRegion)   // 新建区域信息
		cmdbRegionRouter.DELETE("deleteCmdbRegion", cmdbRegionApi.DeleteCmdbRegion) // 删除区域信息
		cmdbRegionRouter.DELETE("deleteCmdbRegionByIds", cmdbRegionApi.DeleteCmdbRegionByIds) // 批量删除区域信息
		cmdbRegionRouter.PUT("updateCmdbRegion", cmdbRegionApi.UpdateCmdbRegion)    // 更新区域信息
	}
	{
		cmdbRegionRouterWithoutRecord.GET("findCmdbRegion", cmdbRegionApi.FindCmdbRegion)        // 根据ID获取区域信息
		cmdbRegionRouterWithoutRecord.GET("getCmdbRegionList", cmdbRegionApi.GetCmdbRegionList)  // 获取区域信息列表
	}
}
