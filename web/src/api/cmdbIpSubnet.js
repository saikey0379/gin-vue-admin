import service from '@/utils/request'

// @Tags CmdbIpSubnet
// @Summary 创建子网管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmdbIpSubnet true "创建子网管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmdbIpSubnet/createCmdbIpSubnet [post]
export const createCmdbIpSubnet = (data) => {
  return service({
    url: '/cmdbIpSubnet/createCmdbIpSubnet',
    method: 'post',
    data
  })
}

// @Tags CmdbIpSubnet
// @Summary 删除子网管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmdbIpSubnet true "删除子网管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmdbIpSubnet/deleteCmdbIpSubnet [delete]
export const deleteCmdbIpSubnet = (data) => {
  return service({
    url: '/cmdbIpSubnet/deleteCmdbIpSubnet',
    method: 'delete',
    data
  })
}

// @Tags CmdbIpSubnet
// @Summary 批量删除子网管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除子网管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmdbIpSubnet/deleteCmdbIpSubnet [delete]
export const deleteCmdbIpSubnetByIds = (data) => {
  return service({
    url: '/cmdbIpSubnet/deleteCmdbIpSubnetByIds',
    method: 'delete',
    data
  })
}

// @Tags CmdbIpSubnet
// @Summary 更新子网管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmdbIpSubnet true "更新子网管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmdbIpSubnet/updateCmdbIpSubnet [put]
export const updateCmdbIpSubnet = (data) => {
  return service({
    url: '/cmdbIpSubnet/updateCmdbIpSubnet',
    method: 'put',
    data
  })
}

// @Tags CmdbIpSubnet
// @Summary 用id查询子网管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmdbIpSubnet true "用id查询子网管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmdbIpSubnet/findCmdbIpSubnet [get]
export const findCmdbIpSubnet = (params) => {
  return service({
    url: '/cmdbIpSubnet/findCmdbIpSubnet',
    method: 'get',
    params
  })
}

// @Tags CmdbIpSubnet
// @Summary 分页获取子网管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取子网管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmdbIpSubnet/getCmdbIpSubnetList [get]
export const getCmdbIpSubnetList = (params) => {
  return service({
    url: '/cmdbIpSubnet/getCmdbIpSubnetList',
    method: 'get',
    params
  })
}

export const validateIpSubnet = (params) => {
  return service({
    url: '/cmdbIpSubnet/validateIpSubnet',
    method: 'get',
    params
  })
}