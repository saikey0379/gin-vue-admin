package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdcInfoRouter struct {
}

// InitIdcInfoRouter 初始化 数据中心 路由信息
func (s *IdcInfoRouter) InitIdcInfoRouter(Router *gin.RouterGroup) {
	idcInfoRouter := Router.Group("idcInfo").Use(middleware.OperationRecord())
	idcInfoRouterWithoutRecord := Router.Group("idcInfo")
	var idcInfoApi = v1.ApiGroupApp.IdcApiGroup.IdcInfoApi
	{
		idcInfoRouter.POST("createIdcInfo", idcInfoApi.CreateIdcInfo)             // 新建数据中心
		idcInfoRouter.DELETE("deleteIdcInfo", idcInfoApi.DeleteIdcInfo)           // 删除数据中心
		idcInfoRouter.DELETE("deleteIdcInfoByIds", idcInfoApi.DeleteIdcInfoByIds) // 批量删除数据中心
		idcInfoRouter.PUT("updateIdcInfo", idcInfoApi.UpdateIdcInfo)              // 更新数据中心
	}
	{
		idcInfoRouterWithoutRecord.GET("findIdcInfo", idcInfoApi.FindIdcInfo)       // 根据ID获取数据中心
		idcInfoRouterWithoutRecord.GET("getIdcInfoList", idcInfoApi.GetIdcInfoList) // 获取数据中心列表
	}
}
