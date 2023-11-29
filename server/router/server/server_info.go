package server

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ServerInfoRouter struct {
}

// InitServerInfoRouter 初始化 ServerInfo 路由信息
func (s *ServerInfoRouter) InitServerInfoRouter(Router *gin.RouterGroup) {
	serverInfoRouter := Router.Group("serverInfo").Use(middleware.OperationRecord())
	serverInfoRouterWithoutRecord := Router.Group("serverInfo")
	var serverInfoApi = v1.ApiGroupApp.ServerApiGroup.ServerInfoApi
	{
		serverInfoRouter.POST("createServerInfo", serverInfoApi.CreateServerInfo)             // 新建ServerInfo
		serverInfoRouter.DELETE("deleteServerInfo", serverInfoApi.DeleteServerInfo)           // 删除ServerInfo
		serverInfoRouter.DELETE("deleteServerInfoByIds", serverInfoApi.DeleteServerInfoByIds) // 批量删除ServerInfo
		serverInfoRouter.PUT("updateServerInfo", serverInfoApi.UpdateServerInfo)              // 更新ServerInfo
	}
	{
		serverInfoRouterWithoutRecord.GET("findServerInfo", serverInfoApi.FindServerInfo)                                                   // 根据ID获取ServerInfo
		serverInfoRouterWithoutRecord.GET("getServerInfoList", serverInfoApi.GetServerInfoList)                                             // 获取ServerInfo列表
		serverInfoRouterWithoutRecord.GET("getServerInfoSnListNotExistsInBareMetal", serverInfoApi.GetServerInfoSnListNotExistsInBareMetal) // 获取未纳管裸金属的SN列表
	}
}

func (s *ServerInfoRouter) InitServerInfoRouterPublic(Router *gin.RouterGroup) {
	serverInfoRouterWithoutRecord := Router.Group("serverInfo")
	var serverInfoApi = v1.ApiGroupApp.ServerApiGroup.ServerInfoApi
	{
		serverInfoRouterWithoutRecord.POST("report", serverInfoApi.Report) // Server信息上报
	}
}
