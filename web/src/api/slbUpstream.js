import service from '@/utils/request'

// @Tags SlbUpstream
// @Summary 创建目标节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbUpstream true "创建目标节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /slbUpstream/createSlbUpstream [post]
export const createSlbUpstream = (data) => {
  return service({
    url: '/slbUpstream/createSlbUpstream',
    method: 'post',
    data
  })
}

// @Tags SlbUpstream
// @Summary 删除目标节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbUpstream true "删除目标节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbUpstream/deleteSlbUpstream [delete]
export const deleteSlbUpstream = (data) => {
  return service({
    url: '/slbUpstream/deleteSlbUpstream',
    method: 'delete',
    data
  })
}

// @Tags SlbUpstream
// @Summary 批量删除目标节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除目标节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbUpstream/deleteSlbUpstream [delete]
export const deleteSlbUpstreamByIds = (data) => {
  return service({
    url: '/slbUpstream/deleteSlbUpstreamByIds',
    method: 'delete',
    data
  })
}

// @Tags SlbUpstream
// @Summary 更新目标节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbUpstream true "更新目标节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slbUpstream/updateSlbUpstream [put]
export const updateSlbUpstream = (data) => {
  return service({
    url: '/slbUpstream/updateSlbUpstream',
    method: 'put',
    data
  })
}

// @Tags SlbUpstream
// @Summary 用id查询目标节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SlbUpstream true "用id查询目标节点"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slbUpstream/findSlbUpstream [get]
export const findSlbUpstream = (params) => {
  return service({
    url: '/slbUpstream/findSlbUpstream',
    method: 'get',
    params
  })
}

// @Tags SlbUpstream
// @Summary 分页获取目标节点列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取目标节点列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slbUpstream/getSlbUpstreamList [get]
export const getSlbUpstreamList = (params) => {
  return service({
    url: '/slbUpstream/getSlbUpstreamList',
    method: 'get',
    params
  })
}
