package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskTimedRouter struct {
}

// InitTaskTimedRouter 初始化 定时任务 路由信息
func (s *TaskTimedRouter) InitTaskTimedRouter(Router *gin.RouterGroup) {
	taskTimedRouter := Router.Group("taskTimed").Use(middleware.OperationRecord())
	taskTimedRouterWithoutRecord := Router.Group("taskTimed")
	var taskTimedApi = v1.ApiGroupApp.TaskApiGroup.TaskTimedApi
	{
		taskTimedRouter.POST("createTaskTimed", taskTimedApi.CreateTaskTimed)   // 新建定时任务
		taskTimedRouter.DELETE("deleteTaskTimed", taskTimedApi.DeleteTaskTimed) // 删除定时任务
		taskTimedRouter.DELETE("deleteTaskTimedByIds", taskTimedApi.DeleteTaskTimedByIds) // 批量删除定时任务
		taskTimedRouter.PUT("updateTaskTimed", taskTimedApi.UpdateTaskTimed)    // 更新定时任务
	}
	{
		taskTimedRouterWithoutRecord.GET("findTaskTimed", taskTimedApi.FindTaskTimed)        // 根据ID获取定时任务
		taskTimedRouterWithoutRecord.GET("getTaskTimedList", taskTimedApi.GetTaskTimedList)  // 获取定时任务列表
	}
}
