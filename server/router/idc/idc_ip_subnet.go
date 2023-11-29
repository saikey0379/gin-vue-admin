package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdcIpSubnetRouter struct {
}

// InitIdcIpSubnetRouter 初始化 数据中心子网 路由信息
func (s *IdcIpSubnetRouter) InitIdcIpSubnetRouter(Router *gin.RouterGroup) {
	idcIpSubnetRouter := Router.Group("idcIpSubnet").Use(middleware.OperationRecord())
	idcIpSubnetRouterWithoutRecord := Router.Group("idcIpSubnet")
	var idcIpSubnetApi = v1.ApiGroupApp.IdcApiGroup.IdcIpSubnetApi
	{
		idcIpSubnetRouter.POST("createIdcIpSubnet", idcIpSubnetApi.CreateIdcIpSubnet)             // 新建数据中心子网
		idcIpSubnetRouter.DELETE("deleteIdcIpSubnet", idcIpSubnetApi.DeleteIdcIpSubnet)           // 删除数据中心子网
		idcIpSubnetRouter.DELETE("deleteIdcIpSubnetByIds", idcIpSubnetApi.DeleteIdcIpSubnetByIds) // 批量删除数据中心子网
		idcIpSubnetRouter.PUT("updateIdcIpSubnet", idcIpSubnetApi.UpdateIdcIpSubnet)              // 更新数据中心子网
	}
	{
		idcIpSubnetRouterWithoutRecord.GET("findIdcIpSubnet", idcIpSubnetApi.FindIdcIpSubnet)       // 根据ID获取数据中心子网
		idcIpSubnetRouterWithoutRecord.GET("getIdcIpSubnetList", idcIpSubnetApi.GetIdcIpSubnetList) // 获取数据中心子网列表
		idcIpSubnetRouterWithoutRecord.GET("validateIpSubnet", idcIpSubnetApi.ValidateIpSubnet)     // 确认IP地址是否在子网内
	}
}

func (s *IdcIpSubnetRouter) InitIdcIpSubnetRouterPublic(Router *gin.RouterGroup) {
	idcIpSubnetRouterWithoutRecord := Router.Group("idcIpSubnetPublic")
	var idcIpSubnetApi = v1.ApiGroupApp.IdcApiGroup.IdcIpSubnetApi
	{
		idcIpSubnetRouterWithoutRecord.GET("validateIpSubnet", idcIpSubnetApi.ValidateIpSubnet) // 确认IP地址是否在子网内
	}
}
