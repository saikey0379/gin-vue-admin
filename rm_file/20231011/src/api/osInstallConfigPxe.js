import service from '@/utils/request'

// @Tags OsInstallConfigPxe
// @Summary 创建PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallConfigPxe true "创建PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pxe/createOsInstallConfigPxe [post]
export const createOsInstallConfigPxe = (data) => {
  return service({
    url: '/pxe/createOsInstallConfigPxe',
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
// @Router /pxe/deleteOsInstallConfigPxe [delete]
export const deleteOsInstallConfigPxe = (data) => {
  return service({
    url: '/pxe/deleteOsInstallConfigPxe',
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
// @Router /pxe/deleteOsInstallConfigPxe [delete]
export const deleteOsInstallConfigPxeByIds = (data) => {
  return service({
    url: '/pxe/deleteOsInstallConfigPxeByIds',
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
// @Router /pxe/updateOsInstallConfigPxe [put]
export const updateOsInstallConfigPxe = (data) => {
  return service({
    url: '/pxe/updateOsInstallConfigPxe',
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
// @Router /pxe/findOsInstallConfigPxe [get]
export const findOsInstallConfigPxe = (params) => {
  return service({
    url: '/pxe/findOsInstallConfigPxe',
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
// @Router /pxe/getOsInstallConfigPxeList [get]
export const getOsInstallConfigPxeList = (params) => {
  return service({
    url: '/pxe/getOsInstallConfigPxeList',
    method: 'get',
    params
  })
}
