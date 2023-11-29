import service from '@/utils/request'

// @Tags OsInstallLog
// @Summary 创建操作系统安装日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallLog true "创建操作系统安装日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /osInstallLog/createOsInstallLog [post]
export const createOsInstallLog = (data) => {
  return service({
    url: '/osInstallLog/createOsInstallLog',
    method: 'post',
    data
  })
}

// @Tags OsInstallLog
// @Summary 删除操作系统安装日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallLog true "删除操作系统安装日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /osInstallLog/deleteOsInstallLog [delete]
export const deleteOsInstallLog = (data) => {
  return service({
    url: '/osInstallLog/deleteOsInstallLog',
    method: 'delete',
    data
  })
}

// @Tags OsInstallLog
// @Summary 批量删除操作系统安装日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除操作系统安装日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /osInstallLog/deleteOsInstallLog [delete]
export const deleteOsInstallLogByIds = (data) => {
  return service({
    url: '/osInstallLog/deleteOsInstallLogByIds',
    method: 'delete',
    data
  })
}

// @Tags OsInstallLog
// @Summary 更新操作系统安装日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallLog true "更新操作系统安装日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /osInstallLog/updateOsInstallLog [put]
export const updateOsInstallLog = (data) => {
  return service({
    url: '/osInstallLog/updateOsInstallLog',
    method: 'put',
    data
  })
}

// @Tags OsInstallLog
// @Summary 用id查询操作系统安装日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OsInstallLog true "用id查询操作系统安装日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /osInstallLog/findOsInstallLog [get]
export const findOsInstallLog = (params) => {
  return service({
    url: '/osInstallLog/findOsInstallLog',
    method: 'get',
    params
  })
}

// @Tags OsInstallLog
// @Summary 分页获取操作系统安装日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取操作系统安装日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /osInstallLog/getOsInstallLogList [get]
export const getOsInstallLogList = (params) => {
  return service({
    url: '/osInstallLog/getOsInstallLogList',
    method: 'get',
    params
  })
}
