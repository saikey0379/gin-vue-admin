package idc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdcRoomRouter struct {
}

// InitIdcRoomRouter 初始化 数据中心房间 路由信息
func (s *IdcRoomRouter) InitIdcRoomRouter(Router *gin.RouterGroup) {
	roomRouter := Router.Group("room").Use(middleware.OperationRecord())
	roomRouterWithoutRecord := Router.Group("room")
	var roomApi = v1.ApiGroupApp.IdcApiGroup.IdcRoomApi
	{
		roomRouter.POST("createIdcRoom", roomApi.CreateIdcRoom)             // 新建数据中心房间
		roomRouter.DELETE("deleteIdcRoom", roomApi.DeleteIdcRoom)           // 删除数据中心房间
		roomRouter.DELETE("deleteIdcRoomByIds", roomApi.DeleteIdcRoomByIds) // 批量删除数据中心房间
		roomRouter.PUT("updateIdcRoom", roomApi.UpdateIdcRoom)              // 更新数据中心房间
	}
	{
		roomRouterWithoutRecord.GET("findIdcRoom", roomApi.FindIdcRoom)       // 根据ID获取数据中心房间
		roomRouterWithoutRecord.GET("getIdcRoomList", roomApi.GetIdcRoomList) // 获取数据中心房间列表
	}
}
