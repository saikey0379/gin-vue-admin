import service from '@/utils/request'

// @Tags CmdbIpSegment
// @Summary 创建网段管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmdbIpSegment true "创建网段管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmdbIpSegment/createCmdbIpSegment [post]
export const createCmdbIpSegment = (data) => {
  return service({
    url: '/cmdbIpSegment/createCmdbIpSegment',
    method: 'post',
    data
  })
}

// @Tags CmdbIpSegment
// @Summary 删除网段管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmdbIpSegment true "删除网段管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmdbIpSegment/deleteCmdbIpSegment [delete]
export const deleteCmdbIpSegment = (data) => {
  return service({
    url: '/cmdbIpSegment/deleteCmdbIpSegment',
    method: 'delete',
    data
  })
}

// @Tags CmdbIpSegment
// @Summary 批量删除网段管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除网段管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmdbIpSegment/deleteCmdbIpSegment [delete]
export const deleteCmdbIpSegmentByIds = (data) => {
  return service({
    url: '/cmdbIpSegment/deleteCmdbIpSegmentByIds',
    method: 'delete',
    data
  })
}

// @Tags CmdbIpSegment
// @Summary 更新网段管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmdbIpSegment true "更新网段管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmdbIpSegment/updateCmdbIpSegment [put]
export const updateCmdbIpSegment = (data) => {
  return service({
    url: '/cmdbIpSegment/updateCmdbIpSegment',
    method: 'put',
    data
  })
}

// @Tags CmdbIpSegment
// @Summary 用id查询网段管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmdbIpSegment true "用id查询网段管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmdbIpSegment/findCmdbIpSegment [get]
export const findCmdbIpSegment = (params) => {
  return service({
    url: '/cmdbIpSegment/findCmdbIpSegment',
    method: 'get',
    params
  })
}

// @Tags CmdbIpSegment
// @Summary 分页获取网段管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取网段管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmdbIpSegment/getCmdbIpSegmentList [get]
export const getCmdbIpSegmentList = (params) => {
  return service({
    url: '/cmdbIpSegment/getCmdbIpSegmentList',
    method: 'get',
    params
  })
}
