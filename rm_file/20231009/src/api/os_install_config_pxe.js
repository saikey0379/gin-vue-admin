import service from '@/utils/request'

// @Tags Pxe
// @Summary 创建PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pxe true "创建PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PxeConfig/createPxe [post]
export const createPxe = (data) => {
  return service({
    url: '/PxeConfig/createPxe',
    method: 'post',
    data
  })
}

// @Tags Pxe
// @Summary 删除PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pxe true "删除PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PxeConfig/deletePxe [delete]
export const deletePxe = (data) => {
  return service({
    url: '/PxeConfig/deletePxe',
    method: 'delete',
    data
  })
}

// @Tags Pxe
// @Summary 批量删除PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PxeConfig/deletePxe [delete]
export const deletePxeByIds = (data) => {
  return service({
    url: '/PxeConfig/deletePxeByIds',
    method: 'delete',
    data
  })
}

// @Tags Pxe
// @Summary 更新PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pxe true "更新PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PxeConfig/updatePxe [put]
export const updatePxe = (data) => {
  return service({
    url: '/PxeConfig/updatePxe',
    method: 'put',
    data
  })
}

// @Tags Pxe
// @Summary 用id查询PXE配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Pxe true "用id查询PXE配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PxeConfig/findPxe [get]
export const findPxe = (params) => {
  return service({
    url: '/PxeConfig/findPxe',
    method: 'get',
    params
  })
}

// @Tags Pxe
// @Summary 分页获取PXE配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PXE配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PxeConfig/getPxeList [get]
export const getPxeList = (params) => {
  return service({
    url: '/PxeConfig/getPxeList',
    method: 'get',
    params
  })
}
