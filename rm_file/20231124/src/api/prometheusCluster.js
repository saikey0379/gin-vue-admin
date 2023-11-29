import service from '@/utils/request'

// @Tags PrometheusCluster
// @Summary 创建监控集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PrometheusCluster true "创建监控集群"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /prometheusCluster/createPrometheusCluster [post]
export const createPrometheusCluster = (data) => {
  return service({
    url: '/prometheusCluster/createPrometheusCluster',
    method: 'post',
    data
  })
}

// @Tags PrometheusCluster
// @Summary 删除监控集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PrometheusCluster true "删除监控集群"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /prometheusCluster/deletePrometheusCluster [delete]
export const deletePrometheusCluster = (data) => {
  return service({
    url: '/prometheusCluster/deletePrometheusCluster',
    method: 'delete',
    data
  })
}

// @Tags PrometheusCluster
// @Summary 批量删除监控集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除监控集群"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /prometheusCluster/deletePrometheusCluster [delete]
export const deletePrometheusClusterByIds = (data) => {
  return service({
    url: '/prometheusCluster/deletePrometheusClusterByIds',
    method: 'delete',
    data
  })
}

// @Tags PrometheusCluster
// @Summary 更新监控集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PrometheusCluster true "更新监控集群"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /prometheusCluster/updatePrometheusCluster [put]
export const updatePrometheusCluster = (data) => {
  return service({
    url: '/prometheusCluster/updatePrometheusCluster',
    method: 'put',
    data
  })
}

// @Tags PrometheusCluster
// @Summary 用id查询监控集群
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PrometheusCluster true "用id查询监控集群"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /prometheusCluster/findPrometheusCluster [get]
export const findPrometheusCluster = (params) => {
  return service({
    url: '/prometheusCluster/findPrometheusCluster',
    method: 'get',
    params
  })
}

// @Tags PrometheusCluster
// @Summary 分页获取监控集群列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取监控集群列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /prometheusCluster/getPrometheusClusterList [get]
export const getPrometheusClusterList = (params) => {
  return service({
    url: '/prometheusCluster/getPrometheusClusterList',
    method: 'get',
    params
  })
}
