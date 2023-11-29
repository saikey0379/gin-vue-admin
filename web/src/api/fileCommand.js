import service from '@/utils/request'

// @Tags FileCommand
// @Summary 创建命令信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileCommand true "创建命令信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileCommand/createFileCommand [post]
export const createFileCommand = (data) => {
  return service({
    url: '/fileCommand/createFileCommand',
    method: 'post',
    data
  })
}

// @Tags FileCommand
// @Summary 删除命令信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileCommand true "删除命令信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileCommand/deleteFileCommand [delete]
export const deleteFileCommand = (data) => {
  return service({
    url: '/fileCommand/deleteFileCommand',
    method: 'delete',
    data
  })
}

// @Tags FileCommand
// @Summary 批量删除命令信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除命令信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileCommand/deleteFileCommand [delete]
export const deleteFileCommandByIds = (data) => {
  return service({
    url: '/fileCommand/deleteFileCommandByIds',
    method: 'delete',
    data
  })
}

// @Tags FileCommand
// @Summary 更新命令信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileCommand true "更新命令信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileCommand/updateFileCommand [put]
export const updateFileCommand = (data) => {
  return service({
    url: '/fileCommand/updateFileCommand',
    method: 'put',
    data
  })
}

// @Tags FileCommand
// @Summary 用id查询命令信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FileCommand true "用id查询命令信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileCommand/findFileCommand [get]
export const findFileCommand = (params) => {
  return service({
    url: '/fileCommand/findFileCommand',
    method: 'get',
    params
  })
}

// @Tags FileCommand
// @Summary 分页获取命令信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取命令信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileCommand/getFileCommandList [get]
export const getFileCommandList = (params) => {
  return service({
    url: '/fileCommand/getFileCommandList',
    method: 'get',
    params
  })
}
