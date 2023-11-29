import service from '@/utils/request'

// @Tags DeviceBareMetal
// @Summary 创建裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceBareMetal true "创建裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /deviceBareMetal/createDeviceBareMetal [post]
export const createDeviceBareMetal = (data) => {
  return service({
    url: '/deviceBareMetal/createDeviceBareMetal',
    method: 'post',
    data
  })
}

// @Tags DeviceBareMetal
// @Summary 删除裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceBareMetal true "删除裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceBareMetal/deleteDeviceBareMetal [delete]
export const deleteDeviceBareMetal = (data) => {
  return service({
    url: '/deviceBareMetal/deleteDeviceBareMetal',
    method: 'delete',
    data
  })
}

// @Tags DeviceBareMetal
// @Summary 批量删除裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deviceBareMetal/deleteDeviceBareMetal [delete]
export const deleteDeviceBareMetalByIds = (data) => {
  return service({
    url: '/deviceBareMetal/deleteDeviceBareMetalByIds',
    method: 'delete',
    data
  })
}

// @Tags DeviceBareMetal
// @Summary 更新裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeviceBareMetal true "更新裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deviceBareMetal/updateDeviceBareMetal [put]
export const updateDeviceBareMetal = (data) => {
  return service({
    url: '/deviceBareMetal/updateDeviceBareMetal',
    method: 'put',
    data
  })
}

// @Tags DeviceBareMetal
// @Summary 用id查询裸金属设备
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DeviceBareMetal true "用id查询裸金属设备"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deviceBareMetal/findDeviceBareMetal [get]
export const findDeviceBareMetal = (params) => {
  return service({
    url: '/deviceBareMetal/findDeviceBareMetal',
    method: 'get',
    params
  })
}

// @Tags DeviceBareMetal
// @Summary 分页获取裸金属设备列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取裸金属设备列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deviceBareMetal/getDeviceBareMetalList [get]
export const getDeviceBareMetalList = (params) => {
  return service({
    url: '/deviceBareMetal/getDeviceBareMetalList',
    method: 'get',
    params
  })
}
