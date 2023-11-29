package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdcInfottRouter struct {
}

// InitIdcInfottRouter 初始化 数据中心 路由信息
func (s *IdcInfottRouter) InitIdcInfottRouter(Router *gin.RouterGroup) {
	idcInfottRouter := Router.Group("idcInfott").Use(middleware.OperationRecord())
	idcInfottRouterWithoutRecord := Router.Group("idcInfott")
	var idcInfottApi = v1.ApiGroupApp.IdcApiGroup.IdcInfottApi
	{
		idcInfottRouter.POST("createIdcInfott", idcInfottApi.CreateIdcInfott)   // 新建数据中心
		idcInfottRouter.DELETE("deleteIdcInfott", idcInfottApi.DeleteIdcInfott) // 删除数据中心
		idcInfottRouter.DELETE("deleteIdcInfottByIds", idcInfottApi.DeleteIdcInfottByIds) // 批量删除数据中心
		idcInfottRouter.PUT("updateIdcInfott", idcInfottApi.UpdateIdcInfott)    // 更新数据中心
	}
	{
		idcInfottRouterWithoutRecord.GET("findIdcInfott", idcInfottApi.FindIdcInfott)        // 根据ID获取数据中心
		idcInfottRouterWithoutRecord.GET("getIdcInfottList", idcInfottApi.GetIdcInfottList)  // 获取数据中心列表
	}
}
