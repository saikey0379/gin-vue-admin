package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdcInfotRouter struct {
}

// InitIdcInfotRouter 初始化 数据中心 路由信息
func (s *IdcInfotRouter) InitIdcInfotRouter(Router *gin.RouterGroup) {
	idcInfotRouter := Router.Group("idcInfot").Use(middleware.OperationRecord())
	idcInfotRouterWithoutRecord := Router.Group("idcInfot")
	var idcInfotApi = v1.ApiGroupApp.IdcApiGroup.IdcInfotApi
	{
		idcInfotRouter.POST("createIdcInfot", idcInfotApi.CreateIdcInfot)   // 新建数据中心
		idcInfotRouter.DELETE("deleteIdcInfot", idcInfotApi.DeleteIdcInfot) // 删除数据中心
		idcInfotRouter.DELETE("deleteIdcInfotByIds", idcInfotApi.DeleteIdcInfotByIds) // 批量删除数据中心
		idcInfotRouter.PUT("updateIdcInfot", idcInfotApi.UpdateIdcInfot)    // 更新数据中心
	}
	{
		idcInfotRouterWithoutRecord.GET("findIdcInfot", idcInfotApi.FindIdcInfot)        // 根据ID获取数据中心
		idcInfotRouterWithoutRecord.GET("getIdcInfotList", idcInfotApi.GetIdcInfotList)  // 获取数据中心列表
	}
}
