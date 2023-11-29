package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmdbIpSegmentRouter struct {
}

// InitCmdbIpSegmentRouter 初始化 网段管理 路由信息
func (s *CmdbIpSegmentRouter) InitCmdbIpSegmentRouter(Router *gin.RouterGroup) {
	cmdbIpSegmentRouter := Router.Group("cmdbIpSegment").Use(middleware.OperationRecord())
	cmdbIpSegmentRouterWithoutRecord := Router.Group("cmdbIpSegment")
	var cmdbIpSegmentApi = v1.ApiGroupApp.CmdbApiGroup.CmdbIpSegmentApi
	{
		cmdbIpSegmentRouter.POST("createCmdbIpSegment", cmdbIpSegmentApi.CreateCmdbIpSegment)   // 新建网段管理
		cmdbIpSegmentRouter.DELETE("deleteCmdbIpSegment", cmdbIpSegmentApi.DeleteCmdbIpSegment) // 删除网段管理
		cmdbIpSegmentRouter.DELETE("deleteCmdbIpSegmentByIds", cmdbIpSegmentApi.DeleteCmdbIpSegmentByIds) // 批量删除网段管理
		cmdbIpSegmentRouter.PUT("updateCmdbIpSegment", cmdbIpSegmentApi.UpdateCmdbIpSegment)    // 更新网段管理
	}
	{
		cmdbIpSegmentRouterWithoutRecord.GET("findCmdbIpSegment", cmdbIpSegmentApi.FindCmdbIpSegment)        // 根据ID获取网段管理
		cmdbIpSegmentRouterWithoutRecord.GET("getCmdbIpSegmentList", cmdbIpSegmentApi.GetCmdbIpSegmentList)  // 获取网段管理列表
	}
}
