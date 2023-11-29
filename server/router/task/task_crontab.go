package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskCrontabRouter struct {
}

// InitTaskCrontabRouter 初始化 计划任务 路由信息
func (s *TaskCrontabRouter) InitTaskCrontabRouter(Router *gin.RouterGroup) {
	taskCrontabRouter := Router.Group("taskCrontab").Use(middleware.OperationRecord())
	taskCrontabRouterWithoutRecord := Router.Group("taskCrontab")
	var taskCrontabApi = v1.ApiGroupApp.TaskApiGroup.TaskCrontabApi
	{
		taskCrontabRouter.POST("createTaskCrontab", taskCrontabApi.CreateTaskCrontab)   // 新建计划任务
		taskCrontabRouter.DELETE("deleteTaskCrontab", taskCrontabApi.DeleteTaskCrontab) // 删除计划任务
		taskCrontabRouter.DELETE("deleteTaskCrontabByIds", taskCrontabApi.DeleteTaskCrontabByIds) // 批量删除计划任务
		taskCrontabRouter.PUT("updateTaskCrontab", taskCrontabApi.UpdateTaskCrontab)    // 更新计划任务
	}
	{
		taskCrontabRouterWithoutRecord.GET("findTaskCrontab", taskCrontabApi.FindTaskCrontab)        // 根据ID获取计划任务
		taskCrontabRouterWithoutRecord.GET("getTaskCrontabList", taskCrontabApi.GetTaskCrontabList)  // 获取计划任务列表
	}
}
