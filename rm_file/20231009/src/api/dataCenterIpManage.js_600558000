import service from '@/utils/request'

// @Tags DataCenterIpManage
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataCenterIpManage true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ipManage/createDataCenterIpManage [post]
export const createDataCenterIpManage = (data) => {
  return service({
    url: '/ipManage/createDataCenterIpManage',
    method: 'post',
    data
  })
}

// @Tags DataCenterIpManage
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataCenterIpManage true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ipManage/deleteDataCenterIpManage [delete]
export const deleteDataCenterIpManage = (data) => {
  return service({
    url: '/ipManage/deleteDataCenterIpManage',
    method: 'delete',
    data
  })
}

// @Tags DataCenterIpManage
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ipManage/deleteDataCenterIpManage [delete]
export const deleteDataCenterIpManageByIds = (data) => {
  return service({
    url: '/ipManage/deleteDataCenterIpManageByIds',
    method: 'delete',
    data
  })
}

// @Tags DataCenterIpManage
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataCenterIpManage true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ipManage/updateDataCenterIpManage [put]
export const updateDataCenterIpManage = (data) => {
  return service({
    url: '/ipManage/updateDataCenterIpManage',
    method: 'put',
    data
  })
}

// @Tags DataCenterIpManage
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DataCenterIpManage true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ipManage/findDataCenterIpManage [get]
export const findDataCenterIpManage = (params) => {
  return service({
    url: '/ipManage/findDataCenterIpManage',
    method: 'get',
    params
  })
}

// @Tags DataCenterIpManage
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ipManage/getDataCenterIpManageList [get]
export const getDataCenterIpManageList = (params) => {
  return service({
    url: '/ipManage/getDataCenterIpManageList',
    method: 'get',
    params
  })
}
