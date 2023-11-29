import service from '@/utils/request'

// @Tags DataCenterInfo
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataCenterInfo true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dc/createDataCenterInfo [post]
export const createDataCenterInfo = (data) => {
  return service({
    url: '/dc/createDataCenterInfo',
    method: 'post',
    data
  })
}

// @Tags DataCenterInfo
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataCenterInfo true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dc/deleteDataCenterInfo [delete]
export const deleteDataCenterInfo = (data) => {
  return service({
    url: '/dc/deleteDataCenterInfo',
    method: 'delete',
    data
  })
}

// @Tags DataCenterInfo
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dc/deleteDataCenterInfo [delete]
export const deleteDataCenterInfoByIds = (data) => {
  return service({
    url: '/dc/deleteDataCenterInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags DataCenterInfo
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataCenterInfo true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dc/updateDataCenterInfo [put]
export const updateDataCenterInfo = (data) => {
  return service({
    url: '/dc/updateDataCenterInfo',
    method: 'put',
    data
  })
}

// @Tags DataCenterInfo
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DataCenterInfo true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dc/findDataCenterInfo [get]
export const findDataCenterInfo = (params) => {
  return service({
    url: '/dc/findDataCenterInfo',
    method: 'get',
    params
  })
}

// @Tags DataCenterInfo
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dc/getDataCenterInfoList [get]
export const getDataCenterInfoList = (params) => {
  return service({
    url: '/dc/getDataCenterInfoList',
    method: 'get',
    params
  })
}
