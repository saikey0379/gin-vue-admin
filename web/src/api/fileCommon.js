import service from '@/utils/request'

// @Tags FileCommon
// @Summary 创建普通文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileCommon true "创建普通文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileCommon/createFileCommon [post]
export const createFileCommon = (data) => {
  return service({
    url: '/fileCommon/createFileCommon',
    method: 'post',
    data
  })
}

// @Tags FileCommon
// @Summary 删除普通文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileCommon true "删除普通文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileCommon/deleteFileCommon [delete]
export const deleteFileCommon = (data) => {
  return service({
    url: '/fileCommon/deleteFileCommon',
    method: 'delete',
    data
  })
}

// @Tags FileCommon
// @Summary 批量删除普通文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除普通文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileCommon/deleteFileCommon [delete]
export const deleteFileCommonByIds = (data) => {
  return service({
    url: '/fileCommon/deleteFileCommonByIds',
    method: 'delete',
    data
  })
}

// @Tags FileCommon
// @Summary 更新普通文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileCommon true "更新普通文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileCommon/updateFileCommon [put]
export const updateFileCommon = (data) => {
  return service({
    url: '/fileCommon/updateFileCommon',
    method: 'put',
    data
  })
}

// @Tags FileCommon
// @Summary 用id查询普通文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FileCommon true "用id查询普通文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileCommon/findFileCommon [get]
export const findFileCommon = (params) => {
  return service({
    url: '/fileCommon/findFileCommon',
    method: 'get',
    params
  })
}

// @Tags FileCommon
// @Summary 分页获取普通文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取普通文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileCommon/getFileCommonList [get]
export const getFileCommonList = (params) => {
  return service({
    url: '/fileCommon/getFileCommonList',
    method: 'get',
    params
  })
}
