import service from '@/utils/request'

// @Tags DataCenterIpService
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataCenterIpService true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ipService/createDataCenterIpService [post]
export const createDataCenterIpService = (data) => {
  return service({
    url: '/ipService/createDataCenterIpService',
    method: 'post',
    data
  })
}

// @Tags DataCenterIpService
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataCenterIpService true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ipService/deleteDataCenterIpService [delete]
export const deleteDataCenterIpService = (data) => {
  return service({
    url: '/ipService/deleteDataCenterIpService',
    method: 'delete',
    data
  })
}

// @Tags DataCenterIpService
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ipService/deleteDataCenterIpService [delete]
export const deleteDataCenterIpServiceByIds = (data) => {
  return service({
    url: '/ipService/deleteDataCenterIpServiceByIds',
    method: 'delete',
    data
  })
}

// @Tags DataCenterIpService
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataCenterIpService true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ipService/updateDataCenterIpService [put]
export const updateDataCenterIpService = (data) => {
  return service({
    url: '/ipService/updateDataCenterIpService',
    method: 'put',
    data
  })
}

// @Tags DataCenterIpService
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DataCenterIpService true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ipService/findDataCenterIpService [get]
export const findDataCenterIpService = (params) => {
  return service({
    url: '/ipService/findDataCenterIpService',
    method: 'get',
    params
  })
}

// @Tags DataCenterIpService
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ipService/getDataCenterIpServiceList [get]
export const getDataCenterIpServiceList = (params) => {
  return service({
    url: '/ipService/getDataCenterIpServiceList',
    method: 'get',
    params
  })
}
