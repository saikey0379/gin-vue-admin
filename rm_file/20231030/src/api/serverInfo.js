import service from '@/utils/request'

// @Tags ServerInfo
// @Summary 创建ServerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServerInfo true "创建ServerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /serverInfo/createServerInfo [post]
export const createServerInfo = (data) => {
  return service({
    url: '/serverInfo/createServerInfo',
    method: 'post',
    data
  })
}

// @Tags ServerInfo
// @Summary 删除ServerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServerInfo true "删除ServerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serverInfo/deleteServerInfo [delete]
export const deleteServerInfo = (data) => {
  return service({
    url: '/serverInfo/deleteServerInfo',
    method: 'delete',
    data
  })
}

// @Tags ServerInfo
// @Summary 批量删除ServerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ServerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serverInfo/deleteServerInfo [delete]
export const deleteServerInfoByIds = (data) => {
  return service({
    url: '/serverInfo/deleteServerInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags ServerInfo
// @Summary 更新ServerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServerInfo true "更新ServerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /serverInfo/updateServerInfo [put]
export const updateServerInfo = (data) => {
  return service({
    url: '/serverInfo/updateServerInfo',
    method: 'put',
    data
  })
}

// @Tags ServerInfo
// @Summary 用id查询ServerInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ServerInfo true "用id查询ServerInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /serverInfo/findServerInfo [get]
export const findServerInfo = (params) => {
  return service({
    url: '/serverInfo/findServerInfo',
    method: 'get',
    params
  })
}

// @Tags ServerInfo
// @Summary 分页获取ServerInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ServerInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serverInfo/getServerInfoList [get]
export const getServerInfoList = (params) => {
  return service({
    url: '/serverInfo/getServerInfoList',
    method: 'get',
    params
  })
}
