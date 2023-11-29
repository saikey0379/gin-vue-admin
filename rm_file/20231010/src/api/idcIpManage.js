import service from '@/utils/request'

// @Tags IdcIpManage
// @Summary 创建数据中心管理网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpManage true "创建数据中心管理网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /idcIpManage/createIdcIpManage [post]
export const createIdcIpManage = (data) => {
  return service({
    url: '/idcIpManage/createIdcIpManage',
    method: 'post',
    data
  })
}

// @Tags IdcIpManage
// @Summary 删除数据中心管理网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpManage true "删除数据中心管理网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpManage/deleteIdcIpManage [delete]
export const deleteIdcIpManage = (data) => {
  return service({
    url: '/idcIpManage/deleteIdcIpManage',
    method: 'delete',
    data
  })
}

// @Tags IdcIpManage
// @Summary 批量删除数据中心管理网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心管理网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idcIpManage/deleteIdcIpManage [delete]
export const deleteIdcIpManageByIds = (data) => {
  return service({
    url: '/idcIpManage/deleteIdcIpManageByIds',
    method: 'delete',
    data
  })
}

// @Tags IdcIpManage
// @Summary 更新数据中心管理网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcIpManage true "更新数据中心管理网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idcIpManage/updateIdcIpManage [put]
export const updateIdcIpManage = (data) => {
  return service({
    url: '/idcIpManage/updateIdcIpManage',
    method: 'put',
    data
  })
}

// @Tags IdcIpManage
// @Summary 用id查询数据中心管理网段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdcIpManage true "用id查询数据中心管理网段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idcIpManage/findIdcIpManage [get]
export const findIdcIpManage = (params) => {
  return service({
    url: '/idcIpManage/findIdcIpManage',
    method: 'get',
    params
  })
}

// @Tags IdcIpManage
// @Summary 分页获取数据中心管理网段列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心管理网段列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idcIpManage/getIdcIpManageList [get]
export const getIdcIpManageList = (params) => {
  return service({
    url: '/idcIpManage/getIdcIpManageList',
    method: 'get',
    params
  })
}
