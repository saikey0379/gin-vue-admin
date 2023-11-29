package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileConfigRouter struct {
}

// InitFileConfigRouter 初始化 配置文件 路由信息
func (s *FileConfigRouter) InitFileConfigRouter(Router *gin.RouterGroup) {
	fileConfigRouter := Router.Group("fileConfig").Use(middleware.OperationRecord())
	fileConfigRouterWithoutRecord := Router.Group("fileConfig")
	var fileConfigApi = v1.ApiGroupApp.FileApiGroup.FileConfigApi
	{
		fileConfigRouter.POST("createFileConfig", fileConfigApi.CreateFileConfig)             // 新建配置文件
		fileConfigRouter.DELETE("deleteFileConfig", fileConfigApi.DeleteFileConfig)           // 删除配置文件
		fileConfigRouter.DELETE("deleteFileConfigByIds", fileConfigApi.DeleteFileConfigByIds) // 批量删除配置文件
		fileConfigRouter.PUT("updateFileConfig", fileConfigApi.UpdateFileConfig)              // 更新配置文件
	}
	{
		fileConfigRouterWithoutRecord.GET("findFileConfig", fileConfigApi.FindFileConfig)       // 根据ID获取配置文件
		fileConfigRouterWithoutRecord.GET("getFileConfigList", fileConfigApi.GetFileConfigList) // 获取配置文件列表
	}
}

func (s *FileConfigRouter) InitFileConfigRouterPublic(Router *gin.RouterGroup) {
	fileConfigRouterWithoutRecord := Router.Group("fileConfig")
	var fileConfigApi = v1.ApiGroupApp.FileApiGroup.FileConfigApi
	{
		fileConfigRouterWithoutRecord.GET("getContent", fileConfigApi.GetContent) // 根据name获取Content
	}
}
