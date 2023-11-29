package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BareMetalRouter struct {
}

// InitBareMetalRouter 初始化 裸金属设备 路由信息
func (s *BareMetalRouter) InitBareMetalRouter(Router *gin.RouterGroup) {
	bmRouter := Router.Group("bm").Use(middleware.OperationRecord())
	bmRouterWithoutRecord := Router.Group("bm")
	var bmApi = v1.ApiGroupApp.DeviceApiGroup.BareMetalApi
	{
		bmRouter.POST("createBareMetal", bmApi.CreateBareMetal)   // 新建裸金属设备
		bmRouter.DELETE("deleteBareMetal", bmApi.DeleteBareMetal) // 删除裸金属设备
		bmRouter.DELETE("deleteBareMetalByIds", bmApi.DeleteBareMetalByIds) // 批量删除裸金属设备
		bmRouter.PUT("updateBareMetal", bmApi.UpdateBareMetal)    // 更新裸金属设备
	}
	{
		bmRouterWithoutRecord.GET("findBareMetal", bmApi.FindBareMetal)        // 根据ID获取裸金属设备
		bmRouterWithoutRecord.GET("getBareMetalList", bmApi.GetBareMetalList)  // 获取裸金属设备列表
	}
}
