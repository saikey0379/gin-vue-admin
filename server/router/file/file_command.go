package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileCommandRouter struct {
}

// InitFileCommandRouter 初始化 命令信息 路由信息
func (s *FileCommandRouter) InitFileCommandRouter(Router *gin.RouterGroup) {
	fileCommandRouter := Router.Group("fileCommand").Use(middleware.OperationRecord())
	fileCommandRouterWithoutRecord := Router.Group("fileCommand")
	var fileCommandApi = v1.ApiGroupApp.FileApiGroup.FileCommandApi
	{
		fileCommandRouter.POST("createFileCommand", fileCommandApi.CreateFileCommand)   // 新建命令信息
		fileCommandRouter.DELETE("deleteFileCommand", fileCommandApi.DeleteFileCommand) // 删除命令信息
		fileCommandRouter.DELETE("deleteFileCommandByIds", fileCommandApi.DeleteFileCommandByIds) // 批量删除命令信息
		fileCommandRouter.PUT("updateFileCommand", fileCommandApi.UpdateFileCommand)    // 更新命令信息
	}
	{
		fileCommandRouterWithoutRecord.GET("findFileCommand", fileCommandApi.FindFileCommand)        // 根据ID获取命令信息
		fileCommandRouterWithoutRecord.GET("getFileCommandList", fileCommandApi.GetFileCommandList)  // 获取命令信息列表
	}
}
