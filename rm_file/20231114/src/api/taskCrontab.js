import service from '@/utils/request'

// @Tags TaskCrontab
// @Summary 创建计划任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskCrontab true "创建计划任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /taskCrontab/createTaskCrontab [post]
export const createTaskCrontab = (data) => {
  return service({
    url: '/taskCrontab/createTaskCrontab',
    method: 'post',
    data
  })
}

// @Tags TaskCrontab
// @Summary 删除计划任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskCrontab true "删除计划任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskCrontab/deleteTaskCrontab [delete]
export const deleteTaskCrontab = (data) => {
  return service({
    url: '/taskCrontab/deleteTaskCrontab',
    method: 'delete',
    data
  })
}

// @Tags TaskCrontab
// @Summary 批量删除计划任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除计划任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskCrontab/deleteTaskCrontab [delete]
export const deleteTaskCrontabByIds = (data) => {
  return service({
    url: '/taskCrontab/deleteTaskCrontabByIds',
    method: 'delete',
    data
  })
}

// @Tags TaskCrontab
// @Summary 更新计划任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskCrontab true "更新计划任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /taskCrontab/updateTaskCrontab [put]
export const updateTaskCrontab = (data) => {
  return service({
    url: '/taskCrontab/updateTaskCrontab',
    method: 'put',
    data
  })
}

// @Tags TaskCrontab
// @Summary 用id查询计划任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TaskCrontab true "用id查询计划任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /taskCrontab/findTaskCrontab [get]
export const findTaskCrontab = (params) => {
  return service({
    url: '/taskCrontab/findTaskCrontab',
    method: 'get',
    params
  })
}

// @Tags TaskCrontab
// @Summary 分页获取计划任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取计划任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /taskCrontab/getTaskCrontabList [get]
export const getTaskCrontabList = (params) => {
  return service({
    url: '/taskCrontab/getTaskCrontabList',
    method: 'get',
    params
  })
}
