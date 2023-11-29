package fileHandle

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type FileResumeTransRouter struct {
}

func (f *FileResumeTransRouter) InitFileResumeTramsRouter(Router *gin.RouterGroup) {
	fileResumeTransRouter := Router.Group("fileResumeTrans")
	fileResumeTransApi := v1.ApiGroupApp.FileHandleApiGroup.FileResumeTransApi
	{
		fileResumeTransRouter.GET("findFile", fileResumeTransApi.FindFile)            // 查询当前文件成功的切片
		fileResumeTransRouter.POST("chunkContinue", fileResumeTransApi.ChunkContinue) // 断点续传
		fileResumeTransRouter.POST("chunkFinish", fileResumeTransApi.ChunkFinish)     // 切片传输完成
		fileResumeTransRouter.POST("chunkRemove", fileResumeTransApi.ChunkRemove)     // 删除切片
	}
}
