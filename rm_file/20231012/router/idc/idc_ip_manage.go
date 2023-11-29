package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdcIpManageRouter struct {
}

// InitIdcIpManageRouter 初始化 数据中心管理网段 路由信息
func (s *IdcIpManageRouter) InitIdcIpManageRouter(Router *gin.RouterGroup) {
	idcIpManageRouter := Router.Group("idcIpManage").Use(middleware.OperationRecord())
	idcIpManageRouterWithoutRecord := Router.Group("idcIpManage")
	var idcIpManageApi = v1.ApiGroupApp.IdcApiGroup.IdcIpManageApi
	{
		idcIpManageRouter.POST("createIdcIpManage", idcIpManageApi.CreateIdcIpManage)             // 新建数据中心管理网段
		idcIpManageRouter.DELETE("deleteIdcIpManage", idcIpManageApi.DeleteIdcIpManage)           // 删除数据中心管理网段
		idcIpManageRouter.DELETE("deleteIdcIpManageByIds", idcIpManageApi.DeleteIdcIpManageByIds) // 批量删除数据中心管理网段
		idcIpManageRouter.PUT("updateIdcIpManage", idcIpManageApi.UpdateIdcIpManage)              // 更新数据中心管理网段
		idcIpManageRouter.GET("validateIpManage", idcIpManageApi.ValidateIpManage)                // 根据IP及网段ID判断是否可用
	}
	{
		idcIpManageRouterWithoutRecord.GET("findIdcIpManage", idcIpManageApi.FindIdcIpManage)       // 根据ID获取数据中心管理网段
		idcIpManageRouterWithoutRecord.GET("getIdcIpManageList", idcIpManageApi.GetIdcIpManageList) // 获取数据中心管理网段列表
	}
}
