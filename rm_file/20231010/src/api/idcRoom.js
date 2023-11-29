import service from '@/utils/request'

// @Tags IdcRoom
// @Summary 创建数据中心房间
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcRoom true "创建数据中心房间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /room/createIdcRoom [post]
export const createIdcRoom = (data) => {
  return service({
    url: '/room/createIdcRoom',
    method: 'post',
    data
  })
}

// @Tags IdcRoom
// @Summary 删除数据中心房间
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcRoom true "删除数据中心房间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /room/deleteIdcRoom [delete]
export const deleteIdcRoom = (data) => {
  return service({
    url: '/room/deleteIdcRoom',
    method: 'delete',
    data
  })
}

// @Tags IdcRoom
// @Summary 批量删除数据中心房间
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据中心房间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /room/deleteIdcRoom [delete]
export const deleteIdcRoomByIds = (data) => {
  return service({
    url: '/room/deleteIdcRoomByIds',
    method: 'delete',
    data
  })
}

// @Tags IdcRoom
// @Summary 更新数据中心房间
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdcRoom true "更新数据中心房间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /room/updateIdcRoom [put]
export const updateIdcRoom = (data) => {
  return service({
    url: '/room/updateIdcRoom',
    method: 'put',
    data
  })
}

// @Tags IdcRoom
// @Summary 用id查询数据中心房间
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdcRoom true "用id查询数据中心房间"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /room/findIdcRoom [get]
export const findIdcRoom = (params) => {
  return service({
    url: '/room/findIdcRoom',
    method: 'get',
    params
  })
}

// @Tags IdcRoom
// @Summary 分页获取数据中心房间列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据中心房间列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /room/getIdcRoomList [get]
export const getIdcRoomList = (params) => {
  return service({
    url: '/room/getIdcRoomList',
    method: 'get',
    params
  })
}
