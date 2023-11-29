import service from '@/utils/request'

// @Tags FileCommadBlacklist
// @Summary 创建命令黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileCommadBlacklist true "创建命令黑名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileCommandBlacklist/createFileCommadBlacklist [post]
export const createFileCommadBlacklist = (data) => {
  return service({
    url: '/fileCommandBlacklist/createFileCommadBlacklist',
    method: 'post',
    data
  })
}

// @Tags FileCommadBlacklist
// @Summary 删除命令黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileCommadBlacklist true "删除命令黑名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileCommandBlacklist/deleteFileCommadBlacklist [delete]
export const deleteFileCommadBlacklist = (data) => {
  return service({
    url: '/fileCommandBlacklist/deleteFileCommadBlacklist',
    method: 'delete',
    data
  })
}

// @Tags FileCommadBlacklist
// @Summary 批量删除命令黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除命令黑名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileCommandBlacklist/deleteFileCommadBlacklist [delete]
export const deleteFileCommadBlacklistByIds = (data) => {
  return service({
    url: '/fileCommandBlacklist/deleteFileCommadBlacklistByIds',
    method: 'delete',
    data
  })
}

// @Tags FileCommadBlacklist
// @Summary 更新命令黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileCommadBlacklist true "更新命令黑名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileCommandBlacklist/updateFileCommadBlacklist [put]
export const updateFileCommadBlacklist = (data) => {
  return service({
    url: '/fileCommandBlacklist/updateFileCommadBlacklist',
    method: 'put',
    data
  })
}

// @Tags FileCommadBlacklist
// @Summary 用id查询命令黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FileCommadBlacklist true "用id查询命令黑名单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileCommandBlacklist/findFileCommadBlacklist [get]
export const findFileCommadBlacklist = (params) => {
  return service({
    url: '/fileCommandBlacklist/findFileCommadBlacklist',
    method: 'get',
    params
  })
}

// @Tags FileCommadBlacklist
// @Summary 分页获取命令黑名单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取命令黑名单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileCommandBlacklist/getFileCommadBlacklistList [get]
export const getFileCommadBlacklistList = (params) => {
  return service({
    url: '/fileCommandBlacklist/getFileCommadBlacklistList',
    method: 'get',
    params
  })
}
