package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdcCabinetRouter struct {
}

// InitIdcCabinetRouter 初始化 数据中心机柜 路由信息
func (s *IdcCabinetRouter) InitIdcCabinetRouter(Router *gin.RouterGroup) {
	idcCabinetRouter := Router.Group("idcCabinet").Use(middleware.OperationRecord())
	idcCabinetRouterWithoutRecord := Router.Group("idcCabinet")
	var idcCabinetApi = v1.ApiGroupApp.IdcApiGroup.IdcCabinetApi
	{
		idcCabinetRouter.POST("createIdcCabinet", idcCabinetApi.CreateIdcCabinet)   // 新建数据中心机柜
		idcCabinetRouter.DELETE("deleteIdcCabinet", idcCabinetApi.DeleteIdcCabinet) // 删除数据中心机柜
		idcCabinetRouter.DELETE("deleteIdcCabinetByIds", idcCabinetApi.DeleteIdcCabinetByIds) // 批量删除数据中心机柜
		idcCabinetRouter.PUT("updateIdcCabinet", idcCabinetApi.UpdateIdcCabinet)    // 更新数据中心机柜
	}
	{
		idcCabinetRouterWithoutRecord.GET("findIdcCabinet", idcCabinetApi.FindIdcCabinet)        // 根据ID获取数据中心机柜
		idcCabinetRouterWithoutRecord.GET("getIdcCabinetList", idcCabinetApi.GetIdcCabinetList)  // 获取数据中心机柜列表
	}
}
