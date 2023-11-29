package monit

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PrometheusClusterRouter struct {
}

// InitPrometheusClusterRouter 初始化 监控集群 路由信息
func (s *PrometheusClusterRouter) InitPrometheusClusterRouter(Router *gin.RouterGroup) {
	prometheusClusterRouter := Router.Group("prometheusCluster").Use(middleware.OperationRecord())
	prometheusClusterRouterWithoutRecord := Router.Group("prometheusCluster")
	var prometheusClusterApi = v1.ApiGroupApp.MonitApiGroup.PrometheusClusterApi
	{
		prometheusClusterRouter.POST("createPrometheusCluster", prometheusClusterApi.CreatePrometheusCluster)             // 新建监控集群
		prometheusClusterRouter.DELETE("deletePrometheusCluster", prometheusClusterApi.DeletePrometheusCluster)           // 删除监控集群
		prometheusClusterRouter.DELETE("deletePrometheusClusterByIds", prometheusClusterApi.DeletePrometheusClusterByIds) // 批量删除监控集群
		prometheusClusterRouter.PUT("updatePrometheusCluster", prometheusClusterApi.UpdatePrometheusCluster)              // 更新监控集群
	}
	{
		prometheusClusterRouterWithoutRecord.GET("findPrometheusCluster", prometheusClusterApi.FindPrometheusCluster)                               // 根据ID获取监控集群
		prometheusClusterRouterWithoutRecord.GET("getPrometheusClusterList", prometheusClusterApi.GetPrometheusClusterList)                         // 获取监控集群列表
		prometheusClusterRouterWithoutRecord.GET("findPromYamlListByClusterId", prometheusClusterApi.FindPromYamlListByClusterId)                   // 根据ID获取监控集群规则YamlList
		prometheusClusterRouterWithoutRecord.GET("findPromYamlDetailByClusterId", prometheusClusterApi.FindPromYamlDetailByClusterId)               // 根据ID获取监控集群规则YamlGlobal
		prometheusClusterRouterWithoutRecord.GET("findPromYamlDetailCurrentByClusterId", prometheusClusterApi.FindPromYamlDetailCurrentByClusterId) // 根据ID获取监控集群规则YamlGlobal
		prometheusClusterRouterWithoutRecord.GET("diffPromYamlByClusterId", prometheusClusterApi.DiffPromYamlByClusterId)                           // 根据ID获取监控集群规则YamlGlobal
		prometheusClusterRouterWithoutRecord.PUT("updatePrometheusRulesByClusterId", prometheusClusterApi.UpdatePrometheusRulesByClusterId)         // 更新监控集群配置

	}
}
