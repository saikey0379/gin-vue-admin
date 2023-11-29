package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileInfoRouter struct {
}

// InitFileInfoRouter 初始化 PXE配置 路由信息
func (s *FileInfoRouter) InitFileInfoRouter(Router *gin.RouterGroup) {
	fileInfoRouter := Router.Group("fileInfo").Use(middleware.OperationRecord())
	fileInfoRouterWithoutRecord := Router.Group("fileInfo")
	var fileInfoApi = v1.ApiGroupApp.FileApiGroup.FileInfoApi
	{
		fileInfoRouter.POST("createFileInfo", fileInfoApi.CreateFileInfo)   // 新建PXE配置
		fileInfoRouter.DELETE("deleteFileInfo", fileInfoApi.DeleteFileInfo) // 删除PXE配置
		fileInfoRouter.DELETE("deleteFileInfoByIds", fileInfoApi.DeleteFileInfoByIds) // 批量删除PXE配置
		fileInfoRouter.PUT("updateFileInfo", fileInfoApi.UpdateFileInfo)    // 更新PXE配置
	}
	{
		fileInfoRouterWithoutRecord.GET("findFileInfo", fileInfoApi.FindFileInfo)        // 根据ID获取PXE配置
		fileInfoRouterWithoutRecord.GET("getFileInfoList", fileInfoApi.GetFileInfoList)  // 获取PXE配置列表
	}
}
