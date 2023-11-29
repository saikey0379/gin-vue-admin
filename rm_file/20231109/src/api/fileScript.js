import service from '@/utils/request'

// @Tags FileScript
// @Summary 创建脚本文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileScript true "创建脚本文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileScript/createFileScript [post]
export const createFileScript = (data) => {
  return service({
    url: '/fileScript/createFileScript',
    method: 'post',
    data
  })
}

// @Tags FileScript
// @Summary 删除脚本文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileScript true "删除脚本文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileScript/deleteFileScript [delete]
export const deleteFileScript = (data) => {
  return service({
    url: '/fileScript/deleteFileScript',
    method: 'delete',
    data
  })
}

// @Tags FileScript
// @Summary 批量删除脚本文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除脚本文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileScript/deleteFileScript [delete]
export const deleteFileScriptByIds = (data) => {
  return service({
    url: '/fileScript/deleteFileScriptByIds',
    method: 'delete',
    data
  })
}

// @Tags FileScript
// @Summary 更新脚本文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileScript true "更新脚本文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileScript/updateFileScript [put]
export const updateFileScript = (data) => {
  return service({
    url: '/fileScript/updateFileScript',
    method: 'put',
    data
  })
}

// @Tags FileScript
// @Summary 用id查询脚本文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FileScript true "用id查询脚本文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileScript/findFileScript [get]
export const findFileScript = (params) => {
  return service({
    url: '/fileScript/findFileScript',
    method: 'get',
    params
  })
}

// @Tags FileScript
// @Summary 分页获取脚本文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取脚本文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileScript/getFileScriptList [get]
export const getFileScriptList = (params) => {
  return service({
    url: '/fileScript/getFileScriptList',
    method: 'get',
    params
  })
}
