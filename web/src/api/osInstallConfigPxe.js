import service from '@/utils/request'

// @Tags OsInstallConfigPxe
// @Summary 创建PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallConfigPxe true "创建PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /osInstallConfigPxe/createOsInstallConfigPxe [post]
export const createOsInstallConfigPxe = (data) => {
  return service({
    url: '/osInstallConfigPxe/createOsInstallConfigPxe',
    method: 'post',
    data
  })
}

// @Tags OsInstallConfigPxe
// @Summary 删除PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallConfigPxe true "删除PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /osInstallConfigPxe/deleteOsInstallConfigPxe [delete]
export const deleteOsInstallConfigPxe = (data) => {
  return service({
    url: '/osInstallConfigPxe/deleteOsInstallConfigPxe',
    method: 'delete',
    data
  })
}

// @Tags OsInstallConfigPxe
// @Summary 批量删除PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /osInstallConfigPxe/deleteOsInstallConfigPxe [delete]
export const deleteOsInstallConfigPxeByIds = (data) => {
  return service({
    url: '/osInstallConfigPxe/deleteOsInstallConfigPxeByIds',
    method: 'delete',
    data
  })
}

// @Tags OsInstallConfigPxe
// @Summary 更新PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallConfigPxe true "更新PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /osInstallConfigPxe/updateOsInstallConfigPxe [put]
export const updateOsInstallConfigPxe = (data) => {
  return service({
    url: '/osInstallConfigPxe/updateOsInstallConfigPxe',
    method: 'put',
    data
  })
}

// @Tags OsInstallConfigPxe
// @Summary 用id查询PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OsInstallConfigPxe true "用id查询PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /osInstallConfigPxe/findOsInstallConfigPxe [get]
export const findOsInstallConfigPxe = (params) => {
  return service({
    url: '/osInstallConfigPxe/findOsInstallConfigPxe',
    method: 'get',
    params
  })
}

// @Tags OsInstallConfigPxe
// @Summary 分页获取PXE配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PXE配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /osInstallConfigPxe/getOsInstallConfigPxeList [get]
export const getOsInstallConfigPxeList = (params) => {
  return service({
    url: '/osInstallConfigPxe/getOsInstallConfigPxeList',
    method: 'get',
    params
  })
}
