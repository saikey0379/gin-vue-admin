import service from '@/utils/request'

// @Tags FileConfig
// @Summary 创建配置文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileConfig true "创建配置文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /fileConfig/createFileConfig [post]
export const createFileConfig = (data) => {
  return service({
    url: '/fileConfig/createFileConfig',
    method: 'post',
    data
  })
}

// @Tags FileConfig
// @Summary 删除配置文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileConfig true "删除配置文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileConfig/deleteFileConfig [delete]
export const deleteFileConfig = (data) => {
  return service({
    url: '/fileConfig/deleteFileConfig',
    method: 'delete',
    data
  })
}

// @Tags FileConfig
// @Summary 批量删除配置文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除配置文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileConfig/deleteFileConfig [delete]
export const deleteFileConfigByIds = (data) => {
  return service({
    url: '/fileConfig/deleteFileConfigByIds',
    method: 'delete',
    data
  })
}

// @Tags FileConfig
// @Summary 更新配置文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FileConfig true "更新配置文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileConfig/updateFileConfig [put]
export const updateFileConfig = (data) => {
  return service({
    url: '/fileConfig/updateFileConfig',
    method: 'put',
    data
  })
}

// @Tags FileConfig
// @Summary 用id查询配置文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FileConfig true "用id查询配置文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileConfig/findFileConfig [get]
export const findFileConfig = (params) => {
  return service({
    url: '/fileConfig/findFileConfig',
    method: 'get',
    params
  })
}

// @Tags FileConfig
// @Summary 分页获取配置文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取配置文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileConfig/getFileConfigList [get]
export const getFileConfigList = (params) => {
  return service({
    url: '/fileConfig/getFileConfigList',
    method: 'get',
    params
  })
}
