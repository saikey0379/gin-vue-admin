package server

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ServerDiscoveryRouter struct {
}

// InitServerDiscoveryRouter 初始化 Server发现 路由信息
func (s *ServerDiscoveryRouter) InitServerDiscoveryRouter(Router *gin.RouterGroup) {
	serverDiscoveryRouter := Router.Group("serverDiscovery").Use(middleware.OperationRecord())
	serverDiscoveryRouterWithoutRecord := Router.Group("serverDiscovery")
	var serverDiscoveryApi = v1.ApiGroupApp.ServerApiGroup.ServerDiscoveryApi
	{
		serverDiscoveryRouter.POST("createServerDiscovery", serverDiscoveryApi.CreateServerDiscovery)   // 新建Server发现
		serverDiscoveryRouter.DELETE("deleteServerDiscovery", serverDiscoveryApi.DeleteServerDiscovery) // 删除Server发现
		serverDiscoveryRouter.DELETE("deleteServerDiscoveryByIds", serverDiscoveryApi.DeleteServerDiscoveryByIds) // 批量删除Server发现
		serverDiscoveryRouter.PUT("updateServerDiscovery", serverDiscoveryApi.UpdateServerDiscovery)    // 更新Server发现
	}
	{
		serverDiscoveryRouterWithoutRecord.GET("findServerDiscovery", serverDiscoveryApi.FindServerDiscovery)        // 根据ID获取Server发现
		serverDiscoveryRouterWithoutRecord.GET("getServerDiscoveryList", serverDiscoveryApi.GetServerDiscoveryList)  // 获取Server发现列表
	}
}
