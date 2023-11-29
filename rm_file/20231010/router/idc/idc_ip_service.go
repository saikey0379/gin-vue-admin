package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdcIpServiceRouter struct {
}

// InitIdcIpServiceRouter 初始化 数据中心业务网段 路由信息
func (s *IdcIpServiceRouter) InitIdcIpServiceRouter(Router *gin.RouterGroup) {
	idcIpServiceRouter := Router.Group("idcIpService").Use(middleware.OperationRecord())
	idcIpServiceRouterWithoutRecord := Router.Group("idcIpService")
	var idcIpServiceApi = v1.ApiGroupApp.IdcApiGroup.IdcIpServiceApi
	{
		idcIpServiceRouter.POST("createIdcIpService", idcIpServiceApi.CreateIdcIpService)   // 新建数据中心业务网段
		idcIpServiceRouter.DELETE("deleteIdcIpService", idcIpServiceApi.DeleteIdcIpService) // 删除数据中心业务网段
		idcIpServiceRouter.DELETE("deleteIdcIpServiceByIds", idcIpServiceApi.DeleteIdcIpServiceByIds) // 批量删除数据中心业务网段
		idcIpServiceRouter.PUT("updateIdcIpService", idcIpServiceApi.UpdateIdcIpService)    // 更新数据中心业务网段
	}
	{
		idcIpServiceRouterWithoutRecord.GET("findIdcIpService", idcIpServiceApi.FindIdcIpService)        // 根据ID获取数据中心业务网段
		idcIpServiceRouterWithoutRecord.GET("getIdcIpServiceList", idcIpServiceApi.GetIdcIpServiceList)  // 获取数据中心业务网段列表
	}
}
