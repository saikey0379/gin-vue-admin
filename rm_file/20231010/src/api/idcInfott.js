import service from '@/utils/request'

// @Tags IdcInfott
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcInfott true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcInfott/createIdcInfott [post]
export const createIdcInfott = (data) => {
  return service({
    url: '/idcInfott/createIdcInfott',
    method: 'post',
    data
  })
}

// @Tags IdcInfott
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcInfott true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcInfott/deleteIdcInfott [delete]
export const deleteIdcInfott = (data) => {
  return service({
    url: '/idcInfott/deleteIdcInfott',
    method: 'delete',
    data
  })
}

// @Tags IdcInfott
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcInfott/deleteIdcInfott [delete]
export const deleteIdcInfottByIds = (data) => {
  return service({
    url: '/idcInfott/deleteIdcInfottByIds',
    method: 'delete',
    data
  })
}

// @Tags IdcInfott
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcInfott true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcInfott/updateIdcInfott [put]
export const updateIdcInfott = (data) => {
  return service({
    url: '/idcInfott/updateIdcInfott',
    method: 'put',
    data
  })
}

// @Tags IdcInfott
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdcInfott true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcInfott/findIdcInfott [get]
export const findIdcInfott = (params) => {
  return service({
    url: '/idcInfott/findIdcInfott',
    method: 'get',
    params
  })
}

// @Tags IdcInfott
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcInfott/getIdcInfottList [get]
export const getIdcInfottList = (params) => {
  return service({
    url: '/idcInfott/getIdcInfottList',
    method: 'get',
    params
  })
}
