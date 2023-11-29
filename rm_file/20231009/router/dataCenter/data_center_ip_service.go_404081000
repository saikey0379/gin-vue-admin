package dataCenter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DataCenterIpServiceRouter struct {
}

// InitDataCenterIpServiceRouter 初始化 数据中心 路由信息
func (s *DataCenterIpServiceRouter) InitDataCenterIpServiceRouter(Router *gin.RouterGroup) {
	ipServiceRouter := Router.Group("ipService").Use(middleware.OperationRecord())
	ipServiceRouterWithoutRecord := Router.Group("ipService")
	var ipServiceApi = v1.ApiGroupApp.DataCenterApiGroup.DataCenterIpServiceApi
	{
		ipServiceRouter.POST("createDataCenterIpService", ipServiceApi.CreateDataCenterIpService)   // 新建数据中心
		ipServiceRouter.DELETE("deleteDataCenterIpService", ipServiceApi.DeleteDataCenterIpService) // 删除数据中心
		ipServiceRouter.DELETE("deleteDataCenterIpServiceByIds", ipServiceApi.DeleteDataCenterIpServiceByIds) // 批量删除数据中心
		ipServiceRouter.PUT("updateDataCenterIpService", ipServiceApi.UpdateDataCenterIpService)    // 更新数据中心
	}
	{
		ipServiceRouterWithoutRecord.GET("findDataCenterIpService", ipServiceApi.FindDataCenterIpService)        // 根据ID获取数据中心
		ipServiceRouterWithoutRecord.GET("getDataCenterIpServiceList", ipServiceApi.GetDataCenterIpServiceList)  // 获取数据中心列表
	}
}
