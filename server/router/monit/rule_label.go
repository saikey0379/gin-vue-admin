package monit

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RuleLabelRouter struct {
}

// InitRuleLabelRouter 初始化 规则标签 路由信息
func (s *RuleLabelRouter) InitRuleLabelRouter(Router *gin.RouterGroup) {
	ruleLabelRouter := Router.Group("ruleLabel").Use(middleware.OperationRecord())
	ruleLabelRouterWithoutRecord := Router.Group("ruleLabel")
	var ruleLabelApi = v1.ApiGroupApp.MonitApiGroup.RuleLabelApi
	{
		ruleLabelRouter.POST("createRuleLabel", ruleLabelApi.CreateRuleLabel)   // 新建规则标签
		ruleLabelRouter.DELETE("deleteRuleLabel", ruleLabelApi.DeleteRuleLabel) // 删除规则标签
		ruleLabelRouter.DELETE("deleteRuleLabelByIds", ruleLabelApi.DeleteRuleLabelByIds) // 批量删除规则标签
		ruleLabelRouter.PUT("updateRuleLabel", ruleLabelApi.UpdateRuleLabel)    // 更新规则标签
	}
	{
		ruleLabelRouterWithoutRecord.GET("findRuleLabel", ruleLabelApi.FindRuleLabel)        // 根据ID获取规则标签
		ruleLabelRouterWithoutRecord.GET("getRuleLabelList", ruleLabelApi.GetRuleLabelList)  // 获取规则标签列表
	}
}
