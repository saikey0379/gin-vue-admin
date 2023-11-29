package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskRealtimeRouter struct {
}

// InitTaskRealtimeRouter 初始化 实时任务 路由信息
func (s *TaskRealtimeRouter) InitTaskRealtimeRouter(Router *gin.RouterGroup) {
	taskRealtimeRouter := Router.Group("taskRealtime").Use(middleware.OperationRecord())
	taskRealtimeRouterWithoutRecord := Router.Group("taskRealtime")
	var taskRealtimeApi = v1.ApiGroupApp.TaskApiGroup.TaskRealtimeApi
	{
		taskRealtimeRouter.POST("createTaskRealtime", taskRealtimeApi.CreateTaskRealtime)   // 新建实时任务
		taskRealtimeRouter.DELETE("deleteTaskRealtime", taskRealtimeApi.DeleteTaskRealtime) // 删除实时任务
		taskRealtimeRouter.DELETE("deleteTaskRealtimeByIds", taskRealtimeApi.DeleteTaskRealtimeByIds) // 批量删除实时任务
		taskRealtimeRouter.PUT("updateTaskRealtime", taskRealtimeApi.UpdateTaskRealtime)    // 更新实时任务
	}
	{
		taskRealtimeRouterWithoutRecord.GET("findTaskRealtime", taskRealtimeApi.FindTaskRealtime)        // 根据ID获取实时任务
		taskRealtimeRouterWithoutRecord.GET("getTaskRealtimeList", taskRealtimeApi.GetTaskRealtimeList)  // 获取实时任务列表
	}
}
