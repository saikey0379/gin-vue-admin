import service from '@/utils/request'

// @Tags DataCenter
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataCenter true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dc/createDataCenter [post]
export const createDataCenter = (data) => {
  return service({
    url: '/dc/createDataCenter',
    method: 'post',
    data
  })
}

// @Tags DataCenter
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataCenter true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dc/deleteDataCenter [delete]
export const deleteDataCenter = (data) => {
  return service({
    url: '/dc/deleteDataCenter',
    method: 'delete',
    data
  })
}

// @Tags DataCenter
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dc/deleteDataCenter [delete]
export const deleteDataCenterByIds = (data) => {
  return service({
    url: '/dc/deleteDataCenterByIds',
    method: 'delete',
    data
  })
}

// @Tags DataCenter
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataCenter true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dc/updateDataCenter [put]
export const updateDataCenter = (data) => {
  return service({
    url: '/dc/updateDataCenter',
    method: 'put',
    data
  })
}

// @Tags DataCenter
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DataCenter true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dc/findDataCenter [get]
export const findDataCenter = (params) => {
  return service({
    url: '/dc/findDataCenter',
    method: 'get',
    params
  })
}

// @Tags DataCenter
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dc/getDataCenterList [get]
export const getDataCenterList = (params) => {
  return service({
    url: '/dc/getDataCenterList',
    method: 'get',
    params
  })
}
