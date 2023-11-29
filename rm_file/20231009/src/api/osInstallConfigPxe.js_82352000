import service from '@/utils/request'

// @Tags ConfigPxe
// @Summary 创建PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigPxe true "创建PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pxe/createConfigPxe [post]
export const createConfigPxe = (data) => {
  return service({
    url: '/pxe/createConfigPxe',
    method: 'post',
    data
  })
}

// @Tags ConfigPxe
// @Summary 删除PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigPxe true "删除PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pxe/deleteConfigPxe [delete]
export const deleteConfigPxe = (data) => {
  return service({
    url: '/pxe/deleteConfigPxe',
    method: 'delete',
    data
  })
}

// @Tags ConfigPxe
// @Summary 批量删除PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pxe/deleteConfigPxe [delete]
export const deleteConfigPxeByIds = (data) => {
  return service({
    url: '/pxe/deleteConfigPxeByIds',
    method: 'delete',
    data
  })
}

// @Tags ConfigPxe
// @Summary 更新PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConfigPxe true "更新PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pxe/updateConfigPxe [put]
export const updateConfigPxe = (data) => {
  return service({
    url: '/pxe/updateConfigPxe',
    method: 'put',
    data
  })
}

// @Tags ConfigPxe
// @Summary 用id查询PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ConfigPxe true "用id查询PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pxe/findConfigPxe [get]
export const findConfigPxe = (params) => {
  return service({
    url: '/pxe/findConfigPxe',
    method: 'get',
    params
  })
}

// @Tags ConfigPxe
// @Summary 分页获取PXE配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PXE配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pxe/getConfigPxeList [get]
export const getConfigPxeList = (params) => {
  return service({
    url: '/pxe/getConfigPxeList',
    method: 'get',
    params
  })
}
