import service from '@/utils/request'

// @Tags ConfigKickstart
// @Summary 创建Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigKickstart true "创建Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ks/createConfigKickstart [post]
export const createConfigKickstart = (data) => {
  return service({
    url: '/ks/createConfigKickstart',
    method: 'post',
    data
  })
}

// @Tags ConfigKickstart
// @Summary 删除Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigKickstart true "删除Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ks/deleteConfigKickstart [delete]
export const deleteConfigKickstart = (data) => {
  return service({
    url: '/ks/deleteConfigKickstart',
    method: 'delete',
    data
  })
}

// @Tags ConfigKickstart
// @Summary 批量删除Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ks/deleteConfigKickstart [delete]
export const deleteConfigKickstartByIds = (data) => {
  return service({
    url: '/ks/deleteConfigKickstartByIds',
    method: 'delete',
    data
  })
}

// @Tags ConfigKickstart
// @Summary 更新Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigKickstart true "更新Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ks/updateConfigKickstart [put]
export const updateConfigKickstart = (data) => {
  return service({
    url: '/ks/updateConfigKickstart',
    method: 'put',
    data
  })
}

// @Tags ConfigKickstart
// @Summary 用id查询Kickstart配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ConfigKickstart true "用id查询Kickstart配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ks/findConfigKickstart [get]
export const findConfigKickstart = (params) => {
  return service({
    url: '/ks/findConfigKickstart',
    method: 'get',
    params
  })
}

// @Tags ConfigKickstart
// @Summary 分页获取Kickstart配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Kickstart配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ks/getConfigKickstartList [get]
export const getConfigKickstartList = (params) => {
  return service({
    url: '/ks/getConfigKickstartList',
    method: 'get',
    params
  })
}
