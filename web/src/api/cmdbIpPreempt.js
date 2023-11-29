import service from '@/utils/request'

// @Tags CmdbIpPreempt
// @Summary 创建预占地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmdbIpPreempt true "创建预占地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmdbIpPreempt/createCmdbIpPreempt [post]
export const createCmdbIpPreempt = (data) => {
  return service({
    url: '/cmdbIpPreempt/createCmdbIpPreempt',
    method: 'post',
    data
  })
}

// @Tags CmdbIpPreempt
// @Summary 删除预占地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmdbIpPreempt true "删除预占地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmdbIpPreempt/deleteCmdbIpPreempt [delete]
export const deleteCmdbIpPreempt = (data) => {
  return service({
    url: '/cmdbIpPreempt/deleteCmdbIpPreempt',
    method: 'delete',
    data
  })
}

// @Tags CmdbIpPreempt
// @Summary 批量删除预占地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除预占地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmdbIpPreempt/deleteCmdbIpPreempt [delete]
export const deleteCmdbIpPreemptByIds = (data) => {
  return service({
    url: '/cmdbIpPreempt/deleteCmdbIpPreemptByIds',
    method: 'delete',
    data
  })
}

// @Tags CmdbIpPreempt
// @Summary 更新预占地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmdbIpPreempt true "更新预占地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmdbIpPreempt/updateCmdbIpPreempt [put]
export const updateCmdbIpPreempt = (data) => {
  return service({
    url: '/cmdbIpPreempt/updateCmdbIpPreempt',
    method: 'put',
    data
  })
}

// @Tags CmdbIpPreempt
// @Summary 用id查询预占地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmdbIpPreempt true "用id查询预占地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmdbIpPreempt/findCmdbIpPreempt [get]
export const findCmdbIpPreempt = (params) => {
  return service({
    url: '/cmdbIpPreempt/findCmdbIpPreempt',
    method: 'get',
    params
  })
}

// @Tags CmdbIpPreempt
// @Summary 分页获取预占地址列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取预占地址列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmdbIpPreempt/getCmdbIpPreemptList [get]
export const getCmdbIpPreemptList = (params) => {
  return service({
    url: '/cmdbIpPreempt/getCmdbIpPreemptList',
    method: 'get',
    params
  })
}
