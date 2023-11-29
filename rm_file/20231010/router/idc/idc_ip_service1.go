package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdcIpService1Router struct {
}

// InitIdcIpService1Router 初始化 数据中心业务网段 路由信息
func (s *IdcIpService1Router) InitIdcIpService1Router(Router *gin.RouterGroup) {
	idcIpService1Router := Router.Group("idcIpService1").Use(middleware.OperationRecord())
	idcIpService1RouterWithoutRecord := Router.Group("idcIpService1")
	var idcIpService1Api = v1.ApiGroupApp.IdcApiGroup.IdcIpService1Api
	{
		idcIpService1Router.POST("createIdcIpService1", idcIpService1Api.CreateIdcIpService1)   // 新建数据中心业务网段
		idcIpService1Router.DELETE("deleteIdcIpService1", idcIpService1Api.DeleteIdcIpService1) // 删除数据中心业务网段
		idcIpService1Router.DELETE("deleteIdcIpService1ByIds", idcIpService1Api.DeleteIdcIpService1ByIds) // 批量删除数据中心业务网段
		idcIpService1Router.PUT("updateIdcIpService1", idcIpService1Api.UpdateIdcIpService1)    // 更新数据中心业务网段
	}
	{
		idcIpService1RouterWithoutRecord.GET("findIdcIpService1", idcIpService1Api.FindIdcIpService1)        // 根据ID获取数据中心业务网段
		idcIpService1RouterWithoutRecord.GET("getIdcIpService1List", idcIpService1Api.GetIdcIpService1List)  // 获取数据中心业务网段列表
	}
}
