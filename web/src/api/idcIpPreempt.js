import service from '@/utils/request'

// @Tags IdcIpPreempt
// @Summary 创建数据中心地址预占
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpPreempt true "创建数据中心地址预占"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpPreempt/createIdcIpPreempt [post]
export const createIdcIpPreempt = (data) => {
  return service({
    url: '/idcIpPreempt/createIdcIpPreempt',
    method: 'post',
    data
  })
}

// @Tags IdcIpPreempt
// @Summary 删除数据中心地址预占
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpPreempt true "删除数据中心地址预占"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpPreempt/deleteIdcIpPreempt [delete]
export const deleteIdcIpPreempt = (data) => {
  return service({
    url: '/idcIpPreempt/deleteIdcIpPreempt',
    method: 'delete',
    data
  })
}

// @Tags IdcIpPreempt
// @Summary 批量删除数据中心地址预占
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心地址预占"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpPreempt/deleteIdcIpPreempt [delete]
export const deleteIdcIpPreemptByIds = (data) => {
  return service({
    url: '/idcIpPreempt/deleteIdcIpPreemptByIds',
    method: 'delete',
    data
  })
}

// @Tags IdcIpPreempt
// @Summary 更新数据中心地址预占
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpPreempt true "更新数据中心地址预占"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpPreempt/updateIdcIpPreempt [put]
export const updateIdcIpPreempt = (data) => {
  return service({
    url: '/idcIpPreempt/updateIdcIpPreempt',
    method: 'put',
    data
  })
}

// @Tags IdcIpPreempt
// @Summary 用id查询数据中心地址预占
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdcIpPreempt true "用id查询数据中心地址预占"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpPreempt/findIdcIpPreempt [get]
export const findIdcIpPreempt = (params) => {
  return service({
    url: '/idcIpPreempt/findIdcIpPreempt',
    method: 'get',
    params
  })
}

// @Tags IdcIpPreempt
// @Summary 分页获取数据中心地址预占列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心地址预占列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpPreempt/getIdcIpPreemptList [get]
export const getIdcIpPreemptList = (params) => {
  return service({
    url: '/idcIpPreempt/getIdcIpPreemptList',
    method: 'get',
    params
  })
}
