import service from '@/utils/request'

// @Tags ServerDiscovery
// @Summary 创建ServerDiscovery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServerDiscovery true "创建ServerDiscovery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /serverDiscovery/createServerDiscovery [post]
export const createServerDiscovery = (data) => {
  return service({
    url: '/serverDiscovery/createServerDiscovery',
    method: 'post',
    data
  })
}

// @Tags ServerDiscovery
// @Summary 删除ServerDiscovery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServerDiscovery true "删除ServerDiscovery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serverDiscovery/deleteServerDiscovery [delete]
export const deleteServerDiscovery = (data) => {
  return service({
    url: '/serverDiscovery/deleteServerDiscovery',
    method: 'delete',
    data
  })
}

// @Tags ServerDiscovery
// @Summary 批量删除ServerDiscovery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ServerDiscovery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serverDiscovery/deleteServerDiscovery [delete]
export const deleteServerDiscoveryByIds = (data) => {
  return service({
    url: '/serverDiscovery/deleteServerDiscoveryByIds',
    method: 'delete',
    data
  })
}

// @Tags ServerDiscovery
// @Summary 更新ServerDiscovery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServerDiscovery true "更新ServerDiscovery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /serverDiscovery/updateServerDiscovery [put]
export const updateServerDiscovery = (data) => {
  return service({
    url: '/serverDiscovery/updateServerDiscovery',
    method: 'put',
    data
  })
}

// @Tags ServerDiscovery
// @Summary 用id查询ServerDiscovery
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ServerDiscovery true "用id查询ServerDiscovery"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /serverDiscovery/findServerDiscovery [get]
export const findServerDiscovery = (params) => {
  return service({
    url: '/serverDiscovery/findServerDiscovery',
    method: 'get',
    params
  })
}

// @Tags ServerDiscovery
// @Summary 分页获取ServerDiscovery列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ServerDiscovery列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serverDiscovery/getServerDiscoveryList [get]
export const getServerDiscoveryList = (params) => {
  return service({
    url: '/serverDiscovery/getServerDiscoveryList',
    method: 'get',
    params
  })
}

export const entryServerDiscoveryByIds = (data) => {
  return service({
    url: '/serverDiscovery/entryServerDiscoveryByIds',
    method: 'post',
    data
  })
}