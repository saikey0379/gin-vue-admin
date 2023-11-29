import service from '@/utils/request'

// @Tags SlbCluster
// @Summary 创建集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbCluster true "创建集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /slbCluster/createSlbCluster [post]
export const createSlbCluster = (data) => {
  return service({
    url: '/slbCluster/createSlbCluster',
    method: 'post',
    data
  })
}

// @Tags SlbCluster
// @Summary 删除集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbCluster true "删除集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbCluster/deleteSlbCluster [delete]
export const deleteSlbCluster = (data) => {
  return service({
    url: '/slbCluster/deleteSlbCluster',
    method: 'delete',
    data
  })
}

// @Tags SlbCluster
// @Summary 批量删除集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbCluster/deleteSlbCluster [delete]
export const deleteSlbClusterByIds = (data) => {
  return service({
    url: '/slbCluster/deleteSlbClusterByIds',
    method: 'delete',
    data
  })
}

// @Tags SlbCluster
// @Summary 更新集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbCluster true "更新集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slbCluster/updateSlbCluster [put]
export const updateSlbCluster = (data) => {
  return service({
    url: '/slbCluster/updateSlbCluster',
    method: 'put',
    data
  })
}

// @Tags SlbCluster
// @Summary 用id查询集群管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SlbCluster true "用id查询集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slbCluster/findSlbCluster [get]
export const findSlbCluster = (params) => {
  return service({
    url: '/slbCluster/findSlbCluster',
    method: 'get',
    params
  })
}

// @Tags SlbCluster
// @Summary 分页获取集群管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取集群管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slbCluster/getSlbClusterList [get]
export const getSlbClusterList = (params) => {
  return service({
    url: '/slbCluster/getSlbClusterList',
    method: 'get',
    params
  })
}
