import service from '@/utils/request'

// @Tags RuleGroup
// @Summary 创建注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RuleGroup true "创建注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ruleGroup/createRuleGroup [post]
export const createRuleGroup = (data) => {
  return service({
    url: '/ruleGroup/createRuleGroup',
    method: 'post',
    data
  })
}

// @Tags RuleGroup
// @Summary 删除注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RuleGroup true "删除注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ruleGroup/deleteRuleGroup [delete]
export const deleteRuleGroup = (data) => {
  return service({
    url: '/ruleGroup/deleteRuleGroup',
    method: 'delete',
    data
  })
}

// @Tags RuleGroup
// @Summary 批量删除注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ruleGroup/deleteRuleGroup [delete]
export const deleteRuleGroupByIds = (data) => {
  return service({
    url: '/ruleGroup/deleteRuleGroupByIds',
    method: 'delete',
    data
  })
}

// @Tags RuleGroup
// @Summary 更新注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RuleGroup true "更新注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ruleGroup/updateRuleGroup [put]
export const updateRuleGroup = (data) => {
  return service({
    url: '/ruleGroup/updateRuleGroup',
    method: 'put',
    data
  })
}

// @Tags RuleGroup
// @Summary 用id查询注解标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RuleGroup true "用id查询注解标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ruleGroup/findRuleGroup [get]
export const findRuleGroup = (params) => {
  return service({
    url: '/ruleGroup/findRuleGroup',
    method: 'get',
    params
  })
}

// @Tags RuleGroup
// @Summary 分页获取注解标签列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取注解标签列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ruleGroup/getRuleGroupList [get]
export const getRuleGroupList = (params) => {
  return service({
    url: '/ruleGroup/getRuleGroupList',
    method: 'get',
    params
  })
}
