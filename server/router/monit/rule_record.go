package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RuleRecordRouter struct {
}

// InitRuleRecordRouter 初始化 规则配置 路由信息
func (s *RuleRecordRouter) InitRuleRecordRouter(Router *gin.RouterGroup) {
	ruleRecordRouter := Router.Group("ruleRecord").Use(middleware.OperationRecord())
	ruleRecordRouterWithoutRecord := Router.Group("ruleRecord")
	var ruleRecordApi = v1.ApiGroupApp.MonitApiGroup.RuleRecordApi
	{
		ruleRecordRouter.POST("createRuleRecord", ruleRecordApi.CreateRuleRecord)   // 新建规则配置
		ruleRecordRouter.DELETE("deleteRuleRecord", ruleRecordApi.DeleteRuleRecord) // 删除规则配置
		ruleRecordRouter.DELETE("deleteRuleRecordByIds", ruleRecordApi.DeleteRuleRecordByIds) // 批量删除规则配置
		ruleRecordRouter.PUT("updateRuleRecord", ruleRecordApi.UpdateRuleRecord)    // 更新规则配置
	}
	{
		ruleRecordRouterWithoutRecord.GET("findRuleRecord", ruleRecordApi.FindRuleRecord)        // 根据ID获取规则配置
		ruleRecordRouterWithoutRecord.GET("getRuleRecordList", ruleRecordApi.GetRuleRecordList)  // 获取规则配置列表
	}
}
