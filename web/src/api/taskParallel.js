import service from '@/utils/request'

// @Tags TaskParallel
// @Summary 创建并行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskParallel true "创建并行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /taskParallel/createTaskParallel [post]
export const createTaskParallel = (data) => {
  return service({
    url: '/taskParallel/createTaskParallel',
    method: 'post',
    data
  })
}

// @Tags TaskParallel
// @Summary 删除并行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskParallel true "删除并行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskParallel/deleteTaskParallel [delete]
export const deleteTaskParallel = (data) => {
  return service({
    url: '/taskParallel/deleteTaskParallel',
    method: 'delete',
    data
  })
}

// @Tags TaskParallel
// @Summary 批量删除并行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除并行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskParallel/deleteTaskParallel [delete]
export const deleteTaskParallelByIds = (data) => {
  return service({
    url: '/taskParallel/deleteTaskParallelByIds',
    method: 'delete',
    data
  })
}

// @Tags TaskParallel
// @Summary 更新并行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskParallel true "更新并行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /taskParallel/updateTaskParallel [put]
export const updateTaskParallel = (data) => {
  return service({
    url: '/taskParallel/updateTaskParallel',
    method: 'put',
    data
  })
}

// @Tags TaskParallel
// @Summary 用id查询并行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TaskParallel true "用id查询并行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /taskParallel/findTaskParallel [get]
export const findTaskParallel = (params) => {
  return service({
    url: '/taskParallel/findTaskParallel',
    method: 'get',
    params
  })
}

// @Tags TaskParallel
// @Summary 分页获取并行任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取并行任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /taskParallel/getTaskParallelList [get]
export const getTaskParallelList = (params) => {
  return service({
    url: '/taskParallel/getTaskParallelList',
    method: 'get',
    params
  })
}
