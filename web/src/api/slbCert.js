import service from '@/utils/request'

// @Tags SlbCert
// @Summary 创建证书管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbCert true "创建证书管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /slbCert/createSlbCert [post]
export const createSlbCert = (data) => {
  return service({
    url: '/slbCert/createSlbCert',
    method: 'post',
    data
  })
}

// @Tags SlbCert
// @Summary 删除证书管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbCert true "删除证书管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbCert/deleteSlbCert [delete]
export const deleteSlbCert = (data) => {
  return service({
    url: '/slbCert/deleteSlbCert',
    method: 'delete',
    data
  })
}

// @Tags SlbCert
// @Summary 批量删除证书管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除证书管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slbCert/deleteSlbCert [delete]
export const deleteSlbCertByIds = (data) => {
  return service({
    url: '/slbCert/deleteSlbCertByIds',
    method: 'delete',
    data
  })
}

// @Tags SlbCert
// @Summary 更新证书管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SlbCert true "更新证书管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slbCert/updateSlbCert [put]
export const updateSlbCert = (data) => {
  return service({
    url: '/slbCert/updateSlbCert',
    method: 'put',
    data
  })
}

// @Tags SlbCert
// @Summary 用id查询证书管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SlbCert true "用id查询证书管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slbCert/findSlbCert [get]
export const findSlbCert = (params) => {
  return service({
    url: '/slbCert/findSlbCert',
    method: 'get',
    params
  })
}

// @Tags SlbCert
// @Summary 分页获取证书管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取证书管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slbCert/getSlbCertList [get]
export const getSlbCertList = (params) => {
  return service({
    url: '/slbCert/getSlbCertList',
    method: 'get',
    params
  })
}
