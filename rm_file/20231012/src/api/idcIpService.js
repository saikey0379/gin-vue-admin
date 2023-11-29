import service from '@/utils/request'

// @Tags IdcIpService
// @Summary 创建数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpService true "创建数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpService/createIdcIpService [post]
export const createIdcIpService = (data) => {
  return service({
    url: '/idcIpService/createIdcIpService',
    method: 'post',
    data
  })
}

// @Tags IdcIpService
// @Summary 删除数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpService true "删除数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpService/deleteIdcIpService [delete]
export const deleteIdcIpService = (data) => {
  return service({
    url: '/idcIpService/deleteIdcIpService',
    method: 'delete',
    data
  })
}

// @Tags IdcIpService
// @Summary 批量删除数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpService/deleteIdcIpService [delete]
export const deleteIdcIpServiceByIds = (data) => {
  return service({
    url: '/idcIpService/deleteIdcIpServiceByIds',
    method: 'delete',
    data
  })
}

// @Tags IdcIpService
// @Summary 更新数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpService true "更新数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpService/updateIdcIpService [put]
export const updateIdcIpService = (data) => {
  return service({
    url: '/idcIpService/updateIdcIpService',
    method: 'put',
    data
  })
}

// @Tags IdcIpService
// @Summary 用id查询数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdcIpService true "用id查询数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpService/findIdcIpService [get]
export const findIdcIpService = (params) => {
  return service({
    url: '/idcIpService/findIdcIpService',
    method: 'get',
    params
  })
}

// @Tags IdcIpService
// @Summary 分页获取数据中心业务网段列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心业务网段列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpService/getIdcIpServiceList [get]
export const getIdcIpServiceList = (params) => {
  return service({
    url: '/idcIpService/getIdcIpServiceList',
    method: 'get',
    params
  })
}

export const validateIpService = (params) => {
  return service({
    url: '/idcIpService/validateIpService',
    method: 'get',
    params
  })
}