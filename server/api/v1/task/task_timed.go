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

type TaskTimedApi struct {
}

var taskTimedService = service.ServiceGroupApp.TaskServiceGroup.TaskTimedService


// CreateTaskTimed 创建定时任务
// @Tags TaskTimed
// @Summary 创建定时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.TaskTimed true "创建定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /taskTimed/createTaskTimed [post]
func (taskTimedApi *TaskTimedApi) CreateTaskTimed(c *gin.Context) {
	var taskTimed task.TaskTimed
	err := c.ShouldBindJSON(&taskTimed)
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
        "ExecutedAt":{utils.NotEmpty()},
    }
	if err := utils.Verify(taskTimed, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := taskTimedService.CreateTaskTimed(&taskTimed); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTaskTimed 删除定时任务
// @Tags TaskTimed
// @Summary 删除定时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.TaskTimed true "删除定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskTimed/deleteTaskTimed [delete]
func (taskTimedApi *TaskTimedApi) DeleteTaskTimed(c *gin.Context) {
	var taskTimed task.TaskTimed
	err := c.ShouldBindJSON(&taskTimed)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := taskTimedService.DeleteTaskTimed(taskTimed); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTaskTimedByIds 批量删除定时任务
// @Tags TaskTimed
// @Summary 批量删除定时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /taskTimed/deleteTaskTimedByIds [delete]
func (taskTimedApi *TaskTimedApi) DeleteTaskTimedByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := taskTimedService.DeleteTaskTimedByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTaskTimed 更新定时任务
// @Tags TaskTimed
// @Summary 更新定时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.TaskTimed true "更新定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /taskTimed/updateTaskTimed [put]
func (taskTimedApi *TaskTimedApi) UpdateTaskTimed(c *gin.Context) {
	var taskTimed task.TaskTimed
	err := c.ShouldBindJSON(&taskTimed)
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
          "ExecutedAt":{utils.NotEmpty()},
      }
    if err := utils.Verify(taskTimed, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := taskTimedService.UpdateTaskTimed(taskTimed); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTaskTimed 用id查询定时任务
// @Tags TaskTimed
// @Summary 用id查询定时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query task.TaskTimed true "用id查询定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /taskTimed/findTaskTimed [get]
func (taskTimedApi *TaskTimedApi) FindTaskTimed(c *gin.Context) {
	var taskTimed task.TaskTimed
	err := c.ShouldBindQuery(&taskTimed)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retaskTimed, err := taskTimedService.GetTaskTimed(taskTimed.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retaskTimed": retaskTimed}, c)
	}
}

// GetTaskTimedList 分页获取定时任务列表
// @Tags TaskTimed
// @Summary 分页获取定时任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query taskReq.TaskTimedSearch true "分页获取定时任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /taskTimed/getTaskTimedList [get]
func (taskTimedApi *TaskTimedApi) GetTaskTimedList(c *gin.Context) {
	var pageInfo taskReq.TaskTimedSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := taskTimedService.GetTaskTimedInfoList(pageInfo); err != nil {
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
