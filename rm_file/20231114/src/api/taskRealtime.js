import service from '@/utils/request'

// @Tags TaskRealtime
// @Summary 创建实时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskRealtime true "创建实时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /taskRealtime/createTaskRealtime [post]
export const createTaskRealtime = (data) => {
  return service({
    url: '/taskRealtime/createTaskRealtime',
    method: 'post',
    data
  })
}

// @Tags TaskRealtime
// @Summary 删除实时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskRealtime true "删除实时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskRealtime/deleteTaskRealtime [delete]
export const deleteTaskRealtime = (data) => {
  return service({
    url: '/taskRealtime/deleteTaskRealtime',
    method: 'delete',
    data
  })
}

// @Tags TaskRealtime
// @Summary 批量删除实时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除实时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskRealtime/deleteTaskRealtime [delete]
export const deleteTaskRealtimeByIds = (data) => {
  return service({
    url: '/taskRealtime/deleteTaskRealtimeByIds',
    method: 'delete',
    data
  })
}

// @Tags TaskRealtime
// @Summary 更新实时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskRealtime true "更新实时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /taskRealtime/updateTaskRealtime [put]
export const updateTaskRealtime = (data) => {
  return service({
    url: '/taskRealtime/updateTaskRealtime',
    method: 'put',
    data
  })
}

// @Tags TaskRealtime
// @Summary 用id查询实时任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TaskRealtime true "用id查询实时任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /taskRealtime/findTaskRealtime [get]
export const findTaskRealtime = (params) => {
  return service({
    url: '/taskRealtime/findTaskRealtime',
    method: 'get',
    params
  })
}

// @Tags TaskRealtime
// @Summary 分页获取实时任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取实时任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /taskRealtime/getTaskRealtimeList [get]
export const getTaskRealtimeList = (params) => {
  return service({
    url: '/taskRealtime/getTaskRealtimeList',
    method: 'get',
    params
  })
}
