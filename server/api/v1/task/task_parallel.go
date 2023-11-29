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

type TaskParallelApi struct {
}

var taskParallelService = service.ServiceGroupApp.TaskServiceGroup.TaskParallelService


// CreateTaskParallel 创建并行任务
// @Tags TaskParallel
// @Summary 创建并行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.TaskParallel true "创建并行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /taskParallel/createTaskParallel [post]
func (taskParallelApi *TaskParallelApi) CreateTaskParallel(c *gin.Context) {
	var taskParallel task.TaskParallel
	err := c.ShouldBindJSON(&taskParallel)
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
	if err := utils.Verify(taskParallel, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := taskParallelService.CreateTaskParallel(&taskParallel); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTaskParallel 删除并行任务
// @Tags TaskParallel
// @Summary 删除并行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.TaskParallel true "删除并行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskParallel/deleteTaskParallel [delete]
func (taskParallelApi *TaskParallelApi) DeleteTaskParallel(c *gin.Context) {
	var taskParallel task.TaskParallel
	err := c.ShouldBindJSON(&taskParallel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := taskParallelService.DeleteTaskParallel(taskParallel); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTaskParallelByIds 批量删除并行任务
// @Tags TaskParallel
// @Summary 批量删除并行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除并行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /taskParallel/deleteTaskParallelByIds [delete]
func (taskParallelApi *TaskParallelApi) DeleteTaskParallelByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := taskParallelService.DeleteTaskParallelByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTaskParallel 更新并行任务
// @Tags TaskParallel
// @Summary 更新并行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body task.TaskParallel true "更新并行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /taskParallel/updateTaskParallel [put]
func (taskParallelApi *TaskParallelApi) UpdateTaskParallel(c *gin.Context) {
	var taskParallel task.TaskParallel
	err := c.ShouldBindJSON(&taskParallel)
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
    if err := utils.Verify(taskParallel, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := taskParallelService.UpdateTaskParallel(taskParallel); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTaskParallel 用id查询并行任务
// @Tags TaskParallel
// @Summary 用id查询并行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query task.TaskParallel true "用id查询并行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /taskParallel/findTaskParallel [get]
func (taskParallelApi *TaskParallelApi) FindTaskParallel(c *gin.Context) {
	var taskParallel task.TaskParallel
	err := c.ShouldBindQuery(&taskParallel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retaskParallel, err := taskParallelService.GetTaskParallel(taskParallel.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retaskParallel": retaskParallel}, c)
	}
}

// GetTaskParallelList 分页获取并行任务列表
// @Tags TaskParallel
// @Summary 分页获取并行任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query taskReq.TaskParallelSearch true "分页获取并行任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /taskParallel/getTaskParallelList [get]
func (taskParallelApi *TaskParallelApi) GetTaskParallelList(c *gin.Context) {
	var pageInfo taskReq.TaskParallelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := taskParallelService.GetTaskParallelInfoList(pageInfo); err != nil {
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
