import service from '@/utils/request'

// @Tags TaskTimed
// @Summary 创建定时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskTimed true "创建定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /taskTimed/createTaskTimed [post]
export const createTaskTimed = (data) => {
  return service({
    url: '/taskTimed/createTaskTimed',
    method: 'post',
    data
  })
}

// @Tags TaskTimed
// @Summary 删除定时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskTimed true "删除定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskTimed/deleteTaskTimed [delete]
export const deleteTaskTimed = (data) => {
  return service({
    url: '/taskTimed/deleteTaskTimed',
    method: 'delete',
    data
  })
}

// @Tags TaskTimed
// @Summary 批量删除定时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskTimed/deleteTaskTimed [delete]
export const deleteTaskTimedByIds = (data) => {
  return service({
    url: '/taskTimed/deleteTaskTimedByIds',
    method: 'delete',
    data
  })
}

// @Tags TaskTimed
// @Summary 更新定时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskTimed true "更新定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /taskTimed/updateTaskTimed [put]
export const updateTaskTimed = (data) => {
  return service({
    url: '/taskTimed/updateTaskTimed',
    method: 'put',
    data
  })
}

// @Tags TaskTimed
// @Summary 用id查询定时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TaskTimed true "用id查询定时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /taskTimed/findTaskTimed [get]
export const findTaskTimed = (params) => {
  return service({
    url: '/taskTimed/findTaskTimed',
    method: 'get',
    params
  })
}

// @Tags TaskTimed
// @Summary 分页获取定时任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取定时任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /taskTimed/getTaskTimedList [get]
export const getTaskTimedList = (params) => {
  return service({
    url: '/taskTimed/getTaskTimedList',
    method: 'get',
    params
  })
}
