package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SlbAccesslistRouter struct {
}

// InitSlbAccesslistRouter 初始化 访问控制 路由信息
func (s *SlbAccesslistRouter) InitSlbAccesslistRouter(Router *gin.RouterGroup) {
	slbAccesslistRouter := Router.Group("slbAccesslist").Use(middleware.OperationRecord())
	slbAccesslistRouterWithoutRecord := Router.Group("slbAccesslist")
	var slbAccesslistApi = v1.ApiGroupApp.SlbApiGroup.SlbAccesslistApi
	{
		slbAccesslistRouter.POST("createSlbAccesslist", slbAccesslistApi.CreateSlbAccesslist)   // 新建访问控制
		slbAccesslistRouter.DELETE("deleteSlbAccesslist", slbAccesslistApi.DeleteSlbAccesslist) // 删除访问控制
		slbAccesslistRouter.DELETE("deleteSlbAccesslistByIds", slbAccesslistApi.DeleteSlbAccesslistByIds) // 批量删除访问控制
		slbAccesslistRouter.PUT("updateSlbAccesslist", slbAccesslistApi.UpdateSlbAccesslist)    // 更新访问控制
	}
	{
		slbAccesslistRouterWithoutRecord.GET("findSlbAccesslist", slbAccesslistApi.FindSlbAccesslist)        // 根据ID获取访问控制
		slbAccesslistRouterWithoutRecord.GET("getSlbAccesslistList", slbAccesslistApi.GetSlbAccesslistList)  // 获取访问控制列表
	}
}
