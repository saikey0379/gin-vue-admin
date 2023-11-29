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
	DeviceBareMetalRouter := Router.Group("DeviceBareMetal").Use(middleware.OperationRecord())
	DeviceBareMetalRouterWithoutRecord := Router.Group("DeviceBareMetal")
	var DeviceBareMetalApi = v1.ApiGroupApp.BareMetalApiGroup.BareMetalApi
	{
		DeviceBareMetalRouter.POST("createBareMetal", DeviceBareMetalApi.CreateBareMetal)   // 新建裸金属管理
		DeviceBareMetalRouter.DELETE("deleteBareMetal", DeviceBareMetalApi.DeleteBareMetal) // 删除裸金属管理
		DeviceBareMetalRouter.DELETE("deleteBareMetalByIds", DeviceBareMetalApi.DeleteBareMetalByIds) // 批量删除裸金属管理
		DeviceBareMetalRouter.PUT("updateBareMetal", DeviceBareMetalApi.UpdateBareMetal)    // 更新裸金属管理
	}
	{
		DeviceBareMetalRouterWithoutRecord.GET("findBareMetal", DeviceBareMetalApi.FindBareMetal)        // 根据ID获取裸金属管理
		DeviceBareMetalRouterWithoutRecord.GET("getBareMetalList", DeviceBareMetalApi.GetBareMetalList)  // 获取裸金属管理列表
	}
}
