import service from '@/utils/request'

// @Tags CmdbRegion
// @Summary 创建区域信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmdbRegion true "创建区域信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmdbRegion/createCmdbRegion [post]
export const createCmdbRegion = (data) => {
  return service({
    url: '/cmdbRegion/createCmdbRegion',
    method: 'post',
    data
  })
}

// @Tags CmdbRegion
// @Summary 删除区域信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmdbRegion true "删除区域信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmdbRegion/deleteCmdbRegion [delete]
export const deleteCmdbRegion = (data) => {
  return service({
    url: '/cmdbRegion/deleteCmdbRegion',
    method: 'delete',
    data
  })
}

// @Tags CmdbRegion
// @Summary 批量删除区域信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除区域信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmdbRegion/deleteCmdbRegion [delete]
export const deleteCmdbRegionByIds = (data) => {
  return service({
    url: '/cmdbRegion/deleteCmdbRegionByIds',
    method: 'delete',
    data
  })
}

// @Tags CmdbRegion
// @Summary 更新区域信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmdbRegion true "更新区域信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmdbRegion/updateCmdbRegion [put]
export const updateCmdbRegion = (data) => {
  return service({
    url: '/cmdbRegion/updateCmdbRegion',
    method: 'put',
    data
  })
}

// @Tags CmdbRegion
// @Summary 用id查询区域信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmdbRegion true "用id查询区域信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmdbRegion/findCmdbRegion [get]
export const findCmdbRegion = (params) => {
  return service({
    url: '/cmdbRegion/findCmdbRegion',
    method: 'get',
    params
  })
}

// @Tags CmdbRegion
// @Summary 分页获取区域信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取区域信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmdbRegion/getCmdbRegionList [get]
export const getCmdbRegionList = (params) => {
  return service({
    url: '/cmdbRegion/getCmdbRegionList',
    method: 'get',
    params
  })
}
