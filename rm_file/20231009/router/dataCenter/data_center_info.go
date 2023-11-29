package dataCenter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DataCenterInfoRouter struct {
}

// InitDataCenterInfoRouter 初始化 数据中心 路由信息
func (s *DataCenterInfoRouter) InitDataCenterInfoRouter(Router *gin.RouterGroup) {
	dcRouter := Router.Group("dc").Use(middleware.OperationRecord())
	dcRouterWithoutRecord := Router.Group("dc")
	var dcApi = v1.ApiGroupApp.DataCenterApiGroup.DataCenterInfoApi
	{
		dcRouter.POST("createDataCenterInfo", dcApi.CreateDataCenterInfo)   // 新建数据中心
		dcRouter.DELETE("deleteDataCenterInfo", dcApi.DeleteDataCenterInfo) // 删除数据中心
		dcRouter.DELETE("deleteDataCenterInfoByIds", dcApi.DeleteDataCenterInfoByIds) // 批量删除数据中心
		dcRouter.PUT("updateDataCenterInfo", dcApi.UpdateDataCenterInfo)    // 更新数据中心
	}
	{
		dcRouterWithoutRecord.GET("findDataCenterInfo", dcApi.FindDataCenterInfo)        // 根据ID获取数据中心
		dcRouterWithoutRecord.GET("getDataCenterInfoList", dcApi.GetDataCenterInfoList)  // 获取数据中心列表
	}
}
