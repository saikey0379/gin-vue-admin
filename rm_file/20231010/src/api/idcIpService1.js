import service from '@/utils/request'

// @Tags IdcIpService1
// @Summary 创建数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpService1 true "创建数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpService1/createIdcIpService1 [post]
export const createIdcIpService1 = (data) => {
  return service({
    url: '/idcIpService1/createIdcIpService1',
    method: 'post',
    data
  })
}

// @Tags IdcIpService1
// @Summary 删除数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpService1 true "删除数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpService1/deleteIdcIpService1 [delete]
export const deleteIdcIpService1 = (data) => {
  return service({
    url: '/idcIpService1/deleteIdcIpService1',
    method: 'delete',
    data
  })
}

// @Tags IdcIpService1
// @Summary 批量删除数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpService1/deleteIdcIpService1 [delete]
export const deleteIdcIpService1ByIds = (data) => {
  return service({
    url: '/idcIpService1/deleteIdcIpService1ByIds',
    method: 'delete',
    data
  })
}

// @Tags IdcIpService1
// @Summary 更新数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpService1 true "更新数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpService1/updateIdcIpService1 [put]
export const updateIdcIpService1 = (data) => {
  return service({
    url: '/idcIpService1/updateIdcIpService1',
    method: 'put',
    data
  })
}

// @Tags IdcIpService1
// @Summary 用id查询数据中心业务网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdcIpService1 true "用id查询数据中心业务网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpService1/findIdcIpService1 [get]
export const findIdcIpService1 = (params) => {
  return service({
    url: '/idcIpService1/findIdcIpService1',
    method: 'get',
    params
  })
}

// @Tags IdcIpService1
// @Summary 分页获取数据中心业务网段列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心业务网段列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpService1/getIdcIpService1List [get]
export const getIdcIpService1List = (params) => {
  return service({
    url: '/idcIpService1/getIdcIpService1List',
    method: 'get',
    params
  })
}
