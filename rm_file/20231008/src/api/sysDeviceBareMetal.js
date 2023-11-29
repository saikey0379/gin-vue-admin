import service from '@/utils/request'

// @Tags BareMetal
// @Summary 创建裸金属管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BareMetal true "创建裸金属管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /DeviceBareMetal/createBareMetal [post]
export const createBareMetal = (data) => {
  return service({
    url: '/DeviceBareMetal/createBareMetal',
    method: 'post',
    data
  })
}

// @Tags BareMetal
// @Summary 删除裸金属管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BareMetal true "删除裸金属管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DeviceBareMetal/deleteBareMetal [delete]
export const deleteBareMetal = (data) => {
  return service({
    url: '/DeviceBareMetal/deleteBareMetal',
    method: 'delete',
    data
  })
}

// @Tags BareMetal
// @Summary 批量删除裸金属管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除裸金属管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DeviceBareMetal/deleteBareMetal [delete]
export const deleteBareMetalByIds = (data) => {
  return service({
    url: '/DeviceBareMetal/deleteBareMetalByIds',
    method: 'delete',
    data
  })
}

// @Tags BareMetal
// @Summary 更新裸金属管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BareMetal true "更新裸金属管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DeviceBareMetal/updateBareMetal [put]
export const updateBareMetal = (data) => {
  return service({
    url: '/DeviceBareMetal/updateBareMetal',
    method: 'put',
    data
  })
}

// @Tags BareMetal
// @Summary 用id查询裸金属管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BareMetal true "用id查询裸金属管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DeviceBareMetal/findBareMetal [get]
export const findBareMetal = (params) => {
  return service({
    url: '/DeviceBareMetal/findBareMetal',
    method: 'get',
    params
  })
}

// @Tags BareMetal
// @Summary 分页获取裸金属管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取裸金属管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DeviceBareMetal/getBareMetalList [get]
export const getBareMetalList = (params) => {
  return service({
    url: '/DeviceBareMetal/getBareMetalList',
    method: 'get',
    params
  })
}
