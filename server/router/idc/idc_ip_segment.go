package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdcIpSegmentRouter struct {
}

// InitIdcIpSegmentRouter 初始化 数据中心网段 路由信息
func (s *IdcIpSegmentRouter) InitIdcIpSegmentRouter(Router *gin.RouterGroup) {
	idcIpSegmentRouter := Router.Group("idcIpSegment").Use(middleware.OperationRecord())
	idcIpSegmentRouterWithoutRecord := Router.Group("idcIpSegment")
	var idcIpSegmentApi = v1.ApiGroupApp.IdcApiGroup.IdcIpSegmentApi
	{
		idcIpSegmentRouter.POST("createIdcIpSegment", idcIpSegmentApi.CreateIdcIpSegment)   // 新建数据中心网段
		idcIpSegmentRouter.DELETE("deleteIdcIpSegment", idcIpSegmentApi.DeleteIdcIpSegment) // 删除数据中心网段
		idcIpSegmentRouter.DELETE("deleteIdcIpSegmentByIds", idcIpSegmentApi.DeleteIdcIpSegmentByIds) // 批量删除数据中心网段
		idcIpSegmentRouter.PUT("updateIdcIpSegment", idcIpSegmentApi.UpdateIdcIpSegment)    // 更新数据中心网段
	}
	{
		idcIpSegmentRouterWithoutRecord.GET("findIdcIpSegment", idcIpSegmentApi.FindIdcIpSegment)        // 根据ID获取数据中心网段
		idcIpSegmentRouterWithoutRecord.GET("getIdcIpSegmentList", idcIpSegmentApi.GetIdcIpSegmentList)  // 获取数据中心网段列表
	}
}
