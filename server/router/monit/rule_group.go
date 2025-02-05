package monit

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RuleGroupRouter struct {
}

// InitRuleGroupRouter 初始化 规则分组 路由信息
func (s *RuleGroupRouter) InitRuleGroupRouter(Router *gin.RouterGroup) {
	ruleGroupRouter := Router.Group("ruleGroup").Use(middleware.OperationRecord())
	ruleGroupRouterWithoutRecord := Router.Group("ruleGroup")
	var ruleGroupApi = v1.ApiGroupApp.MonitApiGroup.RuleGroupApi
	{
		ruleGroupRouter.POST("createRuleGroup", ruleGroupApi.CreateRuleGroup)             // 新建规则分组
		ruleGroupRouter.DELETE("deleteRuleGroup", ruleGroupApi.DeleteRuleGroup)           // 删除规则分组
		ruleGroupRouter.DELETE("deleteRuleGroupByIds", ruleGroupApi.DeleteRuleGroupByIds) // 批量删除规则分组
		ruleGroupRouter.PUT("updateRuleGroup", ruleGroupApi.UpdateRuleGroup)              // 更新规则分组
	}
	{
		ruleGroupRouterWithoutRecord.GET("findRuleGroup", ruleGroupApi.FindRuleGroup)       // 根据ID获取规则分组
		ruleGroupRouterWithoutRecord.GET("getRuleGroupList", ruleGroupApi.GetRuleGroupList) // 获取规则分组列表
	}
}
