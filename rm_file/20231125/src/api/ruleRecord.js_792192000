import service from '@/utils/request'

// @Tags RuleRecord
// @Summary 创建规则配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RuleRecord true "创建规则配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ruleRecord/createRuleRecord [post]
export const createRuleRecord = (data) => {
  return service({
    url: '/ruleRecord/createRuleRecord',
    method: 'post',
    data
  })
}

// @Tags RuleRecord
// @Summary 删除规则配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RuleRecord true "删除规则配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ruleRecord/deleteRuleRecord [delete]
export const deleteRuleRecord = (data) => {
  return service({
    url: '/ruleRecord/deleteRuleRecord',
    method: 'delete',
    data
  })
}

// @Tags RuleRecord
// @Summary 批量删除规则配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除规则配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ruleRecord/deleteRuleRecord [delete]
export const deleteRuleRecordByIds = (data) => {
  return service({
    url: '/ruleRecord/deleteRuleRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags RuleRecord
// @Summary 更新规则配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RuleRecord true "更新规则配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ruleRecord/updateRuleRecord [put]
export const updateRuleRecord = (data) => {
  return service({
    url: '/ruleRecord/updateRuleRecord',
    method: 'put',
    data
  })
}

// @Tags RuleRecord
// @Summary 用id查询规则配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RuleRecord true "用id查询规则配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ruleRecord/findRuleRecord [get]
export const findRuleRecord = (params) => {
  return service({
    url: '/ruleRecord/findRuleRecord',
    method: 'get',
    params
  })
}

// @Tags RuleRecord
// @Summary 分页获取规则配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取规则配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ruleRecord/getRuleRecordList [get]
export const getRuleRecordList = (params) => {
  return service({
    url: '/ruleRecord/getRuleRecordList',
    method: 'get',
    params
  })
}
