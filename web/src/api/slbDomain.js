import service from '@/utils/request'

// @Tags SlbDomain
// @Summary 创建域名管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbDomain true "创建域名管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /slbDomain/createSlbDomain [post]
export const createSlbDomain = (data) => {
  return service({
    url: '/slbDomain/createSlbDomain',
    method: 'post',
    data
  })
}

// @Tags SlbDomain
// @Summary 删除域名管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbDomain true "删除域名管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbDomain/deleteSlbDomain [delete]
export const deleteSlbDomain = (data) => {
  return service({
    url: '/slbDomain/deleteSlbDomain',
    method: 'delete',
    data
  })
}

// @Tags SlbDomain
// @Summary 批量删除域名管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除域名管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbDomain/deleteSlbDomain [delete]
export const deleteSlbDomainByIds = (data) => {
  return service({
    url: '/slbDomain/deleteSlbDomainByIds',
    method: 'delete',
    data
  })
}

// @Tags SlbDomain
// @Summary 更新域名管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbDomain true "更新域名管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slbDomain/updateSlbDomain [put]
export const updateSlbDomain = (data) => {
  return service({
    url: '/slbDomain/updateSlbDomain',
    method: 'put',
    data
  })
}

// @Tags SlbDomain
// @Summary 用id查询域名管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SlbDomain true "用id查询域名管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slbDomain/findSlbDomain [get]
export const findSlbDomain = (params) => {
  return service({
    url: '/slbDomain/findSlbDomain',
    method: 'get',
    params
  })
}

// @Tags SlbDomain
// @Summary 分页获取域名管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取域名管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slbDomain/getSlbDomainList [get]
export const getSlbDomainList = (params) => {
  return service({
    url: '/slbDomain/getSlbDomainList',
    method: 'get',
    params
  })
}
