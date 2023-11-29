package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileCommadBlacklistRouter struct {
}

// InitFileCommadBlacklistRouter 初始化 命令黑名单 路由信息
func (s *FileCommadBlacklistRouter) InitFileCommadBlacklistRouter(Router *gin.RouterGroup) {
	fileCommandBlacklistRouter := Router.Group("fileCommandBlacklist").Use(middleware.OperationRecord())
	fileCommandBlacklistRouterWithoutRecord := Router.Group("fileCommandBlacklist")
	var fileCommandBlacklistApi = v1.ApiGroupApp.FileApiGroup.FileCommadBlacklistApi
	{
		fileCommandBlacklistRouter.POST("createFileCommadBlacklist", fileCommandBlacklistApi.CreateFileCommadBlacklist)   // 新建命令黑名单
		fileCommandBlacklistRouter.DELETE("deleteFileCommadBlacklist", fileCommandBlacklistApi.DeleteFileCommadBlacklist) // 删除命令黑名单
		fileCommandBlacklistRouter.DELETE("deleteFileCommadBlacklistByIds", fileCommandBlacklistApi.DeleteFileCommadBlacklistByIds) // 批量删除命令黑名单
		fileCommandBlacklistRouter.PUT("updateFileCommadBlacklist", fileCommandBlacklistApi.UpdateFileCommadBlacklist)    // 更新命令黑名单
	}
	{
		fileCommandBlacklistRouterWithoutRecord.GET("findFileCommadBlacklist", fileCommandBlacklistApi.FindFileCommadBlacklist)        // 根据ID获取命令黑名单
		fileCommandBlacklistRouterWithoutRecord.GET("getFileCommadBlacklistList", fileCommandBlacklistApi.GetFileCommadBlacklistList)  // 获取命令黑名单列表
	}
}
