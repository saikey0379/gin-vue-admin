import service from '@/utils/request'

// @Tags DeviceServerDiscovery
// @Summary 创建Server发现
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceServerDiscovery true "创建Server发现"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /deviceServerDiscovery/createDeviceServerDiscovery [post]
export const createDeviceServerDiscovery = (data) => {
  return service({
    url: '/deviceServerDiscovery/createDeviceServerDiscovery',
    method: 'post',
    data
  })
}

// @Tags DeviceServerDiscovery
// @Summary 删除Server发现
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceServerDiscovery true "删除Server发现"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceServerDiscovery/deleteDeviceServerDiscovery [delete]
export const deleteDeviceServerDiscovery = (data) => {
  return service({
    url: '/deviceServerDiscovery/deleteDeviceServerDiscovery',
    method: 'delete',
    data
  })
}

// @Tags DeviceServerDiscovery
// @Summary 批量删除Server发现
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Server发现"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceServerDiscovery/deleteDeviceServerDiscovery [delete]
export const deleteDeviceServerDiscoveryByIds = (data) => {
  return service({
    url: '/deviceServerDiscovery/deleteDeviceServerDiscoveryByIds',
    method: 'delete',
    data
  })
}

// @Tags DeviceServerDiscovery
// @Summary 更新Server发现
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceServerDiscovery true "更新Server发现"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deviceServerDiscovery/updateDeviceServerDiscovery [put]
export const updateDeviceServerDiscovery = (data) => {
  return service({
    url: '/deviceServerDiscovery/updateDeviceServerDiscovery',
    method: 'put',
    data
  })
}

// @Tags DeviceServerDiscovery
// @Summary 用id查询Server发现
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DeviceServerDiscovery true "用id查询Server发现"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deviceServerDiscovery/findDeviceServerDiscovery [get]
export const findDeviceServerDiscovery = (params) => {
  return service({
    url: '/deviceServerDiscovery/findDeviceServerDiscovery',
    method: 'get',
    params
  })
}

// @Tags DeviceServerDiscovery
// @Summary 分页获取Server发现列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Server发现列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceServerDiscovery/getDeviceServerDiscoveryList [get]
export const getDeviceServerDiscoveryList = (params) => {
  return service({
    url: '/deviceServerDiscovery/getDeviceServerDiscoveryList',
    method: 'get',
    params
  })
}
