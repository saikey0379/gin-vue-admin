import service from '@/utils/request'

// @Tags FileBinary
// @Summary 创建可执行程序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileBinary true "创建可执行程序"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileBinary/createFileBinary [post]
export const createFileBinary = (data) => {
  return service({
    url: '/fileBinary/createFileBinary',
    method: 'post',
    data
  })
}

// @Tags FileBinary
// @Summary 删除可执行程序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileBinary true "删除可执行程序"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileBinary/deleteFileBinary [delete]
export const deleteFileBinary = (data) => {
  return service({
    url: '/fileBinary/deleteFileBinary',
    method: 'delete',
    data
  })
}

// @Tags FileBinary
// @Summary 批量删除可执行程序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除可执行程序"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileBinary/deleteFileBinary [delete]
export const deleteFileBinaryByIds = (data) => {
  return service({
    url: '/fileBinary/deleteFileBinaryByIds',
    method: 'delete',
    data
  })
}

// @Tags FileBinary
// @Summary 更新可执行程序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileBinary true "更新可执行程序"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileBinary/updateFileBinary [put]
export const updateFileBinary = (data) => {
  return service({
    url: '/fileBinary/updateFileBinary',
    method: 'put',
    data
  })
}

// @Tags FileBinary
// @Summary 用id查询可执行程序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FileBinary true "用id查询可执行程序"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileBinary/findFileBinary [get]
export const findFileBinary = (params) => {
  return service({
    url: '/fileBinary/findFileBinary',
    method: 'get',
    params
  })
}

// @Tags FileBinary
// @Summary 分页获取可执行程序列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取可执行程序列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileBinary/getFileBinaryList [get]
export const getFileBinaryList = (params) => {
  return service({
    url: '/fileBinary/getFileBinaryList',
    method: 'get',
    params
  })
}
