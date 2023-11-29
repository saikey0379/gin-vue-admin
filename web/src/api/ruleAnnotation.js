import service from '@/utils/request'

// @Tags RuleAnnotation
// @Summary 创建注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RuleAnnotation true "创建注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ruleAnnotation/createRuleAnnotation [post]
export const createRuleAnnotation = (data) => {
  return service({
    url: '/ruleAnnotation/createRuleAnnotation',
    method: 'post',
    data
  })
}

// @Tags RuleAnnotation
// @Summary 删除注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RuleAnnotation true "删除注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ruleAnnotation/deleteRuleAnnotation [delete]
export const deleteRuleAnnotation = (data) => {
  return service({
    url: '/ruleAnnotation/deleteRuleAnnotation',
    method: 'delete',
    data
  })
}

// @Tags RuleAnnotation
// @Summary 批量删除注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ruleAnnotation/deleteRuleAnnotation [delete]
export const deleteRuleAnnotationByIds = (data) => {
  return service({
    url: '/ruleAnnotation/deleteRuleAnnotationByIds',
    method: 'delete',
    data
  })
}

// @Tags RuleAnnotation
// @Summary 更新注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RuleAnnotation true "更新注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ruleAnnotation/updateRuleAnnotation [put]
export const updateRuleAnnotation = (data) => {
  return service({
    url: '/ruleAnnotation/updateRuleAnnotation',
    method: 'put',
    data
  })
}

// @Tags RuleAnnotation
// @Summary 用id查询注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RuleAnnotation true "用id查询注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ruleAnnotation/findRuleAnnotation [get]
export const findRuleAnnotation = (params) => {
  return service({
    url: '/ruleAnnotation/findRuleAnnotation',
    method: 'get',
    params
  })
}

// @Tags RuleAnnotation
// @Summary 分页获取注解标签列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取注解标签列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ruleAnnotation/getRuleAnnotationList [get]
export const getRuleAnnotationList = (params) => {
  return service({
    url: '/ruleAnnotation/getRuleAnnotationList',
    method: 'get',
    params
  })
}
