package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/task"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    taskReq "github.com/flipped-aurora/gin-vue-admin/server/model/task/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type TaskCrontabApi struct {
}

var taskCrontabService = service.ServiceGroupApp.TaskServiceGroup.TaskCrontabService


// CreateTaskCrontab 创建计划任务
// @Tags TaskCrontab
// @Summary 创建计划任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.TaskCrontab true "创建计划任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /taskCrontab/createTaskCrontab [post]
func (taskCrontabApi *TaskCrontabApi) CreateTaskCrontab(c *gin.Context) {
	var taskCrontab task.TaskCrontab
	err := c.ShouldBindJSON(&taskCrontab)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Describe":{utils.NotEmpty()},
        "Nodes":{utils.NotEmpty()},
        "FileType":{utils.NotEmpty()},
        "FileId":{utils.NotEmpty()},
        "CronMinute":{utils.NotEmpty()},
        "CronHour":{utils.NotEmpty()},
        "CronDayOfMonth":{utils.NotEmpty()},
        "CronMonth":{utils.NotEmpty()},
        "CronDayOfWeek":{utils.NotEmpty()},
    }
	if err := utils.Verify(taskCrontab, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := taskCrontabService.CreateTaskCrontab(&taskCrontab); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTaskCrontab 删除计划任务
// @Tags TaskCrontab
// @Summary 删除计划任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.TaskCrontab true "删除计划任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskCrontab/deleteTaskCrontab [delete]
func (taskCrontabApi *TaskCrontabApi) DeleteTaskCrontab(c *gin.Context) {
	var taskCrontab task.TaskCrontab
	err := c.ShouldBindJSON(&taskCrontab)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := taskCrontabService.DeleteTaskCrontab(taskCrontab); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTaskCrontabByIds 批量删除计划任务
// @Tags TaskCrontab
// @Summary 批量删除计划任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除计划任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /taskCrontab/deleteTaskCrontabByIds [delete]
func (taskCrontabApi *TaskCrontabApi) DeleteTaskCrontabByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := taskCrontabService.DeleteTaskCrontabByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTaskCrontab 更新计划任务
// @Tags TaskCrontab
// @Summary 更新计划任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.TaskCrontab true "更新计划任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /taskCrontab/updateTaskCrontab [put]
func (taskCrontabApi *TaskCrontabApi) UpdateTaskCrontab(c *gin.Context) {
	var taskCrontab task.TaskCrontab
	err := c.ShouldBindJSON(&taskCrontab)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Describe":{utils.NotEmpty()},
          "Nodes":{utils.NotEmpty()},
          "FileType":{utils.NotEmpty()},
          "FileId":{utils.NotEmpty()},
          "CronMinute":{utils.NotEmpty()},
          "CronHour":{utils.NotEmpty()},
          "CronDayOfMonth":{utils.NotEmpty()},
          "CronMonth":{utils.NotEmpty()},
          "CronDayOfWeek":{utils.NotEmpty()},
      }
    if err := utils.Verify(taskCrontab, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := taskCrontabService.UpdateTaskCrontab(taskCrontab); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTaskCrontab 用id查询计划任务
// @Tags TaskCrontab
// @Summary 用id查询计划任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query task.TaskCrontab true "用id查询计划任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /taskCrontab/findTaskCrontab [get]
func (taskCrontabApi *TaskCrontabApi) FindTaskCrontab(c *gin.Context) {
	var taskCrontab task.TaskCrontab
	err := c.ShouldBindQuery(&taskCrontab)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retaskCrontab, err := taskCrontabService.GetTaskCrontab(taskCrontab.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retaskCrontab": retaskCrontab}, c)
	}
}

// GetTaskCrontabList 分页获取计划任务列表
// @Tags TaskCrontab
// @Summary 分页获取计划任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query taskReq.TaskCrontabSearch true "分页获取计划任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /taskCrontab/getTaskCrontabList [get]
func (taskCrontabApi *TaskCrontabApi) GetTaskCrontabList(c *gin.Context) {
	var pageInfo taskReq.TaskCrontabSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := taskCrontabService.GetTaskCrontabInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
