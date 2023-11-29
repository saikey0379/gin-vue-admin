import service from '@/utils/request'

// @Tags DeviceBareMetalHardware
// @Summary 创建裸金属设备硬件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceBareMetalHardware true "创建裸金属设备硬件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /deviceBareMetalHardware/createDeviceBareMetalHardware [post]
export const createDeviceBareMetalHardware = (data) => {
  return service({
    url: '/deviceBareMetalHardware/createDeviceBareMetalHardware',
    method: 'post',
    data
  })
}

// @Tags DeviceBareMetalHardware
// @Summary 删除裸金属设备硬件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceBareMetalHardware true "删除裸金属设备硬件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceBareMetalHardware/deleteDeviceBareMetalHardware [delete]
export const deleteDeviceBareMetalHardware = (data) => {
  return service({
    url: '/deviceBareMetalHardware/deleteDeviceBareMetalHardware',
    method: 'delete',
    data
  })
}

// @Tags DeviceBareMetalHardware
// @Summary 批量删除裸金属设备硬件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除裸金属设备硬件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceBareMetalHardware/deleteDeviceBareMetalHardware [delete]
export const deleteDeviceBareMetalHardwareByIds = (data) => {
  return service({
    url: '/deviceBareMetalHardware/deleteDeviceBareMetalHardwareByIds',
    method: 'delete',
    data
  })
}

// @Tags DeviceBareMetalHardware
// @Summary 更新裸金属设备硬件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceBareMetalHardware true "更新裸金属设备硬件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deviceBareMetalHardware/updateDeviceBareMetalHardware [put]
export const updateDeviceBareMetalHardware = (data) => {
  return service({
    url: '/deviceBareMetalHardware/updateDeviceBareMetalHardware',
    method: 'put',
    data
  })
}

// @Tags DeviceBareMetalHardware
// @Summary 用id查询裸金属设备硬件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DeviceBareMetalHardware true "用id查询裸金属设备硬件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deviceBareMetalHardware/findDeviceBareMetalHardware [get]
export const findDeviceBareMetalHardware = (params) => {
  return service({
    url: '/deviceBareMetalHardware/findDeviceBareMetalHardware',
    method: 'get',
    params
  })
}

// @Tags DeviceBareMetalHardware
// @Summary 分页获取裸金属设备硬件信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取裸金属设备硬件信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceBareMetalHardware/getDeviceBareMetalHardwareList [get]
export const getDeviceBareMetalHardwareList = (params) => {
  return service({
    url: '/deviceBareMetalHardware/getDeviceBareMetalHardwareList',
    method: 'get',
    params
  })
}
