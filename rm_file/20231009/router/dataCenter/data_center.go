package dataCenter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DataCenterRouter struct {
}

// InitDataCenterRouter 初始化 数据中心 路由信息
func (s *DataCenterRouter) InitDataCenterRouter(Router *gin.RouterGroup) {
	dcRouter := Router.Group("dc").Use(middleware.OperationRecord())
	dcRouterWithoutRecord := Router.Group("dc")
	var dcApi = v1.ApiGroupApp.DataCenterApiGroup.DataCenterApi
	{
		dcRouter.POST("createDataCenter", dcApi.CreateDataCenter)   // 新建数据中心
		dcRouter.DELETE("deleteDataCenter", dcApi.DeleteDataCenter) // 删除数据中心
		dcRouter.DELETE("deleteDataCenterByIds", dcApi.DeleteDataCenterByIds) // 批量删除数据中心
		dcRouter.PUT("updateDataCenter", dcApi.UpdateDataCenter)    // 更新数据中心
	}
	{
		dcRouterWithoutRecord.GET("findDataCenter", dcApi.FindDataCenter)        // 根据ID获取数据中心
		dcRouterWithoutRecord.GET("getDataCenterList", dcApi.GetDataCenterList)  // 获取数据中心列表
	}
}
