package cmdb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmdbIpSubnetRouter struct {
}

// InitCmdbIpSubnetRouter 初始化 子网管理 路由信息
func (s *CmdbIpSubnetRouter) InitCmdbIpSubnetRouter(Router *gin.RouterGroup) {
	cmdbIpSubnetRouter := Router.Group("cmdbIpSubnet").Use(middleware.OperationRecord())
	cmdbIpSubnetRouterWithoutRecord := Router.Group("cmdbIpSubnet")
	var cmdbIpSubnetApi = v1.ApiGroupApp.CmdbApiGroup.CmdbIpSubnetApi
	{
		cmdbIpSubnetRouter.POST("createCmdbIpSubnet", cmdbIpSubnetApi.CreateCmdbIpSubnet)   // 新建子网管理
		cmdbIpSubnetRouter.DELETE("deleteCmdbIpSubnet", cmdbIpSubnetApi.DeleteCmdbIpSubnet) // 删除子网管理
		cmdbIpSubnetRouter.DELETE("deleteCmdbIpSubnetByIds", cmdbIpSubnetApi.DeleteCmdbIpSubnetByIds) // 批量删除子网管理
		cmdbIpSubnetRouter.PUT("updateCmdbIpSubnet", cmdbIpSubnetApi.UpdateCmdbIpSubnet)    // 更新子网管理
	}
	{
		cmdbIpSubnetRouterWithoutRecord.GET("findCmdbIpSubnet", cmdbIpSubnetApi.FindCmdbIpSubnet)        // 根据ID获取子网管理
		cmdbIpSubnetRouterWithoutRecord.GET("getCmdbIpSubnetList", cmdbIpSubnetApi.GetCmdbIpSubnetList)  // 获取子网管理列表
		cmdbIpSubnetRouterWithoutRecord.GET("validateIpSubnet", cmdbIpSubnetApi.ValidateIpSubnet)        // 确认IP地址是否在子网内

	}
}

func (s *CmdbIpSubnetRouter) InitCmdbIpSubnetRouterPublic(Router *gin.RouterGroup) {
	cmdbIpSubnetRouterWithoutRecord := Router.Group("cmdbIpSubnetPublic")
	var cmdbIpSubnetApi = v1.ApiGroupApp.CmdbApiGroup.CmdbIpSubnetApi
	{
		cmdbIpSubnetRouterWithoutRecord.GET("validateIpSubnet", cmdbIpSubnetApi.ValidateIpSubnet)       // 确认IP地址是否在子网内
	}
}