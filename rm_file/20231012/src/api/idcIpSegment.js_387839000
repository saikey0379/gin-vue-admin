import service from '@/utils/request'

// @Tags IdcIpSegment
// @Summary 创建数据中心IP网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpSegment true "创建数据中心IP网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpSegment/createIdcIpSegment [post]
export const createIdcIpSegment = (data) => {
  return service({
    url: '/idcIpSegment/createIdcIpSegment',
    method: 'post',
    data
  })
}

// @Tags IdcIpSegment
// @Summary 删除数据中心IP网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpSegment true "删除数据中心IP网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpSegment/deleteIdcIpSegment [delete]
export const deleteIdcIpSegment = (data) => {
  return service({
    url: '/idcIpSegment/deleteIdcIpSegment',
    method: 'delete',
    data
  })
}

// @Tags IdcIpSegment
// @Summary 批量删除数据中心IP网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心IP网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpSegment/deleteIdcIpSegment [delete]
export const deleteIdcIpSegmentByIds = (data) => {
  return service({
    url: '/idcIpSegment/deleteIdcIpSegmentByIds',
    method: 'delete',
    data
  })
}

// @Tags IdcIpSegment
// @Summary 更新数据中心IP网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpSegment true "更新数据中心IP网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpSegment/updateIdcIpSegment [put]
export const updateIdcIpSegment = (data) => {
  return service({
    url: '/idcIpSegment/updateIdcIpSegment',
    method: 'put',
    data
  })
}

// @Tags IdcIpSegment
// @Summary 用id查询数据中心IP网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdcIpSegment true "用id查询数据中心IP网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpSegment/findIdcIpSegment [get]
export const findIdcIpSegment = (params) => {
  return service({
    url: '/idcIpSegment/findIdcIpSegment',
    method: 'get',
    params
  })
}

// @Tags IdcIpSegment
// @Summary 分页获取数据中心IP网段列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心IP网段列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpSegment/getIdcIpSegmentList [get]
export const getIdcIpSegmentList = (params) => {
  return service({
    url: '/idcIpSegment/getIdcIpSegmentList',
    method: 'get',
    params
  })
}
