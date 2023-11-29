package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileCommonRouter struct {
}

// InitFileCommonRouter 初始化 普通文件 路由信息
func (s *FileCommonRouter) InitFileCommonRouter(Router *gin.RouterGroup) {
	fileCommonRouter := Router.Group("fileCommon").Use(middleware.OperationRecord())
	fileCommonRouterWithoutRecord := Router.Group("fileCommon")
	var fileCommonApi = v1.ApiGroupApp.FileApiGroup.FileCommonApi
	{
		fileCommonRouter.POST("createFileCommon", fileCommonApi.CreateFileCommon)   // 新建普通文件
		fileCommonRouter.DELETE("deleteFileCommon", fileCommonApi.DeleteFileCommon) // 删除普通文件
		fileCommonRouter.DELETE("deleteFileCommonByIds", fileCommonApi.DeleteFileCommonByIds) // 批量删除普通文件
		fileCommonRouter.PUT("updateFileCommon", fileCommonApi.UpdateFileCommon)    // 更新普通文件
	}
	{
		fileCommonRouterWithoutRecord.GET("findFileCommon", fileCommonApi.FindFileCommon)        // 根据ID获取普通文件
		fileCommonRouterWithoutRecord.GET("getFileCommonList", fileCommonApi.GetFileCommonList)  // 获取普通文件列表
	}
}
