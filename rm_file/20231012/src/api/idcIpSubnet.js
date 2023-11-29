import service from '@/utils/request'

// @Tags IdcIpSubnet
// @Summary 创建数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpSubnet true "创建数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpSubnet/createIdcIpSubnet [post]
export const createIdcIpSubnet = (data) => {
  return service({
    url: '/idcIpSubnet/createIdcIpSubnet',
    method: 'post',
    data
  })
}

// @Tags IdcIpSubnet
// @Summary 删除数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpSubnet true "删除数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpSubnet/deleteIdcIpSubnet [delete]
export const deleteIdcIpSubnet = (data) => {
  return service({
    url: '/idcIpSubnet/deleteIdcIpSubnet',
    method: 'delete',
    data
  })
}

// @Tags IdcIpSubnet
// @Summary 批量删除数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpSubnet/deleteIdcIpSubnet [delete]
export const deleteIdcIpSubnetByIds = (data) => {
  return service({
    url: '/idcIpSubnet/deleteIdcIpSubnetByIds',
    method: 'delete',
    data
  })
}

// @Tags IdcIpSubnet
// @Summary 更新数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpSubnet true "更新数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpSubnet/updateIdcIpSubnet [put]
export const updateIdcIpSubnet = (data) => {
  return service({
    url: '/idcIpSubnet/updateIdcIpSubnet',
    method: 'put',
    data
  })
}

// @Tags IdcIpSubnet
// @Summary 用id查询数据中心子网
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdcIpSubnet true "用id查询数据中心子网"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpSubnet/findIdcIpSubnet [get]
export const findIdcIpSubnet = (params) => {
  return service({
    url: '/idcIpSubnet/findIdcIpSubnet',
    method: 'get',
    params
  })
}

// @Tags IdcIpSubnet
// @Summary 分页获取数据中心子网列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心子网列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpSubnet/getIdcIpSubnetList [get]
export const getIdcIpSubnetList = (params) => {
  return service({
    url: '/idcIpSubnet/getIdcIpSubnetList',
    method: 'get',
    params
  })
}
