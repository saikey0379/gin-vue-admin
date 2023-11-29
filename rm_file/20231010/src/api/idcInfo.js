import service from '@/utils/request'

// @Tags IdcInfo
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcInfo true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcInfo/createIdcInfo [post]
export const createIdcInfo = (data) => {
  return service({
    url: '/idcInfo/createIdcInfo',
    method: 'post',
    data
  })
}

// @Tags IdcInfo
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcInfo true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcInfo/deleteIdcInfo [delete]
export const deleteIdcInfo = (data) => {
  return service({
    url: '/idcInfo/deleteIdcInfo',
    method: 'delete',
    data
  })
}

// @Tags IdcInfo
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcInfo/deleteIdcInfo [delete]
export const deleteIdcInfoByIds = (data) => {
  return service({
    url: '/idcInfo/deleteIdcInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags IdcInfo
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcInfo true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcInfo/updateIdcInfo [put]
export const updateIdcInfo = (data) => {
  return service({
    url: '/idcInfo/updateIdcInfo',
    method: 'put',
    data
  })
}

// @Tags IdcInfo
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdcInfo true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcInfo/findIdcInfo [get]
export const findIdcInfo = (params) => {
  return service({
    url: '/idcInfo/findIdcInfo',
    method: 'get',
    params
  })
}

// @Tags IdcInfo
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcInfo/getIdcInfoList [get]
export const getIdcInfoList = (params) => {
  return service({
    url: '/idcInfo/getIdcInfoList',
    method: 'get',
    params
  })
}
