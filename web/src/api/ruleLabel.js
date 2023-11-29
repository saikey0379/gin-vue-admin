import service from '@/utils/request'

// @Tags RuleLabel
// @Summary 创建规则标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RuleLabel true "创建规则标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ruleLabel/createRuleLabel [post]
export const createRuleLabel = (data) => {
  return service({
    url: '/ruleLabel/createRuleLabel',
    method: 'post',
    data
  })
}

// @Tags RuleLabel
// @Summary 删除规则标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RuleLabel true "删除规则标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ruleLabel/deleteRuleLabel [delete]
export const deleteRuleLabel = (data) => {
  return service({
    url: '/ruleLabel/deleteRuleLabel',
    method: 'delete',
    data
  })
}

// @Tags RuleLabel
// @Summary 批量删除规则标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除规则标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ruleLabel/deleteRuleLabel [delete]
export const deleteRuleLabelByIds = (data) => {
  return service({
    url: '/ruleLabel/deleteRuleLabelByIds',
    method: 'delete',
    data
  })
}

// @Tags RuleLabel
// @Summary 更新规则标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RuleLabel true "更新规则标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ruleLabel/updateRuleLabel [put]
export const updateRuleLabel = (data) => {
  return service({
    url: '/ruleLabel/updateRuleLabel',
    method: 'put',
    data
  })
}

// @Tags RuleLabel
// @Summary 用id查询规则标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RuleLabel true "用id查询规则标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ruleLabel/findRuleLabel [get]
export const findRuleLabel = (params) => {
  return service({
    url: '/ruleLabel/findRuleLabel',
    method: 'get',
    params
  })
}

// @Tags RuleLabel
// @Summary 分页获取规则标签列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取规则标签列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ruleLabel/getRuleLabelList [get]
export const getRuleLabelList = (params) => {
  return service({
    url: '/ruleLabel/getRuleLabelList',
    method: 'get',
    params
  })
}
