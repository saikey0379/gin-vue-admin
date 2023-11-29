package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskParallelRouter struct {
}

// InitTaskParallelRouter 初始化 并行任务 路由信息
func (s *TaskParallelRouter) InitTaskParallelRouter(Router *gin.RouterGroup) {
	taskParallelRouter := Router.Group("taskParallel").Use(middleware.OperationRecord())
	taskParallelRouterWithoutRecord := Router.Group("taskParallel")
	var taskParallelApi = v1.ApiGroupApp.TaskApiGroup.TaskParallelApi
	{
		taskParallelRouter.POST("createTaskParallel", taskParallelApi.CreateTaskParallel)   // 新建并行任务
		taskParallelRouter.DELETE("deleteTaskParallel", taskParallelApi.DeleteTaskParallel) // 删除并行任务
		taskParallelRouter.DELETE("deleteTaskParallelByIds", taskParallelApi.DeleteTaskParallelByIds) // 批量删除并行任务
		taskParallelRouter.PUT("updateTaskParallel", taskParallelApi.UpdateTaskParallel)    // 更新并行任务
	}
	{
		taskParallelRouterWithoutRecord.GET("findTaskParallel", taskParallelApi.FindTaskParallel)        // 根据ID获取并行任务
		taskParallelRouterWithoutRecord.GET("getTaskParallelList", taskParallelApi.GetTaskParallelList)  // 获取并行任务列表
	}
}
