package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileBinaryRouter struct {
}

// InitFileBinaryRouter 初始化 可执行程序 路由信息
func (s *FileBinaryRouter) InitFileBinaryRouter(Router *gin.RouterGroup) {
	fileBinaryRouter := Router.Group("fileBinary").Use(middleware.OperationRecord())
	fileBinaryRouterWithoutRecord := Router.Group("fileBinary")
	var fileBinaryApi = v1.ApiGroupApp.FileApiGroup.FileBinaryApi
	{
		fileBinaryRouter.POST("createFileBinary", fileBinaryApi.CreateFileBinary)   // 新建可执行程序
		fileBinaryRouter.DELETE("deleteFileBinary", fileBinaryApi.DeleteFileBinary) // 删除可执行程序
		fileBinaryRouter.DELETE("deleteFileBinaryByIds", fileBinaryApi.DeleteFileBinaryByIds) // 批量删除可执行程序
		fileBinaryRouter.PUT("updateFileBinary", fileBinaryApi.UpdateFileBinary)    // 更新可执行程序
	}
	{
		fileBinaryRouterWithoutRecord.GET("findFileBinary", fileBinaryApi.FindFileBinary)        // 根据ID获取可执行程序
		fileBinaryRouterWithoutRecord.GET("getFileBinaryList", fileBinaryApi.GetFileBinaryList)  // 获取可执行程序列表
	}
}
