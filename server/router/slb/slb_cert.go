package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SlbCertRouter struct {
}

// InitSlbCertRouter 初始化 证书管理 路由信息
func (s *SlbCertRouter) InitSlbCertRouter(Router *gin.RouterGroup) {
	slbCertRouter := Router.Group("slbCert").Use(middleware.OperationRecord())
	slbCertRouterWithoutRecord := Router.Group("slbCert")
	var slbCertApi = v1.ApiGroupApp.SlbApiGroup.SlbCertApi
	{
		slbCertRouter.POST("createSlbCert", slbCertApi.CreateSlbCert)   // 新建证书管理
		slbCertRouter.DELETE("deleteSlbCert", slbCertApi.DeleteSlbCert) // 删除证书管理
		slbCertRouter.DELETE("deleteSlbCertByIds", slbCertApi.DeleteSlbCertByIds) // 批量删除证书管理
		slbCertRouter.PUT("updateSlbCert", slbCertApi.UpdateSlbCert)    // 更新证书管理
	}
	{
		slbCertRouterWithoutRecord.GET("findSlbCert", slbCertApi.FindSlbCert)        // 根据ID获取证书管理
		slbCertRouterWithoutRecord.GET("getSlbCertList", slbCertApi.GetSlbCertList)  // 获取证书管理列表
	}
}
