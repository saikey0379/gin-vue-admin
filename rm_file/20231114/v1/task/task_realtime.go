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

type TaskRealtimeApi struct {
}

var taskRealtimeService = service.ServiceGroupApp.TaskServiceGroup.TaskRealtimeService


// CreateTaskRealtime 创建实时任务
// @Tags TaskRealtime
// @Summary 创建实时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.TaskRealtime true "创建实时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /taskRealtime/createTaskRealtime [post]
func (taskRealtimeApi *TaskRealtimeApi) CreateTaskRealtime(c *gin.Context) {
	var taskRealtime task.TaskRealtime
	err := c.ShouldBindJSON(&taskRealtime)
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
    }
	if err := utils.Verify(taskRealtime, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := taskRealtimeService.CreateTaskRealtime(&taskRealtime); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTaskRealtime 删除实时任务
// @Tags TaskRealtime
// @Summary 删除实时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.TaskRealtime true "删除实时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskRealtime/deleteTaskRealtime [delete]
func (taskRealtimeApi *TaskRealtimeApi) DeleteTaskRealtime(c *gin.Context) {
	var taskRealtime task.TaskRealtime
	err := c.ShouldBindJSON(&taskRealtime)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := taskRealtimeService.DeleteTaskRealtime(taskRealtime); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTaskRealtimeByIds 批量删除实时任务
// @Tags TaskRealtime
// @Summary 批量删除实时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除实时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /taskRealtime/deleteTaskRealtimeByIds [delete]
func (taskRealtimeApi *TaskRealtimeApi) DeleteTaskRealtimeByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := taskRealtimeService.DeleteTaskRealtimeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTaskRealtime 更新实时任务
// @Tags TaskRealtime
// @Summary 更新实时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.TaskRealtime true "更新实时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /taskRealtime/updateTaskRealtime [put]
func (taskRealtimeApi *TaskRealtimeApi) UpdateTaskRealtime(c *gin.Context) {
	var taskRealtime task.TaskRealtime
	err := c.ShouldBindJSON(&taskRealtime)
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
      }
    if err := utils.Verify(taskRealtime, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := taskRealtimeService.UpdateTaskRealtime(taskRealtime); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTaskRealtime 用id查询实时任务
// @Tags TaskRealtime
// @Summary 用id查询实时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query task.TaskRealtime true "用id查询实时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /taskRealtime/findTaskRealtime [get]
func (taskRealtimeApi *TaskRealtimeApi) FindTaskRealtime(c *gin.Context) {
	var taskRealtime task.TaskRealtime
	err := c.ShouldBindQuery(&taskRealtime)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retaskRealtime, err := taskRealtimeService.GetTaskRealtime(taskRealtime.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retaskRealtime": retaskRealtime}, c)
	}
}

// GetTaskRealtimeList 分页获取实时任务列表
// @Tags TaskRealtime
// @Summary 分页获取实时任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query taskReq.TaskRealtimeSearch true "分页获取实时任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /taskRealtime/getTaskRealtimeList [get]
func (taskRealtimeApi *TaskRealtimeApi) GetTaskRealtimeList(c *gin.Context) {
	var pageInfo taskReq.TaskRealtimeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := taskRealtimeService.GetTaskRealtimeInfoList(pageInfo); err != nil {
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
