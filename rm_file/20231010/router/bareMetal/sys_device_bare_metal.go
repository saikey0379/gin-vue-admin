package bareMetal

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BareMetalRouter struct {
}

// InitBareMetalRouter 初始化 裸金属管理 路由信息
func (s *BareMetalRouter) InitBareMetalRouter(Router *gin.RouterGroup) {
	bmRouter := Router.Group("bm").Use(middleware.OperationRecord())
	bmRouterWithoutRecord := Router.Group("bm")
	var bmApi = v1.ApiGroupApp.BareMetalApiGroup.BareMetalApi
	{
		bmRouter.POST("createBareMetal", bmApi.CreateBareMetal)             // 新建裸金属管理
		bmRouter.DELETE("deleteBareMetal", bmApi.DeleteBareMetal)           // 删除裸金属管理
		bmRouter.DELETE("deleteBareMetalByIds", bmApi.DeleteBareMetalByIds) // 批量删除裸金属管理
		bmRouter.PUT("updateBareMetal", bmApi.UpdateBareMetal)              // 更新裸金属管理
	}
	{
		bmRouterWithoutRecord.GET("findBareMetal", bmApi.FindBareMetal)       // 根据ID获取裸金属管理
		bmRouterWithoutRecord.GET("getBareMetalList", bmApi.GetBareMetalList) // 获取裸金属管理列表
	}
	{
		
	}
}
