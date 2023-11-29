package osInstall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OsInstallQueueRouter struct {
}

// InitOsInstallQueueRouter 初始化 操作系统安装队列 路由信息
func (s *OsInstallQueueRouter) InitOsInstallQueueRouter(Router *gin.RouterGroup) {
	queueRouter := Router.Group("queue").Use(middleware.OperationRecord())
	queueRouterWithoutRecord := Router.Group("queue")
	var queueApi = v1.ApiGroupApp.OsInstallApiGroup.OsInstallQueueApi
	{
		queueRouter.POST("createOsInstallQueue", queueApi.CreateOsInstallQueue)   // 新建操作系统安装队列
		queueRouter.DELETE("deleteOsInstallQueue", queueApi.DeleteOsInstallQueue) // 删除操作系统安装队列
		queueRouter.DELETE("deleteOsInstallQueueByIds", queueApi.DeleteOsInstallQueueByIds) // 批量删除操作系统安装队列
		queueRouter.PUT("updateOsInstallQueue", queueApi.UpdateOsInstallQueue)    // 更新操作系统安装队列
	}
	{
		queueRouterWithoutRecord.GET("findOsInstallQueue", queueApi.FindOsInstallQueue)        // 根据ID获取操作系统安装队列
		queueRouterWithoutRecord.GET("getOsInstallQueueList", queueApi.GetOsInstallQueueList)  // 获取操作系统安装队列列表
	}
}
