package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PrometheusClusterRouter struct {
}

// InitPrometheusClusterRouter 初始化 集群管理 路由信息
func (s *PrometheusClusterRouter) InitPrometheusClusterRouter(Router *gin.RouterGroup) {
	prometheusClusterRouter := Router.Group("prometheusCluster").Use(middleware.OperationRecord())
	prometheusClusterRouterWithoutRecord := Router.Group("prometheusCluster")
	var prometheusClusterApi = v1.ApiGroupApp.MonitApiGroup.PrometheusClusterApi
	{
		prometheusClusterRouter.POST("createPrometheusCluster", prometheusClusterApi.CreatePrometheusCluster)   // 新建集群管理
		prometheusClusterRouter.DELETE("deletePrometheusCluster", prometheusClusterApi.DeletePrometheusCluster) // 删除集群管理
		prometheusClusterRouter.DELETE("deletePrometheusClusterByIds", prometheusClusterApi.DeletePrometheusClusterByIds) // 批量删除集群管理
		prometheusClusterRouter.PUT("updatePrometheusCluster", prometheusClusterApi.UpdatePrometheusCluster)    // 更新集群管理
	}
	{
		prometheusClusterRouterWithoutRecord.GET("findPrometheusCluster", prometheusClusterApi.FindPrometheusCluster)        // 根据ID获取集群管理
		prometheusClusterRouterWithoutRecord.GET("getPrometheusClusterList", prometheusClusterApi.GetPrometheusClusterList)  // 获取集群管理列表
	}
}
