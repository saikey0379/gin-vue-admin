package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileScriptRouter struct {
}

// InitFileScriptRouter 初始化 脚本文件 路由信息
func (s *FileScriptRouter) InitFileScriptRouter(Router *gin.RouterGroup) {
	fileScriptRouter := Router.Group("fileScript").Use(middleware.OperationRecord())
	fileScriptRouterWithoutRecord := Router.Group("fileScript")
	var fileScriptApi = v1.ApiGroupApp.FileApiGroup.FileScriptApi
	{
		fileScriptRouter.POST("createFileScript", fileScriptApi.CreateFileScript)             // 新建脚本文件
		fileScriptRouter.DELETE("deleteFileScript", fileScriptApi.DeleteFileScript)           // 删除脚本文件
		fileScriptRouter.DELETE("deleteFileScriptByIds", fileScriptApi.DeleteFileScriptByIds) // 批量删除脚本文件
		fileScriptRouter.PUT("updateFileScript", fileScriptApi.UpdateFileScript)              // 更新脚本文件
	}
	{
		fileScriptRouterWithoutRecord.GET("findFileScript", fileScriptApi.FindFileScript)       // 根据ID获取脚本文件
		fileScriptRouterWithoutRecord.GET("getFileScriptList", fileScriptApi.GetFileScriptList) // 获取脚本文件列表
	}
}

func (s *FileScriptRouter) InitFileScriptRouterPublic(Router *gin.RouterGroup) {
	fileScriptRouterWithoutRecord := Router.Group("fileScript")
	var fileScriptApi = v1.ApiGroupApp.FileApiGroup.FileScriptApi
	{
		fileScriptRouterWithoutRecord.GET("getContent", fileScriptApi.GetContent) // 根据name获取Content
	}
}
