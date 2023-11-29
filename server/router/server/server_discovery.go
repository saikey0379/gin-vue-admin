package server

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ServerDiscoveryRouter struct {
}

// InitServerDiscoveryRouter 初始化 ServerDiscovery 路由信息
func (s *ServerDiscoveryRouter) InitServerDiscoveryRouter(Router *gin.RouterGroup) {
	serverDiscoveryRouter := Router.Group("serverDiscovery").Use(middleware.OperationRecord())
	serverDiscoveryRouterWithoutRecord := Router.Group("serverDiscovery")
	var serverDiscoveryApi = v1.ApiGroupApp.ServerApiGroup.ServerDiscoveryApi
	{
		serverDiscoveryRouter.POST("createServerDiscovery", serverDiscoveryApi.CreateServerDiscovery)             // 新建ServerDiscovery
		serverDiscoveryRouter.DELETE("deleteServerDiscovery", serverDiscoveryApi.DeleteServerDiscovery)           // 删除ServerDiscovery
		serverDiscoveryRouter.DELETE("deleteServerDiscoveryByIds", serverDiscoveryApi.DeleteServerDiscoveryByIds) // 批量删除ServerDiscovery
		serverDiscoveryRouter.PUT("updateServerDiscovery", serverDiscoveryApi.UpdateServerDiscovery)              // 更新ServerDiscovery
		serverDiscoveryRouter.POST("entryServerDiscoveryByIds", serverDiscoveryApi.EntryServerDiscoveryByIds)     // 新建ServerDiscovery

	}
	{
		serverDiscoveryRouterWithoutRecord.GET("findServerDiscovery", serverDiscoveryApi.FindServerDiscovery)       // 根据ID获取ServerDiscovery
		serverDiscoveryRouterWithoutRecord.GET("getServerDiscoveryList", serverDiscoveryApi.GetServerDiscoveryList) // 获取ServerDiscovery列表
	}
}
