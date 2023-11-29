import service from '@/utils/request'

// @Tags SlbAccesslist
// @Summary 创建访问控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbAccesslist true "创建访问控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /slbAccesslist/createSlbAccesslist [post]
export const createSlbAccesslist = (data) => {
  return service({
    url: '/slbAccesslist/createSlbAccesslist',
    method: 'post',
    data
  })
}

// @Tags SlbAccesslist
// @Summary 删除访问控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbAccesslist true "删除访问控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbAccesslist/deleteSlbAccesslist [delete]
export const deleteSlbAccesslist = (data) => {
  return service({
    url: '/slbAccesslist/deleteSlbAccesslist',
    method: 'delete',
    data
  })
}

// @Tags SlbAccesslist
// @Summary 批量删除访问控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除访问控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbAccesslist/deleteSlbAccesslist [delete]
export const deleteSlbAccesslistByIds = (data) => {
  return service({
    url: '/slbAccesslist/deleteSlbAccesslistByIds',
    method: 'delete',
    data
  })
}

// @Tags SlbAccesslist
// @Summary 更新访问控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbAccesslist true "更新访问控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slbAccesslist/updateSlbAccesslist [put]
export const updateSlbAccesslist = (data) => {
  return service({
    url: '/slbAccesslist/updateSlbAccesslist',
    method: 'put',
    data
  })
}

// @Tags SlbAccesslist
// @Summary 用id查询访问控制
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SlbAccesslist true "用id查询访问控制"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slbAccesslist/findSlbAccesslist [get]
export const findSlbAccesslist = (params) => {
  return service({
    url: '/slbAccesslist/findSlbAccesslist',
    method: 'get',
    params
  })
}

// @Tags SlbAccesslist
// @Summary 分页获取访问控制列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取访问控制列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slbAccesslist/getSlbAccesslistList [get]
export const getSlbAccesslistList = (params) => {
  return service({
    url: '/slbAccesslist/getSlbAccesslistList',
    method: 'get',
    params
  })
}
