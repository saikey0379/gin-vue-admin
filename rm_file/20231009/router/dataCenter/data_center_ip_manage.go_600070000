package dataCenter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DataCenterIpManageRouter struct {
}

// InitDataCenterIpManageRouter 初始化 数据中心 路由信息
func (s *DataCenterIpManageRouter) InitDataCenterIpManageRouter(Router *gin.RouterGroup) {
	ipManageRouter := Router.Group("ipManage").Use(middleware.OperationRecord())
	ipManageRouterWithoutRecord := Router.Group("ipManage")
	var ipManageApi = v1.ApiGroupApp.DataCenterApiGroup.DataCenterIpManageApi
	{
		ipManageRouter.POST("createDataCenterIpManage", ipManageApi.CreateDataCenterIpManage)   // 新建数据中心
		ipManageRouter.DELETE("deleteDataCenterIpManage", ipManageApi.DeleteDataCenterIpManage) // 删除数据中心
		ipManageRouter.DELETE("deleteDataCenterIpManageByIds", ipManageApi.DeleteDataCenterIpManageByIds) // 批量删除数据中心
		ipManageRouter.PUT("updateDataCenterIpManage", ipManageApi.UpdateDataCenterIpManage)    // 更新数据中心
	}
	{
		ipManageRouterWithoutRecord.GET("findDataCenterIpManage", ipManageApi.FindDataCenterIpManage)        // 根据ID获取数据中心
		ipManageRouterWithoutRecord.GET("getDataCenterIpManageList", ipManageApi.GetDataCenterIpManageList)  // 获取数据中心列表
	}
}
