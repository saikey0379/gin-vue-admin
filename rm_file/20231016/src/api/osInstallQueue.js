import service from '@/utils/request'

// @Tags OsInstallQueue
// @Summary 创建操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallQueue true "创建操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /queue/createOsInstallQueue [post]
export const createOsInstallQueue = (data) => {
  return service({
    url: '/queue/createOsInstallQueue',
    method: 'post',
    data
  })
}

// @Tags OsInstallQueue
// @Summary 删除操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallQueue true "删除操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /queue/deleteOsInstallQueue [delete]
export const deleteOsInstallQueue = (data) => {
  return service({
    url: '/queue/deleteOsInstallQueue',
    method: 'delete',
    data
  })
}

// @Tags OsInstallQueue
// @Summary 批量删除操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /queue/deleteOsInstallQueue [delete]
export const deleteOsInstallQueueByIds = (data) => {
  return service({
    url: '/queue/deleteOsInstallQueueByIds',
    method: 'delete',
    data
  })
}

// @Tags OsInstallQueue
// @Summary 更新操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallQueue true "更新操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /queue/updateOsInstallQueue [put]
export const updateOsInstallQueue = (data) => {
  return service({
    url: '/queue/updateOsInstallQueue',
    method: 'put',
    data
  })
}

// @Tags OsInstallQueue
// @Summary 用id查询操作系统安装队列
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OsInstallQueue true "用id查询操作系统安装队列"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /queue/findOsInstallQueue [get]
export const findOsInstallQueue = (params) => {
  return service({
    url: '/queue/findOsInstallQueue',
    method: 'get',
    params
  })
}

// @Tags OsInstallQueue
// @Summary 分页获取操作系统安装队列列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取操作系统安装队列列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /queue/getOsInstallQueueList [get]
export const getOsInstallQueueList = (params) => {
  return service({
    url: '/queue/getOsInstallQueueList',
    method: 'get',
    params
  })
}
