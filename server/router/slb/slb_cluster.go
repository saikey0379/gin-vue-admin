package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SlbClusterRouter struct {
}

// InitSlbClusterRouter 初始化 集群管理 路由信息
func (s *SlbClusterRouter) InitSlbClusterRouter(Router *gin.RouterGroup) {
	slbClusterRouter := Router.Group("slbCluster").Use(middleware.OperationRecord())
	slbClusterRouterWithoutRecord := Router.Group("slbCluster")
	var slbClusterApi = v1.ApiGroupApp.SlbApiGroup.SlbClusterApi
	{
		slbClusterRouter.POST("createSlbCluster", slbClusterApi.CreateSlbCluster)             // 新建集群管理
		slbClusterRouter.DELETE("deleteSlbCluster", slbClusterApi.DeleteSlbCluster)           // 删除集群管理
		slbClusterRouter.DELETE("deleteSlbClusterByIds", slbClusterApi.DeleteSlbClusterByIds) // 批量删除集群管理
		slbClusterRouter.PUT("updateSlbCluster", slbClusterApi.UpdateSlbCluster)              // 更新集群管理
	}
	{
		slbClusterRouterWithoutRecord.GET("findSlbCluster", slbClusterApi.FindSlbCluster)       // 根据ID获取集群管理
		slbClusterRouterWithoutRecord.GET("getSlbClusterList", slbClusterApi.GetSlbClusterList) // 获取集群管理列表
		slbClusterRouterWithoutRecord.GET("uploadSlbRsa", slbClusterApi.UploadSlbRsa)           // 上传Rsa证书
	}
}
