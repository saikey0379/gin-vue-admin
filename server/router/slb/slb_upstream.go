package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SlbUpstreamRouter struct {
}

// InitSlbUpstreamRouter 初始化 目标节点 路由信息
func (s *SlbUpstreamRouter) InitSlbUpstreamRouter(Router *gin.RouterGroup) {
	slbUpstreamRouter := Router.Group("slbUpstream").Use(middleware.OperationRecord())
	slbUpstreamRouterWithoutRecord := Router.Group("slbUpstream")
	var slbUpstreamApi = v1.ApiGroupApp.SlbApiGroup.SlbUpstreamApi
	{
		slbUpstreamRouter.POST("createSlbUpstream", slbUpstreamApi.CreateSlbUpstream)   // 新建目标节点
		slbUpstreamRouter.DELETE("deleteSlbUpstream", slbUpstreamApi.DeleteSlbUpstream) // 删除目标节点
		slbUpstreamRouter.DELETE("deleteSlbUpstreamByIds", slbUpstreamApi.DeleteSlbUpstreamByIds) // 批量删除目标节点
		slbUpstreamRouter.PUT("updateSlbUpstream", slbUpstreamApi.UpdateSlbUpstream)    // 更新目标节点
	}
	{
		slbUpstreamRouterWithoutRecord.GET("findSlbUpstream", slbUpstreamApi.FindSlbUpstream)        // 根据ID获取目标节点
		slbUpstreamRouterWithoutRecord.GET("getSlbUpstreamList", slbUpstreamApi.GetSlbUpstreamList)  // 获取目标节点列表
	}
}
