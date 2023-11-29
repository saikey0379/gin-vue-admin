package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmdbIpPreemptRouter struct {
}

// InitCmdbIpPreemptRouter 初始化 预占地址 路由信息
func (s *CmdbIpPreemptRouter) InitCmdbIpPreemptRouter(Router *gin.RouterGroup) {
	cmdbIpPreemptRouter := Router.Group("cmdbIpPreempt").Use(middleware.OperationRecord())
	cmdbIpPreemptRouterWithoutRecord := Router.Group("cmdbIpPreempt")
	var cmdbIpPreemptApi = v1.ApiGroupApp.CmdbApiGroup.CmdbIpPreemptApi
	{
		cmdbIpPreemptRouter.POST("createCmdbIpPreempt", cmdbIpPreemptApi.CreateCmdbIpPreempt)   // 新建预占地址
		cmdbIpPreemptRouter.DELETE("deleteCmdbIpPreempt", cmdbIpPreemptApi.DeleteCmdbIpPreempt) // 删除预占地址
		cmdbIpPreemptRouter.DELETE("deleteCmdbIpPreemptByIds", cmdbIpPreemptApi.DeleteCmdbIpPreemptByIds) // 批量删除预占地址
		cmdbIpPreemptRouter.PUT("updateCmdbIpPreempt", cmdbIpPreemptApi.UpdateCmdbIpPreempt)    // 更新预占地址
	}
	{
		cmdbIpPreemptRouterWithoutRecord.GET("findCmdbIpPreempt", cmdbIpPreemptApi.FindCmdbIpPreempt)        // 根据ID获取预占地址
		cmdbIpPreemptRouterWithoutRecord.GET("getCmdbIpPreemptList", cmdbIpPreemptApi.GetCmdbIpPreemptList)  // 获取预占地址列表
	}
}
