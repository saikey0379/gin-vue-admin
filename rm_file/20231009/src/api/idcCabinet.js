import service from '@/utils/request'

// @Tags IdcCabinet
// @Summary 创建机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcCabinet true "创建机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcCabinet/createIdcCabinet [post]
export const createIdcCabinet = (data) => {
  return service({
    url: '/idcCabinet/createIdcCabinet',
    method: 'post',
    data
  })
}

// @Tags IdcCabinet
// @Summary 删除机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcCabinet true "删除机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcCabinet/deleteIdcCabinet [delete]
export const deleteIdcCabinet = (data) => {
  return service({
    url: '/idcCabinet/deleteIdcCabinet',
    method: 'delete',
    data
  })
}

// @Tags IdcCabinet
// @Summary 批量删除机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcCabinet/deleteIdcCabinet [delete]
export const deleteIdcCabinetByIds = (data) => {
  return service({
    url: '/idcCabinet/deleteIdcCabinetByIds',
    method: 'delete',
    data
  })
}

// @Tags IdcCabinet
// @Summary 更新机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcCabinet true "更新机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcCabinet/updateIdcCabinet [put]
export const updateIdcCabinet = (data) => {
  return service({
    url: '/idcCabinet/updateIdcCabinet',
    method: 'put',
    data
  })
}

// @Tags IdcCabinet
// @Summary 用id查询机柜
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdcCabinet true "用id查询机柜"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcCabinet/findIdcCabinet [get]
export const findIdcCabinet = (params) => {
  return service({
    url: '/idcCabinet/findIdcCabinet',
    method: 'get',
    params
  })
}

// @Tags IdcCabinet
// @Summary 分页获取机柜列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取机柜列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcCabinet/getIdcCabinetList [get]
export const getIdcCabinetList = (params) => {
  return service({
    url: '/idcCabinet/getIdcCabinetList',
    method: 'get',
    params
  })
}
