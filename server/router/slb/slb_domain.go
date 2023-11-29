package slb

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SlbDomainRouter struct {
}

// InitSlbDomainRouter 初始化 域名管理 路由信息
func (s *SlbDomainRouter) InitSlbDomainRouter(Router *gin.RouterGroup) {
	slbDomainRouter := Router.Group("slbDomain").Use(middleware.OperationRecord())
	slbDomainRouterWithoutRecord := Router.Group("slbDomain")
	var slbDomainApi = v1.ApiGroupApp.SlbApiGroup.SlbDomainApi
	{
		slbDomainRouter.POST("createSlbDomain", slbDomainApi.CreateSlbDomain)   // 新建域名管理
		slbDomainRouter.DELETE("deleteSlbDomain", slbDomainApi.DeleteSlbDomain) // 删除域名管理
		slbDomainRouter.DELETE("deleteSlbDomainByIds", slbDomainApi.DeleteSlbDomainByIds) // 批量删除域名管理
		slbDomainRouter.PUT("updateSlbDomain", slbDomainApi.UpdateSlbDomain)    // 更新域名管理
	}
	{
		slbDomainRouterWithoutRecord.GET("findSlbDomain", slbDomainApi.FindSlbDomain)        // 根据ID获取域名管理
		slbDomainRouterWithoutRecord.GET("getSlbDomainList", slbDomainApi.GetSlbDomainList)  // 获取域名管理列表
	}
}
