import service from '@/utils/request'

// @Tags IdcInfot
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcInfot true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcInfot/createIdcInfot [post]
export const createIdcInfot = (data) => {
  return service({
    url: '/idcInfot/createIdcInfot',
    method: 'post',
    data
  })
}

// @Tags IdcInfot
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcInfot true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcInfot/deleteIdcInfot [delete]
export const deleteIdcInfot = (data) => {
  return service({
    url: '/idcInfot/deleteIdcInfot',
    method: 'delete',
    data
  })
}

// @Tags IdcInfot
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcInfot/deleteIdcInfot [delete]
export const deleteIdcInfotByIds = (data) => {
  return service({
    url: '/idcInfot/deleteIdcInfotByIds',
    method: 'delete',
    data
  })
}

// @Tags IdcInfot
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcInfot true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcInfot/updateIdcInfot [put]
export const updateIdcInfot = (data) => {
  return service({
    url: '/idcInfot/updateIdcInfot',
    method: 'put',
    data
  })
}

// @Tags IdcInfot
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdcInfot true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcInfot/findIdcInfot [get]
export const findIdcInfot = (params) => {
  return service({
    url: '/idcInfot/findIdcInfot',
    method: 'get',
    params
  })
}

// @Tags IdcInfot
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcInfot/getIdcInfotList [get]
export const getIdcInfotList = (params) => {
  return service({
    url: '/idcInfot/getIdcInfotList',
    method: 'get',
    params
  })
}
