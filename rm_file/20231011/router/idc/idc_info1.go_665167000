package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdcInfo1Router struct {
}

// InitIdcInfo1Router 初始化 数据中心 路由信息
func (s *IdcInfo1Router) InitIdcInfo1Router(Router *gin.RouterGroup) {
	idcInfo1Router := Router.Group("idcInfo1").Use(middleware.OperationRecord())
	idcInfo1RouterWithoutRecord := Router.Group("idcInfo1")
	var idcInfo1Api = v1.ApiGroupApp.IdcApiGroup.IdcInfo1Api
	{
		idcInfo1Router.POST("createIdcInfo1", idcInfo1Api.CreateIdcInfo1)   // 新建数据中心
		idcInfo1Router.DELETE("deleteIdcInfo1", idcInfo1Api.DeleteIdcInfo1) // 删除数据中心
		idcInfo1Router.DELETE("deleteIdcInfo1ByIds", idcInfo1Api.DeleteIdcInfo1ByIds) // 批量删除数据中心
		idcInfo1Router.PUT("updateIdcInfo1", idcInfo1Api.UpdateIdcInfo1)    // 更新数据中心
	}
	{
		idcInfo1RouterWithoutRecord.GET("findIdcInfo1", idcInfo1Api.FindIdcInfo1)        // 根据ID获取数据中心
		idcInfo1RouterWithoutRecord.GET("getIdcInfo1List", idcInfo1Api.GetIdcInfo1List)  // 获取数据中心列表
	}
}
