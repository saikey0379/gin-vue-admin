import service from '@/utils/request'

// @Tags IdcInfo1
// @Summary 创建数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcInfo1 true "创建数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcInfo1/createIdcInfo1 [post]
export const createIdcInfo1 = (data) => {
  return service({
    url: '/idcInfo1/createIdcInfo1',
    method: 'post',
    data
  })
}

// @Tags IdcInfo1
// @Summary 删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcInfo1 true "删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcInfo1/deleteIdcInfo1 [delete]
export const deleteIdcInfo1 = (data) => {
  return service({
    url: '/idcInfo1/deleteIdcInfo1',
    method: 'delete',
    data
  })
}

// @Tags IdcInfo1
// @Summary 批量删除数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcInfo1/deleteIdcInfo1 [delete]
export const deleteIdcInfo1ByIds = (data) => {
  return service({
    url: '/idcInfo1/deleteIdcInfo1ByIds',
    method: 'delete',
    data
  })
}

// @Tags IdcInfo1
// @Summary 更新数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcInfo1 true "更新数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcInfo1/updateIdcInfo1 [put]
export const updateIdcInfo1 = (data) => {
  return service({
    url: '/idcInfo1/updateIdcInfo1',
    method: 'put',
    data
  })
}

// @Tags IdcInfo1
// @Summary 用id查询数据中心
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdcInfo1 true "用id查询数据中心"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcInfo1/findIdcInfo1 [get]
export const findIdcInfo1 = (params) => {
  return service({
    url: '/idcInfo1/findIdcInfo1',
    method: 'get',
    params
  })
}

// @Tags IdcInfo1
// @Summary 分页获取数据中心列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcInfo1/getIdcInfo1List [get]
export const getIdcInfo1List = (params) => {
  return service({
    url: '/idcInfo1/getIdcInfo1List',
    method: 'get',
    params
  })
}
