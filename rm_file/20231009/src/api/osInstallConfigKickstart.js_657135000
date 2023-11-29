import service from '@/utils/request'

// @Tags OsInstallConfigKickstart
// @Summary 创建Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallConfigKickstart true "创建Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ks/createOsInstallConfigKickstart [post]
export const createOsInstallConfigKickstart = (data) => {
  return service({
    url: '/ks/createOsInstallConfigKickstart',
    method: 'post',
    data
  })
}

// @Tags OsInstallConfigKickstart
// @Summary 删除Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallConfigKickstart true "删除Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ks/deleteOsInstallConfigKickstart [delete]
export const deleteOsInstallConfigKickstart = (data) => {
  return service({
    url: '/ks/deleteOsInstallConfigKickstart',
    method: 'delete',
    data
  })
}

// @Tags OsInstallConfigKickstart
// @Summary 批量删除Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ks/deleteOsInstallConfigKickstart [delete]
export const deleteOsInstallConfigKickstartByIds = (data) => {
  return service({
    url: '/ks/deleteOsInstallConfigKickstartByIds',
    method: 'delete',
    data
  })
}

// @Tags OsInstallConfigKickstart
// @Summary 更新Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OsInstallConfigKickstart true "更新Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ks/updateOsInstallConfigKickstart [put]
export const updateOsInstallConfigKickstart = (data) => {
  return service({
    url: '/ks/updateOsInstallConfigKickstart',
    method: 'put',
    data
  })
}

// @Tags OsInstallConfigKickstart
// @Summary 用id查询Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OsInstallConfigKickstart true "用id查询Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ks/findOsInstallConfigKickstart [get]
export const findOsInstallConfigKickstart = (params) => {
  return service({
    url: '/ks/findOsInstallConfigKickstart',
    method: 'get',
    params
  })
}

// @Tags OsInstallConfigKickstart
// @Summary 分页获取Kickstart配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Kickstart配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ks/getOsInstallConfigKickstartList [get]
export const getOsInstallConfigKickstartList = (params) => {
  return service({
    url: '/ks/getOsInstallConfigKickstartList',
    method: 'get',
    params
  })
}
