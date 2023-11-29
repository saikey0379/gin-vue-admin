package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdcIpPreemptRouter struct {
}

// InitIdcIpPreemptRouter 初始化 数据中心地址预占 路由信息
func (s *IdcIpPreemptRouter) InitIdcIpPreemptRouter(Router *gin.RouterGroup) {
	idcIpPreemptRouter := Router.Group("idcIpPreempt").Use(middleware.OperationRecord())
	idcIpPreemptRouterWithoutRecord := Router.Group("idcIpPreempt")
	var idcIpPreemptApi = v1.ApiGroupApp.IdcApiGroup.IdcIpPreemptApi
	{
		idcIpPreemptRouter.POST("createIdcIpPreempt", idcIpPreemptApi.CreateIdcIpPreempt)   // 新建数据中心地址预占
		idcIpPreemptRouter.DELETE("deleteIdcIpPreempt", idcIpPreemptApi.DeleteIdcIpPreempt) // 删除数据中心地址预占
		idcIpPreemptRouter.DELETE("deleteIdcIpPreemptByIds", idcIpPreemptApi.DeleteIdcIpPreemptByIds) // 批量删除数据中心地址预占
		idcIpPreemptRouter.PUT("updateIdcIpPreempt", idcIpPreemptApi.UpdateIdcIpPreempt)    // 更新数据中心地址预占
	}
	{
		idcIpPreemptRouterWithoutRecord.GET("findIdcIpPreempt", idcIpPreemptApi.FindIdcIpPreempt)        // 根据ID获取数据中心地址预占
		idcIpPreemptRouterWithoutRecord.GET("getIdcIpPreemptList", idcIpPreemptApi.GetIdcIpPreemptList)  // 获取数据中心地址预占列表
	}
}
